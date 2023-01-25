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
)

type Filter string

const (
	NoFilter        Filter = "NO_FILTER"
	FilterHVVListed Filter = "HVV_LISTED"
)
