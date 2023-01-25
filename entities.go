package geofox

type BaseRequest struct {
	Language   string `json:"language,omitempty"`
	Version    int32  `json:"version,omitempty"`
	FilterType Filter `json:"filterType,omitempty"`
}

type InitRequest struct {
	BaseRequest
}

type InitResponse struct {
	ReturnCode     string `json:"returnCode"`
	BeginOfService string `json:"beginOfService"`
	EndOfService   string `json:"endOfService"`
	ID             string `json:"id"`
	DataID         string `json:"dataId"`
	BuildDate      string `json:"buildDate"`
	BuildTime      string `json:"buildTime"`
	BuildText      string `json:"buildText"`
}

type LSRequest struct {
	BaseRequest
	DataReleaseID     string             `json:"dataReleaseID"`
	ModificationTypes []ModificationType `json:"modificationTypes"`
	CoordinateType    CoordinateType     `json:"coordinateType,omitempty"`
	FilterEquivalent  bool               `json:"filterEquivalent,omitempty"`
}

type LSResponse struct {
	DataReleaseID string             `json:"dataReleaseID,omitempty"`
	Stations      []StationListEntry `json:"stations,omitempty"`
}

type StationListEntry struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	City         string        `json:"city"`
	CombinedName string        `json:"combinedName"`
	Shortcuts    []string      `json:"shortcuts"`
	Aliasses     []string      `json:"aliasses"`
	VehicleTypes []VehicleType `json:"vehicleTypes"`
	Coordniate   Coordniate    `json:"coordinate"`
	Exists       bool          `json:"exists"`
}

type Coordniate struct {
	X    float64        `json:"x"`
	Y    float64        `json:"y"`
	Type CoordinateType `json:"type"`
}

type LLRequest struct {
	BaseRequest
	WithSublines      bool               `json:"withSublines"`
	DataReleaseID     string             `json:"dataReleaseID"`
	ModificationTypes []ModificationType `json:"modificationTypes"`
}

type LLResponse struct {
	ReturnCode    GeofoxReturnCode `json:"returnCode"`
	ErrorText     string           `json:"errorText"`
	ErrorDevInfo  string           `json:"errorDevInfo"`
	DataReleaseID string           `json:"dataReleaseID"`
	Lines         []LineListEntry  `json:"lines"`
}

type LineListEntry struct {
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	CarrierNameShort string             `json:"carrierNameShort"`
	CarrierNameLong  string             `json:"carrierNameLong"`
	Sublines         []SublineListEntry `json:"sublines"`
	Exists           bool               `json:"exists"`
	Type             ServiceType        `json:"type"`
}

type SublineListEntry struct {
	SublineNumber   string         `json:"sublineNumber"`
	VehicleType     VehicleType    `json:"vehicleType"`
	StationSequence []StationLight `json:"stationSequence"`
}

type ServiceType struct {
	SimpleType SimpleType `json:"simpleType"`
	ShortInfo  string     `json:"shortInfo"`
	LongInfo   string     `json:"longInfo"`
	Model      string     `json:"model"`
}

type StationLight struct {
	ID   string `json:"id"`
	Name string `json:"json"`
}
