package geofox

type BaseResponse struct {
	ReturnCode   GeofoxReturnCode `json:"returnCode"`
	ErrorText    *string          `json:"errorText,omitempty"`
	ErrorDevInfo *string          `json:"errorDevInfo,omitempty"`
}

type TicketListResponse struct {
	BaseResponse
	TicketInfos []TicketListTicketInfos `json:"ticketInfos"`
}

type TariffZoneNeighboursResponse struct {
	BaseResponse
	TariffZones []TariffZone `json:"tariffZones"`
}

type TariffMetaDataResponse struct {
	BaseResponse
	TariffKinds    *[]TariffKind   `json:"tariffKinds,omitempty"`
	TariffLevels   *[]TariffLevel  `json:"tariffLevels,omitempty"`
	TariffCounties *[]TariffCounty `json:"tariffCounties,omitempty"`
	TariffZones    *[]TariffZone   `json:"tariffZones,omitempty"`
	TariffRings    *[]string       `json:"tariffRings,omitempty"`
}

type SingleTicketOptimizerResponse struct {
	BaseResponse
	Tickets *[]TariffOptimizerTicket `json:"tickets,omitempty"`
}

type LSResponse struct {
	BaseResponse
	DataReleaseID *string             `json:"dataReleaseID,omitempty"`
	Stations      *[]StationListEntry `json:"stations,omitempty"`
}

type LLResponse struct {
	BaseResponse
	DataReleaseID *string          `json:"dataReleaseID,omitempty"`
	Lines         *[]LineListEntry `json:"lines,omitempty"`
}

type InitResponse struct {
	BaseResponse
	BeginOfService string     `json:"beginOfService"`
	EndOfService   string     `json:"endOfService"`
	ID             string     `json:"id"`
	DataID         string     `json:"dataId"`
	BuildDate      string     `json:"buildDate"`
	BuildTime      string     `json:"buildTime"`
	BuildText      string     `json:"buildText"`
	Version        string     `json:"version"`
	Properties     []Property `json:"properties,omitempty"`
}

type VehicleMapResponse struct {
	BaseResponse
	Journeys []Journey `json:"journeys"`
}

type TrackCoordinatesResponse struct {
	BaseResponse
	TrackIDs []string         `json:"trackIDs"`
	Tracks   []VehicleMapPath `json:"tracks"`
}

type TariffResponse struct {
	BaseResponse
	TariffInfos *[]TariffInfo `json:"tariffInfos,omitempty"`
}

type StationInformationResponse struct {
	BaseResponse
	PartialStations *[]PartialStation `json:"partialStations,omitempty"`
	LastUpdate      *GTITime          `json:"lastUpdate,omitzero"`
}

type GRResponse struct {
	BaseResponse
	Schedules         *[]Schedule      `json:"schedules,omitempty"`
	RealtimeSchedules *[]Schedule      `json:"realtimeSchedules,omitempty"`
	RealtimeAffected  *bool            `json:"realtimeAffected,omitempty"` //default: false
	IndividualTrack   *IndividualTrack `json:"individualTrack,omitzero"`
}

type IndividualRouteResponse struct {
	BaseResponse
	Routes []IndividualRoute `json:"routes"`
}

type AnnouncementResponse struct {
	BaseResponse
	Announcements *[]Announcement `json:"announcements,omitempty"`
	LastUpdate    string          `json:"lastUpdate"`
}

type DLResponse struct {
	BaseResponse
	Time        GTITime               `json:"time"`
	Departures  *[]Departure          `json:"departures,omitempty"`
	Filter      *[]DLFilterEntry      `json:"filter,omitempty"`
	ServiceType *TransportServiceType `json:"serviceTypes,omitempty"`
}

type DCResponse struct {
	BaseResponse
	Courseelements *[]CourseElement `json:"courseElements,omitempty"`
	Extra          *bool            `json:"extra,omitempty"`     //default: false
	Cancelled      *bool            `json:"cancelled,omitempty"` //default: false
	Attributes     *[]Attribute     `json:"attributes,omitempty"`
}

type PCResponse struct {
	BaseResponse
	IsHVV *bool `json:"isHVV,omitempty"`
}

type CNResponse struct {
	BaseResponse
	Results *[]RegionalSDName `json:"results,omitempty"`
}
