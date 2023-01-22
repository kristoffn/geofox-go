package geofox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kristoffn/geofox-go/model"
)

func (a *API) Init(ctx context.Context) (*model.InitResponse, error) {
	uri := fmt.Sprintf("%s://%s%s/init", defaultScheme, defaultHostname, defaultBasePath)

	req := model.InitRequest{}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal InitRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	responseBytes, err := a.makeRequest(ctx, http.MethodPost, uri, payload)
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
