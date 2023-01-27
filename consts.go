package geofox

const (
	defaultScheme   = "https"
	defaultHostname = "gti.geofox.de"
	defaultBasePath = "/gti/public"
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
