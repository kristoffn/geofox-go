package coordinates

type CoordinateType string

const (
	EPSG4326  CoordinateType = "EPSG_4326"
	EPSG31466 CoordinateType = "EPSG_31466"
	EPSG31467 CoordinateType = "EPSG_31467"
	EPSG31468 CoordinateType = "EPSG_31468"
	EPSG31469 CoordinateType = "EPSG_31469"
)

type Coordniate struct {
	X    float64        `json:"x"`
	Y    float64        `json:"y"`
	Type CoordinateType `json:"type"`
}
