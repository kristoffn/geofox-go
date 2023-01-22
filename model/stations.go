package model

import "github.com/kristoffn/geofox-go/coordinates"

type LSRequest struct {
	BaseRequest
	DataReleaseID     string                     `json:"dataReleaseID"`
	ModificationTypes []ModificationType         `json:"modificationTypes"`
	CoordinateType    coordinates.CoordinateType `json:"coordinateType,omitempty"`
	FilterEquivalent  bool                       `json:"filterEquivalent,omitempty"`
}

type LSResponse struct {
	DataReleaseID string             `json:"dataReleaseID,omitempty"`
	Stations      []StationListEntry `json:"stations,omitempty"`
}

type StationListEntry struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	City         string                 `json:"city"`
	CombinedName string                 `json:"combinedName"`
	Shortcuts    []string               `json:"shortcuts"`
	Aliasses     []string               `json:"aliasses"`
	VehicleTypes []VehicleType          `json:"vehicleTypes"`
	Coordniate   coordinates.Coordniate `json:"coordinate"`
	Exists       bool                   `json:"exists"`
}
