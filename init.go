package geofox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type InitRequest struct{}

type InitResponse struct {
	ReturnCode     string `json:"returnCode"`
	BeginOfService string `json:"beginOfService"`
	EndOfService   string `json:"endOfService"`
	ID             string `json:"id"`
	DataID         string `json:"dataId"`
	BudildDate     string `json:"buildDate"`
	BuildTime      string `json:"buildTime"`
	BuildText      string `json:"buildText"`
}

func (api *API) Init(ctx context.Context) (*InitResponse, error) {
	uri := "/init"
	res, err := api.makeRequest(ctx, http.MethodGet, uri,
		&InitRequest{},
		nil)
	if err != nil {
		return nil, err
	}
	result := InitResponse{}
	if err := json.Unmarshal(res.Body, &result); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return &result, nil
}
