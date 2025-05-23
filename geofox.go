package geofox

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/kristoffn/geofox-go/internal/consts"
	"go.uber.org/zap"
)

const (
	AuthTypeHmacSHA1 = "HmacSHA1"
)

type API struct {
	Username    string
	Password    string
	AuthType    string
	BaseURL     string
	debug       bool
	initialized bool
	logger      *zap.Logger
}

type APIResponse struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
}

func New(username, password string) (*API, error) {
	api, err := newClient()
	if err != nil {
		return nil, err
	}

	api.Username = username
	api.Password = password
	api.BaseURL = fmt.Sprintf(
		"%s://%s%s",
		consts.DefaultScheme,
		consts.DefaultHostname,
		consts.DefaultBasePath,
	)

	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}
	api.logger = logger

	return api, nil
}

func newClient() (*API, error) {
	api := &API{
		AuthType: AuthTypeHmacSHA1,
		debug:    false,
	}
	return api, nil
}

func createSHA1HmacSignature(input []byte, key []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(input)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getBodyBytes(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

func (a *API) sendRequest(ctx context.Context, method, uri string, params any) (*APIResponse, error) {
	var reqBody io.Reader
	if params != nil {
		if r, ok := params.(io.Reader); ok {
			reqBody = r
		} else if paramBytes, ok := params.([]byte); ok {
			reqBody = bytes.NewReader(paramBytes)
		} else {
			var jsonBody []byte
			jsonBody, err := json.Marshal(params)
			if err != nil {
				return nil, fmt.Errorf("error marshalling params to JSON: %w", err)
			}
			reqBody = bytes.NewReader(jsonBody)
		}
	}

	resp, err := a.request(ctx, method, uri, reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Log response details for debugging
	if a.debug {
		a.logger.Debug("API response",
			zap.String("method", method),
			zap.String("uri", uri),
			zap.Int("status_code", resp.StatusCode),
			zap.String("response_body", string(body)),
		)
	}

	// Handle error responses
	if resp.StatusCode >= http.StatusBadRequest {
		errorMsg := string(body)
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return nil, &ErrorUnauthorized{
				StatusCode: resp.StatusCode,
				ReturnCode: RCCommError,
				Message:    errorMsg,
			}
		case http.StatusForbidden:
			return nil, &ErrorForbidden{
				StatusCode: resp.StatusCode,
				ReturnCode: RCCommError,
				Message:    errorMsg,
			}
		case http.StatusNotFound:
			return nil, &ErrorNotFound{
				StatusCode: resp.StatusCode,
				ReturnCode: RCCommError,
				Message:    errorMsg,
			}
		case http.StatusTooManyRequests:
			return nil, &ErrorTooManyRequests{
				StatusCode: resp.StatusCode,
				ReturnCode: RCCommError,
				Message:    errorMsg,
			}
		case http.StatusInternalServerError:
			return nil, &ErrorInternalServerError{
				StatusCode: resp.StatusCode,
				ReturnCode: RCCommError,
				Message:    errorMsg,
			}
		default:
			return nil, &ErrorGeneric{
				StatusCode: resp.StatusCode,
				ReturnCode: RCCommError,
				Message:    errorMsg,
			}
		}
	}

	return &APIResponse{
		Body:       body,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
	}, nil
}

func (a *API) request(ctx context.Context, method, uri string, reqBody io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, a.BaseURL+uri, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	bodyBytes, err := getBodyBytes(req)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("geofox-auth-signature", createSHA1HmacSignature(bodyBytes, []byte(a.Password)))
	req.Header.Add("geofox-auth-user", a.Username)
	req.Header.Add("geofox-auth-type", a.AuthType)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("X-Platform", "web")

	if a.debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			a.logger.Error("Failed to dump request",
				zap.Error(err),
				zap.String("method", method),
				zap.String("uri", uri),
			)
		} else {
			a.logger.Debug("Request dump",
				zap.String("method", method),
				zap.String("uri", uri),
				zap.String("dump", string(dump)),
			)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if a.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			a.logger.Error("Failed to dump response",
				zap.Error(err),
				zap.String("method", method),
				zap.String("uri", uri),
			)
		} else {
			a.logger.Debug("Response dump",
				zap.String("method", method),
				zap.String("uri", uri),
				zap.String("dump", string(dump)),
			)
		}
	}

	return resp, nil
}
