package geofox

type BaseRequest struct {
	Language   string `json:"language,omitempty"`   // default: de
	Version    int32  `json:"version,omitempty"`    // default: 1
	FilterType Filter `json:"filterType,omitempty"` // default: NO_FILTER
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
	WithReturnJourney bool                              `json:"withReturnJourney,omitempty"`
	NumberOfAdults    int32                             `json:"numberOfAdults,omitempty"`
	NumberOfChildren  int32                             `json:"numberOfChildren,omitempty"`
	Tickets           []TariffOptimizerTicket           `json:"tickets,omitempty"`
	Route             SingleTicketOptimizerRequestRoute `json:"route"`
}

type LSRequest struct {
	BaseRequest
	DataReleaseID     string             `json:"dataReleaseID"`
	ModificationTypes []ModificationType `json:"modificationTypes"`
	CoordinateType    CoordinateType     `json:"coordinateType,omitempty"`
	FilterEquivalent  bool               `json:"filterEquivalent,omitempty"`
}

type LLRequest struct {
	BaseRequest
	WithSublines      bool               `json:"withSublines"`
	DataReleaseID     string             `json:"dataReleaseID"`
	ModificationTypes []ModificationType `json:"modificationTypes"`
}

type InitRequest struct {
	BaseRequest
}

type VehicleMapRequest struct {
	BaseRequest
	BoundingBox    BoundingBox    `json:"boundingBox"`
	PeriodBegin    int64          `json:"periodBegin,omitempty"`
	PeriodEnd      int64          `json:"periodEnd,omitempty"`
	WithoutCoords  bool           `json:"withoutCoords,omitempty"`
	CoordinateType CoordinateType `json:"coordinateType,omitempty"` //default: EPSG_4326
	VehicleTypes   []VehicleType  `json:"vehicleTypes,omitempty"`
	RealTime       bool           `json:"realtime,omitempty"`
}

type TrackCoordinatesRequest struct {
	BaseRequest
	CoordinateType CoordinateType `json:"coordinateType,omitempty"` //default: EPSG_4326
	StopPointKeys  []string       `json:"stopPointKeys"`
}

type TariffRequest struct {
	BaseRequest
	ScheduleElements     []ScheduleElementLight `json:"scheduleElements"`
	Departure            GTITime                `json:"departure"`
	Arrival              GTITime                `json:"arrival"`
	ReturnReduced        bool                   `json:"returnReduced"`
	ReturnPartialTickets bool                   `json:"returnPartialTickets"`
}

type StationInformationRequest struct {
	BaseRequest
	Station SDName `json:"station"`
}

type GRRequest struct {
	BaseRequest
	Start                      SDName                `json:"start"`
	Destination                SDName                `json:"dest"`
	Via                        SDName                `json:"via,omitempty"`
	Time                       GTITime               `json:"time"`
	TimeIsDeparture            bool                  `json:"timeIsDeparture,omitempty"`
	WithPaths                  bool                  `json:"withPaths,omitempty"`
	NumberOfSchedules          int32                 `json:"numberOfSchedules,omitempty"`
	Penalties                  []Penalty             `json:"penalties,omitempty"`
	TariffDetails              bool                  `json:"tariffDetails,omitempty"`   //default: false
	ContinousSearch            bool                  `json:"continousSearch,omitempty"` //default: false
	ContinousSearchByServiceId ContSearchByServiceId `json:"contSearchByServiceId,omitempty"`
	CoordinateType             CoordinateType        `json:"coordinateType,omitempty"`  //default: EPSG_4326
	SchedulesBefore            int32                 `json:"schedulesBefore,omitempty"` //default: 0
	SchedulesAfter             int32                 `json:"schedulesAfter,omitempty"`  //default: 0
	TariffInfoSelector         []TariffInfoSelector  `json:"tariffInfoSelector,omitempty"`
	Returnreduced              bool                  `json:"returnReduced,omitempty"`        //default: false
	ReturnPartialTickets       bool                  `json:"returnPartialTickets,omitempty"` //default: true
	Realtime                   RealTimeType          `json:"realtime,omitempty"`             //default: AUTO
	IntermediateStops          bool                  `json:"intermediateStops,omitempty"`    //default: false
	UseStationPosition         bool                  `json:"useStationPosition,omitempty"`   //default: true
	ForcedStart                SDName                `json:"forcedStart,omitempty"`
	ForcedDest                 SDName                `json:"forcedDest,omitempty"`
	ToStartBy                  TransportMode         `json:"toStartBy,omitempty"`
	ToDestinationBy            TransportMode         `json:"toDestBy,omitempty"`
	ReturnContentSearchData    bool                  `json:"returnContSearchData,omitempty"`
	UseBikeAndRide             bool                  `json:"useBikeAndRide,omitempty"` //defaut: false
}

type IndividualRouteRequest struct {
	BaseRequest
	Starts       []SDName                `json:"starts"`
	Destinations []SDName                `json:"dests"`
	MaxLength    int32                   `json:"maxLength,omitempty"`
	MaxResults   int32                   `json:"maxResults,omitempty"`
	Type         CoordinateType          `json:"type,omitempty"`        //default: EPSG_4326
	ServiceType  TransportMode           `json:"serviceType,omitempty"` //default: FOOTPATH
	Profile      BicycleTransportProfile `json:"profile,omitempty"`     //default: FOOT_NORMAL
	Speed        string                  `json:"speed,omitempty"`       //default: NORMAL
}

type AnnouncementRequest struct {
	BaseRequest
	Names                 []string  `json:"names,omitempty"`
	TimeRange             TimeRange `json:"timeRange,omitempty"`
	Full                  bool      `json:"full,omitempty"` //default: false
	FilterPlanned         string    `json:"filterPlanned,omitempty"`
	ShowBroadcastRelevant bool      `json:"showBroadcastRelevant,omitempty"` //default: false
}

type DLRequest struct {
	BaseRequest
	Station                   SDName               `json:"station,omitempty"`
	Stations                  []SDName             `json:"stations,omitempty"`
	Time                      GTITime              `json:"time"`
	MaxList                   int32                `json:"maxList,omitempty"`
	MaxTimeOffset             int32                `json:"maxTimeOffset,omitempty"`             //defauly: 120
	AllStationsInChangingNode bool                 `json:"allStationsInChangingNode,omitempty"` //default: true
	UseRealTime               bool                 `json:"useRealtime,omitempty"`               //default: false
	ReturnFilters             bool                 `json:"returnFilters,omitempty"`             //default: false
	Filter                    DLFilterEntry        `json:"filter,omitempty"`
	ServiceTypes              TransportServiceType `json:"serviceTypes,omitempty"`
	Departure                 bool                 `json:"departure,omitempty"` //default: true
}

type DCRequest struct {
	BaseRequest
	LineKey        string                     `json:"lineKey"`
	Station        SDName                     `json:"station"`
	Time           string                     `json:"time"`
	Direction      string                     `json:"direction,omitempty"`
	Origin         string                     `json:"origin,omitempty"`
	ServiceID      int32                      `json:"serviceId,omitempty"`      //default: -1
	Segments       DepartureCourseSegmentType `json:"segments,omitempty"`       //default: ALL
	ShowPath       bool                       `json:"showPath,omitempty"`       //default: false
	CoordinateType CoordinateType             `json:"coordinateType,omitempty"` //default: EPSG_4326
}

type PCRequest struct {
	BaseRequest
	PostalCode int32 `json:"postalCode,omitempty"`
}

type CNRequest struct {
	BaseRequest
	Name            SDName         `json:"theName"`
	MaxList         int32          `json:"maxList"`
	MaxDistance     int32          `json:"maxDistance"`
	CoordinateType  CoordinateType `json:"coordinateType"`
	TariffDetails   bool           `json:"tariffDetails"`
	AllowTypeSwitch bool           `json:"allowTypeSwitch"`
}
