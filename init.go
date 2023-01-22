package geofox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type InitRequest struct {
	BaseRequest
}

type InitResponse struct {
	ReturnCode     string `json:"returnCode"`
	BeginOfService string `json:"beginOfService"`
	EndOfService   string `json:"endOfService"`
	ID             string `json:"id"`
	DataID         string `json:"dataId"`
	BuildDate      string `json:"buildDate"`
	BuildTime      string `json:"buildTime"`
	BuildText      string `json:"buildText"`
}

func (a *API) Init(ctx context.Context) (*InitResponse, error) {
	uri := fmt.Sprintf("%s://%s%s/init", defaultScheme, defaultHostname, defaultBasePath)

	req := InitRequest{}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal InitRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	responseBytes, err := a.makeRequest(ctx, http.MethodPost, uri, payload)
	if err != nil {
		return nil, err
	}

	var resp InitResponse
	if err = json.Unmarshal(responseBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &resp, nil
}
