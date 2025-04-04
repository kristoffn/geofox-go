package geofox

import (
	"github.com/kristoffn/geofox-go/internal/consts"
	"github.com/kristoffn/geofox-go/internal/types"
)

type BaseRequest struct {
	Language   string        `json:"language,omitempty"`   // default: de
	Version    int32         `json:"version,omitempty"`    // default: 1
	FilterType consts.Filter `json:"filterType,omitempty"` // default: NO_FILTER
}

type TicketListRequest struct {
	BaseRequest
	StationKey string `json:"stationKey,omitempty"`
}

type TariffZoneNeighboursRequest struct {
	BaseRequest
}

type TariffMetaDataRequest struct {
	BaseRequest
}

type SingleTicketOptimizerRequest struct {
	BaseRequest
	WithReturnJourney bool                                    `json:"withReturnJourney,omitempty"`
	NumberOfAdults    int32                                   `json:"numberOfAdults,omitempty"`
	NumberOfChildren  int32                                   `json:"numberOfChildren,omitempty"`
	Tickets           []types.TariffOptimizerTicket           `json:"tickets,omitempty"`
	Route             types.SingleTicketOptimizerRequestRoute `json:"route"`
}

type LSRequest struct {
	BaseRequest
	DataReleaseID     string                    `json:"dataReleaseID"`
	ModificationTypes []consts.ModificationType `json:"modificationTypes"`
	CoordinateType    consts.CoordinateType     `json:"coordinateType,omitempty"`
	FilterEquivalent  bool                      `json:"filterEquivalent,omitempty"`
}

type LLRequest struct {
	BaseRequest
	WithSublines      bool                      `json:"withSublines"`
	DataReleaseID     string                    `json:"dataReleaseID"`
	ModificationTypes []consts.ModificationType `json:"modificationTypes"`
}

type InitRequest struct {
	BaseRequest
}

type VehicleMapRequest struct {
	BaseRequest
	BoundingBox    types.BoundingBox     `json:"boundingBox"`
	PeriodBegin    int64                 `json:"periodBegin,omitempty"`
	PeriodEnd      int64                 `json:"periodEnd,omitempty"`
	WithoutCoords  bool                  `json:"withoutCoords,omitempty"`
	CoordinateType consts.CoordinateType `json:"coordinateType,omitempty"` //default: EPSG_4326
	VehicleTypes   []consts.VehicleType  `json:"vehicleTypes,omitempty"`
	RealTime       bool                  `json:"realtime,omitempty"`
}

type TrackCoordinatesRequest struct {
	BaseRequest
	CoordinateType consts.CoordinateType `json:"coordinateType,omitempty"` //default: EPSG_4326
	StopPointKeys  []string              `json:"stopPointKeys"`
}

type TariffRequest struct {
	BaseRequest
	ScheduleElements     []types.ScheduleElementLight `json:"scheduleElements"`
	Departure            types.GTITime                `json:"departure"`
	Arrival              types.GTITime                `json:"arrival"`
	ReturnReduced        bool                         `json:"returnReduced"`
	ReturnPartialTickets bool                         `json:"returnPartialTickets"`
}

type StationInformationRequest struct {
	BaseRequest
	Station types.SDName `json:"station"`
}

type GRRequest struct {
	BaseRequest
	Start                      types.SDName                `json:"start"`
	Destination                types.SDName                `json:"dest"`
	Via                        types.SDName                `json:"via,omitzero"`
	Time                       types.GTITime               `json:"time"`
	TimeIsDeparture            bool                        `json:"timeIsDeparture,omitempty"`
	WithPaths                  bool                        `json:"withPaths,omitempty"`
	NumberOfSchedules          int32                       `json:"numberOfSchedules,omitempty"`
	Penalties                  []types.Penalty             `json:"penalties,omitempty"`
	TariffDetails              bool                        `json:"tariffDetails,omitempty"`   //default: false
	ContinousSearch            bool                        `json:"continousSearch,omitempty"` //default: false
	ContinousSearchByServiceId types.ContSearchByServiceId `json:"contSearchByServiceId,omitzero"`
	CoordinateType             consts.CoordinateType       `json:"coordinateType,omitempty"`  //default: EPSG_4326
	SchedulesBefore            int32                       `json:"schedulesBefore,omitempty"` //default: 0
	SchedulesAfter             int32                       `json:"schedulesAfter,omitempty"`  //default: 0
	TariffInfoSelector         []types.TariffInfoSelector  `json:"tariffInfoSelector,omitempty"`
	Returnreduced              bool                        `json:"returnReduced,omitempty"`        //default: false
	ReturnPartialTickets       bool                        `json:"returnPartialTickets,omitempty"` //default: true
	Realtime                   consts.RealTimeType         `json:"realtime,omitempty"`             //default: AUTO
	IntermediateStops          bool                        `json:"intermediateStops,omitempty"`    //default: false
	UseStationPosition         bool                        `json:"useStationPosition,omitempty"`   //default: true
	ForcedStart                types.SDName                `json:"forcedStart,omitzero"`
	ForcedDest                 types.SDName                `json:"forcedDest,omitzero"`
	ToStartBy                  consts.TransportMode        `json:"toStartBy,omitempty"`
	ToDestinationBy            consts.TransportMode        `json:"toDestBy,omitempty"`
	ReturnContentSearchData    bool                        `json:"returnContSearchData,omitempty"`
	UseBikeAndRide             bool                        `json:"useBikeAndRide,omitempty"` //defaut: false
}

type IndividualRouteRequest struct {
	BaseRequest
	Starts       []types.SDName                 `json:"starts"`
	Destinations []types.SDName                 `json:"dests"`
	MaxLength    int32                          `json:"maxLength,omitempty"`
	MaxResults   int32                          `json:"maxResults,omitempty"`
	Type         consts.CoordinateType          `json:"type,omitempty"`        //default: EPSG_4326
	ServiceType  consts.TransportMode           `json:"serviceType,omitempty"` //default: FOOTPATH
	Profile      consts.BicycleTransportProfile `json:"profile,omitempty"`     //default: FOOT_NORMAL
	Speed        string                         `json:"speed,omitempty"`       //default: NORMAL
}

type AnnouncementRequest struct {
	BaseRequest
	Names                 []string        `json:"names,omitempty"`
	TimeRange             types.TimeRange `json:"timeRange,omitzero"`
	Full                  bool            `json:"full,omitempty"` //default: false
	FilterPlanned         string          `json:"filterPlanned,omitempty"`
	ShowBroadcastRelevant bool            `json:"showBroadcastRelevant,omitempty"` //default: false
}

type DLRequest struct {
	BaseRequest
	Station                   types.SDName                `json:"station,omitzero"`
	Stations                  []types.SDName              `json:"stations,omitempty"`
	Time                      types.GTITime               `json:"time"`
	MaxList                   int32                       `json:"maxList,omitempty"`
	MaxTimeOffset             int32                       `json:"maxTimeOffset,omitempty"`             //defauly: 120
	AllStationsInChangingNode bool                        `json:"allStationsInChangingNode,omitempty"` //default: true
	UseRealTime               bool                        `json:"useRealtime,omitempty"`               //default: false
	ReturnFilters             bool                        `json:"returnFilters,omitempty"`             //default: false
	Filter                    types.DLFilterEntry         `json:"filter,omitzero"`
	ServiceTypes              consts.TransportServiceType `json:"serviceTypes,omitempty"`
	Departure                 bool                        `json:"departure,omitempty"` //default: true
}

type DCRequest struct {
	BaseRequest
	LineKey        string                            `json:"lineKey"`
	Station        types.SDName                      `json:"station"`
	Time           string                            `json:"time"`
	Direction      string                            `json:"direction,omitempty"`
	Origin         string                            `json:"origin,omitempty"`
	ServiceID      int32                             `json:"serviceId,omitempty"`      //default: -1
	Segments       consts.DepartureCourseSegmentType `json:"segments,omitempty"`       //default: ALL
	ShowPath       bool                              `json:"showPath,omitempty"`       //default: false
	CoordinateType consts.CoordinateType             `json:"coordinateType,omitempty"` //default: EPSG_4326
}

type PCRequest struct {
	BaseRequest
	PostalCode int32 `json:"postalCode,omitempty"`
}

type CNRequest struct {
	BaseRequest
	Name            types.SDName          `json:"theName"`
	MaxList         int32                 `json:"maxList,omitempty"`
	MaxDistance     int32                 `json:"maxDistance,omitempty"`
	CoordinateType  consts.CoordinateType `json:"coordinateType,omitempty"`  //default: EPSG_4326
	TariffDetails   bool                  `json:"tariffDetails,omitempty"`   //default: false
	AllowTypeSwitch bool                  `json:"allowTypeSwitch,omitempty"` //default: true
}
