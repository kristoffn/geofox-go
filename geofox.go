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
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/kristoffn/geofox-go/coordinates"
	"github.com/kristoffn/geofox-go/model"
)

const (
	AuthTypeHmacSHA1 = "HmacSHA1"
)

type API struct {
	Username    string
	Password    string
	AuthType    string
	debug       bool
	initialized bool
}

func New(username, password string, opts ...Option) (*API, error) {
	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.Username = username
	api.Password = password

	return api, nil
}

func newClient(opts ...Option) (*API, error) {
	api := &API{
		AuthType: AuthTypeHmacSHA1,
		debug:    false,
	}
	if err := api.parseOptions(opts...); err != nil {
		return nil, fmt.Errorf("failed to parse options: %w", err)
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

func (a *API) makeRequest(ctx context.Context, method, url string, payload *strings.Reader) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, payload)
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

	if a.debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		log.Println(string(dump))
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if a.debug {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			return nil, err
		}
		log.Println(string(dump))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *API) Init(ctx context.Context) (*model.InitResponse, error) {
	req := model.InitRequest{}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal InitRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	responseBytes, err := a.makeRequest(ctx, http.MethodPost, endpointInit, payload)
	if err != nil {
		return nil, err
	}

	var resp model.InitResponse
	if err = json.Unmarshal(responseBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	a.initialized = true
	return &resp, nil
}

func (a *API) ListStations(ctx context.Context, modTypes []model.ModificationType, coordType coordinates.CoordinateType) (*model.LSResponse, error) {
	req := model.LSRequest{
		ModificationTypes: modTypes,
		CoordinateType:    coordType,
		FilterEquivalent:  true,
	}
	req.Language = model.LanguageGerman
	req.Version = 51

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal LSRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	responseBytes, err := a.makeRequest(ctx, http.MethodPost, endpointListStations, payload)
	if err != nil {
		return nil, err
	}

	var resp model.LSResponse
	if err = json.Unmarshal(responseBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &resp, nil
}
