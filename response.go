package geofox

import (
	"github.com/kristoffn/geofox-go/internal/consts"
	"github.com/kristoffn/geofox-go/internal/types"
)

type BaseResponse struct {
	ReturnCode   GeofoxReturnCode `json:"returnCode"`
	ErrorText    *string          `json:"errorText,omitempty"`
	ErrorDevInfo *string          `json:"errorDevInfo,omitempty"`
}

type TicketListResponse struct {
	BaseResponse
	TicketInfos []types.TicketListTicketInfos `json:"ticketInfos"`
}

type TariffZoneNeighboursResponse struct {
	BaseResponse
	TariffZones []types.TariffZone `json:"tariffZones"`
}

type TariffMetaDataResponse struct {
	BaseResponse
	TariffKinds    *[]types.TariffKind   `json:"tariffKinds,omitempty"`
	TariffLevels   *[]types.TariffLevel  `json:"tariffLevels,omitempty"`
	TariffCounties *[]types.TariffCounty `json:"tariffCounties,omitempty"`
	TariffZones    *[]types.TariffZone   `json:"tariffZones,omitempty"`
	TariffRings    *[]string             `json:"tariffRings,omitempty"`
}

type SingleTicketOptimizerResponse struct {
	BaseResponse
	Tickets *[]types.TariffOptimizerTicket `json:"tickets,omitempty"`
}

type LSResponse struct {
	BaseResponse
	DataReleaseID *string                   `json:"dataReleaseID,omitempty"`
	Stations      *[]types.StationListEntry `json:"stations,omitempty"`
}

type LLResponse struct {
	BaseResponse
	DataReleaseID *string                `json:"dataReleaseID,omitempty"`
	Lines         *[]types.LineListEntry `json:"lines,omitempty"`
}

type InitResponse struct {
	BaseResponse
	BeginOfService string           `json:"beginOfService"`
	EndOfService   string           `json:"endOfService"`
	ID             string           `json:"id"`
	DataID         string           `json:"dataId"`
	BuildDate      string           `json:"buildDate"`
	BuildTime      string           `json:"buildTime"`
	BuildText      string           `json:"buildText"`
	Version        string           `json:"version"`
	Properties     []types.Property `json:"properties,omitempty"`
}

type VehicleMapResponse struct {
	BaseResponse
	Journeys []types.Journey `json:"journeys"`
}

type TrackCoordinatesResponse struct {
	BaseResponse
	TrackIDs []string               `json:"trackIDs"`
	Tracks   []types.VehicleMapPath `json:"tracks"`
}

type TariffResponse struct {
	BaseResponse
	TariffInfos *[]types.TariffInfo `json:"tariffInfos,omitempty"`
}

type StationInformationResponse struct {
	BaseResponse
	PartialStations *[]types.PartialStation `json:"partialStations,omitempty"`
	LastUpdate      *types.GTITime          `json:"lastUpdate,omitzero"`
}

type GRResponse struct {
	BaseResponse
	Schedules         *[]types.Schedule      `json:"schedules,omitempty"`
	RealtimeSchedules *[]types.Schedule      `json:"realtimeSchedules,omitempty"`
	RealtimeAffected  *bool                  `json:"realtimeAffected,omitempty"` //default: false
	IndividualTrack   *types.IndividualTrack `json:"individualTrack,omitzero"`
}

type IndividualRouteResponse struct {
	BaseResponse
	Routes []types.IndividualRoute `json:"routes"`
}

type AnnouncementResponse struct {
	BaseResponse
	Announcements *[]types.Announcement `json:"announcements,omitempty"`
	LastUpdate    string                `json:"lastUpdate"`
}

type DLResponse struct {
	BaseResponse
	Time        types.GTITime                `json:"time"`
	Departures  *[]types.Departure           `json:"departures,omitempty"`
	Filter      *[]types.DLFilterEntry       `json:"filter,omitempty"`
	ServiceType *consts.TransportServiceType `json:"serviceTypes,omitempty"`
}

type DCResponse struct {
	BaseResponse
	Courseelements *[]types.CourseElement `json:"courseElements,omitempty"`
	Extra          *bool                  `json:"extra,omitempty"`     //default: false
	Cancelled      *bool                  `json:"cancelled,omitempty"` //default: false
	Attributes     *[]types.Attribute     `json:"attributes,omitempty"`
}

type PCResponse struct {
	BaseResponse
	IsHVV *bool `json:"isHVV,omitempty"`
}

type CNResponse struct {
	BaseResponse
	Results *[]types.RegionalSDName `json:"results,omitempty"`
}
