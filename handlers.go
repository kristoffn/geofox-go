package geofox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kristoffn/geofox-go/coordinates"
	"github.com/kristoffn/geofox-go/model"
)

func (a *API) Init(ctx context.Context) (*model.InitResponse, error) {
	req := model.InitRequest{}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal InitRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	resp, err := a.sendRequest(ctx, http.MethodPost, a.BaseURL+"/init", payload)
	if err != nil {
		return nil, err
	}

	var initResp model.InitResponse
	if err = json.Unmarshal(resp.Body, &initResp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	a.initialized = true
	return &initResp, nil
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

	resp, err := a.sendRequest(ctx, http.MethodPost, a.BaseURL+"/listStations", payload)
	if err != nil {
		return nil, err
	}

	var lsResp model.LSResponse
	if err = json.Unmarshal(resp.Body, &lsResp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &lsResp, nil
}
