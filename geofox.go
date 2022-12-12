package geofox

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

type API struct {
	APIUsername string
	APIKey      string
	BaseURL     string
	headers     http.Header
	httpClient  *http.Client
	rateLimiter *rate.Limiter
	retryPolicy RetryPolicy
	logger      Logger
	Debug       bool
}

// New initializes the API configuration
func New(username, key string, opts ...Option) (*API, error) {
	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}
	api.APIUsername = username
	api.APIKey = key

	return api, nil
}

func newClient(opts ...Option) (*API, error) {
	silentLogger := log.New(io.Discard, "", log.LstdFlags)
	api := &API{
		BaseURL:     fmt.Sprintf("%s://%s%s", defaultScheme, defaultHostname, defaultBasePath),
		headers:     make(http.Header),
		rateLimiter: rate.NewLimiter(rate.Limit(4), 1), // 4rps equates to default api limit (1200 req/5 min)
		retryPolicy: RetryPolicy{
			MaxRetries:    3,
			MinRetryDelay: time.Duration(1) * time.Second,
			MaxRetryDelay: time.Duration(30) * time.Second,
		},
		logger: silentLogger,
	}

	err := api.parseOptions(opts...)
	if err != nil {
		return nil, fmt.Errorf("options parsing failed: %w", err)
	}

	// Fall back to http.DefaultClient if the package user does not provide
	// their own.
	if api.httpClient == nil {
		api.httpClient = http.DefaultClient
	}

	return api, nil
}

type APIResponse struct {
	Body       []byte
	Status     string
	StatusCode int
	Headers    http.Header
}

func (api *API) makeRequest(ctx context.Context, method, uri string, params interface{}, headers http.Header) (*APIResponse, error) {
	var err error
	var resp *http.Response
	var respErr error
	var respBody []byte

	for i := 0; i <= api.retryPolicy.MaxRetries; i++ {
		var reqBody io.Reader
		if params != nil {
			if r, ok := params.(io.Reader); ok {
				reqBody = r
			} else if paramBytes, ok := params.([]byte); ok {
				reqBody = bytes.NewReader(paramBytes)
			} else {
				var jsonBody []byte
				jsonBody, err = json.Marshal(params)
				if err != nil {
					return nil, fmt.Errorf("error marshalling params to JSON: %w", err)
				}
				reqBody = bytes.NewReader(jsonBody)
			}
		}

		if i > 0 {
			sleepDuration := time.Duration(math.Pow(2, float64(i-1)) * float64(api.retryPolicy.MinRetryDelay))

			if sleepDuration > api.retryPolicy.MaxRetryDelay {
				sleepDuration = api.retryPolicy.MaxRetryDelay
			}
			api.logger.Printf("Sleeping %s before retry attempt number %d for request %s %s", sleepDuration.String(), i, method, uri)
			select {
			case <-time.After(sleepDuration):
			case <-ctx.Done():
				return nil, fmt.Errorf("operation aborted during backoff: %w", ctx.Err())
			}
		}

		err = api.rateLimiter.Wait(ctx)
		if err != nil {
			return nil, fmt.Errorf("error caused by request rate limiting: %w", err)
		}

		resp, respErr = api.request(ctx, method, uri, reqBody, headers)

		// short circuit processing on context timeouts
		if respErr != nil && errors.Is(respErr, context.DeadlineExceeded) {
			return nil, respErr
		}

		if respErr != nil || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
				respErr = errors.New("exceeded available rate limit retries")
			}

			if respErr == nil {
				respBody, err = io.ReadAll(resp.Body)
				resp.Body.Close()

				respErr = fmt.Errorf("could not read response body: %w", err)

				api.logger.Printf("Request: %s %s got an error response %d: %s\n", method, uri, resp.StatusCode,
					strings.Replace(strings.Replace(string(respBody), "\n", "", -1), "\t", "", -1))
			} else {
				api.logger.Printf("Error performing request: %s %s : %s \n", method, uri, respErr.Error())
			}
			continue
		} else {
			respBody, err = io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return nil, fmt.Errorf("could not read response body: %w", err)
			}
			break
		}
	}

	if respErr != nil {
		return nil, respErr
	}

	//TODO: Improve error handling
	if resp.StatusCode >= http.StatusBadRequest {
		switch resp.StatusCode {
		case http.StatusBadRequest:
			return nil, &RequestError{
				err: &Error{
					StatusCode: resp.StatusCode,
					Message:    "bad request",
				}}
		case http.StatusNotFound:
			return nil, &NotFoundError{err: &Error{
				StatusCode: resp.StatusCode,
				Message:    "not found",
			}}
		case http.StatusUnauthorized:
			return nil, &AuthorizationError{err: &Error{
				StatusCode: resp.StatusCode,
				Message:    "unauthorized",
			}}
		case http.StatusTooManyRequests:
			return nil, &RatelimitError{err: &Error{
				StatusCode: resp.StatusCode,
				Message:    "rate limit exceeded",
			}}
		case http.StatusInternalServerError:
			return nil, &ServiceError{err: &Error{
				StatusCode: resp.StatusCode,
				Message:    "internal server error",
			}}
		default:
			return nil, &ServiceError{err: &Error{
				StatusCode: resp.StatusCode,
				Message:    "unknown error",
			}}
		}

	}

	return &APIResponse{
		Body:       respBody,
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header,
	}, nil
}

func (api *API) request(ctx context.Context, method, uri string, reqBody io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, api.BaseURL+uri, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}
	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, api.headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	return resp, nil
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}

type Logger interface {
	Printf(format string, v ...interface{})
}

// RetryPolicy specifies number of retries and min/max retry delays
// This config is used when the client exponentially backs off after errored requests.
type RetryPolicy struct {
	MaxRetries    int
	MinRetryDelay time.Duration
	MaxRetryDelay time.Duration
}
