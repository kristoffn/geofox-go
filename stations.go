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

func (a *API) ListStations(ctx context.Context, modTypes []model.ModificationType, coordType coordinates.CoordinateType) (*model.LSResponse, error) {
	uri := fmt.Sprintf("%s://%s%s/listStations", defaultScheme, defaultHostname, defaultBasePath)

	req := model.LSRequest{
		ModificationTypes: modTypes,
		CoordinateType:    coordType,
		FilterEquivalent:  true,
	}
	req.Language = LanguageGerman
	req.Version = 51

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal LSRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	responseBytes, err := a.makeRequest(ctx, http.MethodPost, uri, payload)
	if err != nil {
		return nil, err
	}

	var resp model.LSResponse
	if err = json.Unmarshal(responseBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &resp, nil
}
