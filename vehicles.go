package geofox

type VehicleType string

const (
	VehicleTypeRegionalBus  VehicleType = "REGIONALBUS"
	VehicleTypeMetroBus     VehicleType = "METROBUS"
	VehicleTypeNachtBus     VehicleType = "NACHTBUS"
	VehicleTypeSchnellBus   VehicleType = "SCHNELLBUS"
	VehicleTypeXpressBus    VehicleType = "XPRESSBUS"
	VehicleTypeAST          VehicleType = "AST"
	VehicleTypeSchiff       VehicleType = "SCHIFF"
	VehicleTypeUBahn        VehicleType = "U_BAHN"
	VehicleTypeSBahn        VehicleType = "S_BAHN"
	VehicleTypeABahn        VehicleType = "A_BAHN"
	VehicleTypeRegionalbahn VehicleType = "R_BAHN"
	VehicleTypeFernbahn     VehicleType = "F_BAHN"
)
