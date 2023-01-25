package geofox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (a *API) Init(ctx context.Context) (*InitResponse, error) {
	req := InitRequest{}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal InitRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	resp, err := a.sendRequest(ctx, http.MethodPost, "/init", payload)
	if err != nil {
		return nil, err
	}

	var initResp InitResponse
	if err = json.Unmarshal(resp.Body, &initResp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	a.initialized = true
	return &initResp, nil
}

func (a *API) ListStations(ctx context.Context, modTypes []ModificationType, coordType CoordinateType) (*LSResponse, error) {
	req := LSRequest{
		ModificationTypes: modTypes,
		CoordinateType:    coordType,
		FilterEquivalent:  true,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal LSRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	resp, err := a.sendRequest(ctx, http.MethodPost, "/listStations", payload)
	if err != nil {
		return nil, err
	}

	var lsResp LSResponse
	if err = json.Unmarshal(resp.Body, &lsResp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &lsResp, nil
}
