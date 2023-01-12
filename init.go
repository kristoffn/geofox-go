package geofox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	initMessage = "{}"
)

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

func (a *API) Init() (*InitResponse, error) {
	uri := fmt.Sprintf("%s://%s%s/init", defaultScheme, defaultHostname, defaultBasePath)
	payload := strings.NewReader(initMessage)

	responseBytes, err := a.makeRequest(http.MethodPost, uri, payload)
	if err != nil {
		return nil, err
	}

	var resp InitResponse
	if err = json.Unmarshal(responseBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &resp, nil
}
