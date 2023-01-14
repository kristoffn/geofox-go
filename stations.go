package geofox

type ListStationsRequest struct {
	DataReleaseID    string
	CoordinateType   CoordinateType
	ModificationType ModificationType
	FilterEqvivalent bool
}

type Station struct {
	ID           string
	Name         string
	City         string
	CombinedName string
	Shortcuts    []string
	Aliases      []string
	VehicleTypes []VehicleType
	Coordinate   Coordniate
	Exists       bool
}

type ListStationsResponse struct {
	DataReleaseID string
	Stations      []Station
}
