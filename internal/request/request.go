package request

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

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

func sendRequest(ctx context.Context, method, uri string, params any) (error) {

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

	resp, err := request(ctx, method, uri, reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	retrun nil
}

func request(ctx context.Context, method, uri string, reqBody io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, a.BaseURL+uri, reqBody)
	if err != nil {
		return nil, err
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

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
