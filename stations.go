package geofox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type LSFilterType string

const (
	LSNoFilter  LSFilterType = "NO_FILTER"
	LSHVVLister LSFilterType = "HVV_LISTED"
)

type LSRequest struct {
	Language          string             `json:"language,omitempty"`
	Version           int                `json:"version,omitempty"`
	FilterType        string             `json:"filterType,omitempty"`
	DataReleaseID     string             `json:"dataReleaseID"`
	ModificationTypes []ModificationType `json:"modificationTypes"`
	CoordinateType    CoordinateType     `json:"coordinateType,omitempty"`
	FilterEquivalent  bool               `json:"filterEquivalent,omitempty"`
}

type LSResponse struct {
	DataReleaseID string             `json:"dataReleaseID,omitempty"`
	Stations      []StationListEntry `json:"stations,omitempty"`
}

type StationListEntry struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	City         string        `json:"city"`
	CombinedName string        `json:"combinedName"`
	Shortcuts    []string      `json:"shortcuts"`
	Aliasses     []string      `json:"aliasses"`
	VehicleTypes []VehicleType `json:"vehicleTypes"`
	Coordniate   Coordniate    `json:"coordinate"`
	Exists       bool          `json:"exists"`
}

type Station StationListEntry

func (a *API) ListStations(modTypes []ModificationType, coordType CoordinateType) (map[string]Station, error) {
	uri := fmt.Sprintf("%s://%s%s/listStations", defaultScheme, defaultHostname, defaultBasePath)

	req := LSRequest{
		Language:          "de",
		ModificationTypes: modTypes,
		Version:           51,
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal LSRequest object: %s", err.Error())
	}
	payload := strings.NewReader(string(reqBytes))

	responseBytes, err := a.makeRequest(http.MethodPost, uri, payload)
	if err != nil {
		return nil, err
	}

	var resp LSResponse
	if err = json.Unmarshal(responseBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to marshal body bytes: %s", err.Error())
	}

	stations := make(map[string]Station)
	for _, station := range resp.Stations {
		stations[station.ID] = Station(station)
	}
	return stations, nil
}
