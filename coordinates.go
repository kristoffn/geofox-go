package geofox

type CoordinateType string

const (
	CoordinateTypeEPSG4326  CoordinateType = "EPSG_4326"
	CoordinateTypeEPSG31466 CoordinateType = "EPSG_31466"
	CoordinateTypeEPSG31467 CoordinateType = "EPSG_31467"
	CoordinateTypeEPSG31468 CoordinateType = "EPSG_31468"
	CoordinateTypeEPSG31469 CoordinateType = "EPSG_31469"
)

type Coordniate struct {
	X    float64        `json:"x"`
	Y    float64        `json:"y"`
	Type CoordinateType `json:"type"`
}
