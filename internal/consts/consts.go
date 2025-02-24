package consts

const (
	DefaultScheme   = "https"
	DefaultHostname = "gti.geofox.de"
	DefaultBasePath = "/gti/public"
)

type CoordinateType string

const (
	EPSG4326  CoordinateType = "EPSG_4326"
	EPSG31466 CoordinateType = "EPSG_31466"
	EPSG31467 CoordinateType = "EPSG_31467"
	EPSG31468 CoordinateType = "EPSG_31468"
	EPSG31469 CoordinateType = "EPSG_31469"
)

type Language string

const (
	LangEnglish = "en"
	LangGerman  = "de"
)

type ModificationType string

const (
	ModTypeMain     ModificationType = "MAIN"
	ModTypePosition ModificationType = "POSITION"
	ModTypeSequence ModificationType = "SEQUENCE"
)

type VehicleType string

const (
	VTypeRegionalBus  VehicleType = "REGIONALBUS"
	VTypeMetroBus     VehicleType = "METROBUS"
	VTypeNachtBus     VehicleType = "NACHTBUS"
	VTypeSchnellBus   VehicleType = "SCHNELLBUS"
	VTypeXpressBus    VehicleType = "XPRESSBUS"
	VTypeAST          VehicleType = "AST"
	VTypeSchiff       VehicleType = "SCHIFF"
	VTypeUBahn        VehicleType = "U_BAHN"
	VTypeSBahn        VehicleType = "S_BAHN"
	VTypeABahn        VehicleType = "A_BAHN"
	VTypeRegionalbahn VehicleType = "R_BAHN"
	VTypeFernbahn     VehicleType = "F_BAHN"
	VTypeEILBus       VehicleType = "EILBUS"
)

type Filter string

const (
	NoFilter        Filter = "NO_FILTER"
	FilterHVVListed Filter = "HVV_LISTED"
)

type SimpleType string

const (
	STBus                SimpleType = "BUS"
	STTrain              SimpleType = "TRAIN"
	STShip               SimpleType = "SHIP"
	STFootpath           SimpleType = "FOOTPATH"
	STBicycle            SimpleType = "BICYCLE"
	STAirplane           SimpleType = "AIRPLANE"
	STChange             SimpleType = "CHANGE"
	STChangeSamePlatform SimpleType = "CHANGE_SAME_PLATFORM"
)

type ExtraFareType string

const (
	NoExtraFare       ExtraFareType = "NO"
	ExtraFarePossible ExtraFareType = "POSSIBLE"
	ExtraFareRequired ExtraFareType = "REQUIRED"
)

type TariffRegionType string

const (
	TariffRegionZone         TariffRegionType = "ZONE"
	TariffRegionGHZone       TariffRegionType = "GH_ZONE"
	TariffRegionRing         TariffRegionType = "RING"
	TariffRegionCounty       TariffRegionType = "COUNTY"
	TariffRegionGH           TariffRegionType = "GH"
	TariffRegionNet          TariffRegionType = "NET"
	TariffRegionZG           TariffRegionType = "ZG"
	TariffRegionStadtverkehr TariffRegionType = "STADTVERKEHR"
)

type SDType string

const (
	SDUnknown     SDType = "UNKNOWN"
	SDStation     SDType = "STATION"
	SDAddress     SDType = "ADDRESS"
	SDPOI         SDType = "POI"
	SDCoodrdinate SDType = "COORDINATE"
)

type PersonType string

const (
	PersonTypeAll        PersonType = "ALL"
	PersonTypeAdult      PersonType = "ADULT"
	PersonTypeElderly    PersonType = "ELDERLY"
	PersonTypeApprentice PersonType = "APPRENTICE"
	PersonTypePupil      PersonType = "PUPIL"
	PersonTypeStudent    PersonType = "STUDENT"
	PersonTypeChild      PersonType = "CHILD"
)

type RegionType string

const (
	RegionTypeRing   RegionType = "RING"
	RegionTypeZone   RegionType = "ZONE"
	RegionTypeCounty RegionType = "COUNTY"
	RegionTypeGhZone RegionType = "GH_ZONE"
)

type RealTimeType string

const (
	RTPlandata RealTimeType = "PLANDATA"
	RT         RealTimeType = "REALTIME"
	RTAuto     RealTimeType = "AUTO"
)

type TransportMode string

const (
	TMBus                 TransportMode = "BUS"
	TMTrain               TransportMode = "TRAIN"
	TMShip                TransportMode = "SHIP"
	TMFootpath            TransportMode = "FOOTPATH"
	TMBicycle             TransportMode = "BICYCLE"
	TMAirplan             TransportMode = "AIRPLANE"
	TMChange              TransportMode = "CHANGE"
	TMChangeSamePlatform  TransportMode = "CHANGE_SAME_PLATFORM"
	TMActivityBikeAndRide TransportMode = "ACTIVITY_BIKE_AND_RIDE"
)

type BicycleTransportProfile string

const (
	BTPNormal     BicycleTransportProfile = "BICYCLE_NORMAL"
	BTPRacing     BicycleTransportProfile = "BICYCLE_RACING"
	BTPQuietRoads BicycleTransportProfile = "BICYCLE_QUIET_ROADS"
	BTPMainRoads  BicycleTransportProfile = "BICYCLE_MAIN_ROADS"
	BTPBadWeather BicycleTransportProfile = "BICYCLE_BAD_WEATHER"
	BTPFootNormal BicycleTransportProfile = "FOOT_NORMAL"
)

type TransportServiceType string

const (
	TSZug        TransportServiceType = "ZUG"
	TSUBahn      TransportServiceType = "UBAHN"
	TSSBahn      TransportServiceType = "SBAHN"
	TSAKN        TransportServiceType = "AKN"
	TSRBahn      TransportServiceType = "RBAHN"
	TSFernbahn   TransportServiceType = "FERNBAHN"
	TSBus        TransportServiceType = "BUS"
	TSStadtbus   TransportServiceType = "STADTBUS"
	TSMetrobus   TransportServiceType = "METROBUS"
	TSSchnellbus TransportServiceType = "SCHNELLBUS"
	TSNachtbus   TransportServiceType = "NACHTBUS"
	TSXpressbus  TransportServiceType = "XPRESSBUS"
	TSEilbus     TransportServiceType = "EILBUS"
	TSAST        TransportServiceType = "AST"
	TSFaehre     TransportServiceType = "FAEHRE"
)

type DepartureCourseSegmentType string

const (
	DCSBefore DepartureCourseSegmentType = "BEFORE"
	DCSAfter  DepartureCourseSegmentType = "AFTER"
	DCSAll    DepartureCourseSegmentType = "ALL"
)

type DayType string

const (
	DTWeekday DayType = "WEEKDAY"
	DTWeekend DayType = "WEEKEND"
)

type TicketClassType string

const (
	TCNone    TicketClassType = "NONE"
	TCSecond  TicketClassType = "SECOND"
	TCFirst   TicketClassType = "FIRST"
	TCSchnell TicketClassType = "SCHNELL"
)

type DiscountType string

const (
	DiscontNone   DiscountType = "NONE"
	DiscontOnline DiscountType = "ONLINE"
	DiscontSocial DiscountType = "SOCIAL"
)

type TicketType string

const (
	TicketTypeOccasional TicketType = "OCCASIONAL_TICKET"
	TicketTypeSeasonal   TicketType = "SEASON_TICKET"
)

type ElevatorButtonType string

const (
	ElevatorButtonBraille ElevatorButtonType = "BRAILLE"
	ElevatorButtonAcustic ElevatorButtonType = "ACUSTIC"
	ElevatorButtonCombi   ElevatorButtonType = "COMBI"
	ElevatorButtonUnkown  ElevatorButtonType = "UNKNOWN"
)

type ElevatorStateType string

const (
	ElevatorStateReady      ElevatorStateType = "READY"
	ElevatorStateOutOfOrder ElevatorStateType = "OUTOFORDER"
	ElevatorStateUnkown     ElevatorStateType = "UNKNOWN"
)

type LocationType string

const (
	LocationSingleLine        LocationType = "SINGLE_LINE"
	LocationAllLinesOfCarrier LocationType = "ALL_LINES_OF_CARRIER"
	LocationCompleteNetwork   LocationType = "COMPLETE_NET"
)
