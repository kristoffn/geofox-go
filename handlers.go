package geofox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/kristoffn/geofox-go/internal/consts"
	"github.com/kristoffn/geofox-go/internal/types"
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

func (a *API) ListStations(ctx context.Context, modTypes []consts.ModificationType, coordType consts.CoordinateType) (*LSResponse, error) {
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

func (a *API) ListLines(ctx context.Context, modTypes []consts.ModificationType, withSublines bool) (*LLResponse, error) {
	req := LLRequest{
		ModificationTypes: modTypes,
		WithSublines:      withSublines,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal LLRequest object: %s", err.Error())
	}
	reqReader := strings.NewReader(string(reqBytes))

	resp, err := a.sendRequest(ctx, http.MethodPost, "/listLines", reqReader)
	if err != nil {
		return nil, err
	}

	var llResp LLResponse
	if err = json.Unmarshal(resp.Body, &llResp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &llResp, nil
}

func (a *API) CheckByName(ctx context.Context, name string, typ consts.SDType) (*CNResponse, error) {
	sdName := types.SDName{
		Name: name,
		Type: typ,
	}
	return a.checkName(ctx, &sdName)
}

func (a *API) checkName(ctx context.Context, name *types.SDName) (*CNResponse, error) {
	if name == nil {
		return nil, errors.New("name object is nil")
	}
	req := CNRequest{
		Name:           *name,
		MaxList:        1,
		CoordinateType: consts.EPSG4326,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal CNRequest object: %s", err.Error())
	}
	reqReader := strings.NewReader(string(reqBytes))

	resp, err := a.sendRequest(ctx, http.MethodPost, "/checkName", reqReader)
	if err != nil {
		return nil, err
	}

	var cnResp CNResponse
	if err = json.Unmarshal(resp.Body, &cnResp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &cnResp, nil
}

func (a *API) GetAnnouncements(ctx context.Context, lines []string, fullAnnouncement, showBroadcastRelevant bool) (*AnnouncementResponse, error) {
	req := AnnouncementRequest{
		Names:                 lines,
		Full:                  fullAnnouncement,
		ShowBroadcastRelevant: showBroadcastRelevant,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AnnouncementRequest object: %s", err.Error())
	}
	reqReader := strings.NewReader(string(reqBytes))

	resp, err := a.sendRequest(ctx, http.MethodPost, "/getAnnouncements", reqReader)
	if err != nil {
		return nil, err
	}
	var announcementResp AnnouncementResponse
	if err = json.Unmarshal(resp.Body, &announcementResp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}
	return &announcementResp, nil
}
