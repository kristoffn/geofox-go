package geofox

type BaseResponse struct {
	ReturnCode   GeofoxReturnCode `json:"returnCode,omitempty"`
	ErrorText    string           `json:"errorText,omitempty"`
	ErrorDevInfo string           `json:"errorDevInfo,omitempty"`
}

type InitResponse struct {
	BaseResponse
	BeginOfService string     `json:"beginOfService"`
	EndOfService   string     `json:"endOfService"`
	ID             string     `json:"id"`
	DataID         string     `json:"dataId"`
	BuildDate      string     `json:"buildDate"`
	BuildTime      string     `json:"buildTime"`
	BuildText      string     `json:"buildText"`
	Version        string     `json:"version"`
	Properties     []Property `json:"properties"`
}

type LSResponse struct {
	BaseResponse
	DataReleaseID string             `json:"dataReleaseID,omitempty"`
	Stations      []StationListEntry `json:"stations,omitempty"`
}

type LLResponse struct {
	BaseResponse
	DataReleaseID string          `json:"dataReleaseID"`
	Lines         []LineListEntry `json:"lines"`
}

type TariffResponse struct {
	BaseResponse
	TariffInfos []TariffInfo `json:"tariffInfos"`
}

type CNResponse struct {
	BaseResponse
	Results []RegionalSDName `json:"results"`
}

type RegionalSDName struct {
	Name                  string        `json:"name"`
	City                  string        `json:"city"`
	CombinedName          string        `json:"combinedName"`
	ID                    string        `json:"id"`
	GlobalID              string        `json:"globalId"`
	Type                  SDType        `json:"type"`
	Coordinate            Coordniate    `json:"coordinate"`
	Layer                 int32         `json:"layer"`
	TariffDetails         TariffDetails `json:"tariffDetails"`
	ServiceTypes          []string      `json:"serviceTypes"`
	HasStationInformation bool          `json:"hasStationInformation"`
	Distance              int32         `json:"distance"`
	Time                  int32         `json:"time"`
}

type TariffInfo struct {
	TariffName    string             `json:"tariffName"`
	TariffRegions []TariffRegionInfo `json:"tariffRegions"`
	RegionTexts   []string           `json:"regionTexts"`
	ExtraFareType ExtraFareType      `json:"extraFareType"`
	TicketInfos   []TicketInfo       `json:"ticketInfos"`
	TicketRemarks string             `json:"ticketRemarks"`
}

type TariffRegionInfo struct {
	RegionType   TariffRegionType   `json:"regionType"`
	Alternatives []TariffRegionList `json:"alternatives"`
}

type TariffRegionList struct {
	Regions []string `json:"regions"`
}

type TicketInfo struct {
	TariffKindID          int32            `json:"tariffKindID"`
	TariffKindLabel       string           `json:"tariffKindLabel"`
	TariffLevelId         int32            `json:"tariffLevelID"`
	TariffLevelLabel      string           `json:"tariffLevelLabel"`
	TariffGroupID         int32            `json:"tariffGroupID"`
	TariffGroupLabel      string           `json:"tariffGroupLabel"`
	ViaPathID             int              `json:"viaPathId"`
	BasePrice             float64          `json:"basePrice"`
	ExtraFarePrice        float64          `json:"extraFarePrice"`
	ReducedBasePrice      float64          `json:"reducedBasePrice"`
	ReducedExtraFarePrice float64          `json:"reducedExtraFarePrice"`
	Currency              string           `json:"currency"`
	RegionType            TariffRegionType `json:"regionType"`
	NotRecommended        bool             `json:"notRecommended"`
	ShopLinkRegular       string           `json:"shopLinkRegular"`
	ShopLinkExtraFare     string           `json:"shopLinkExtraFare"`
	StartStationID        string           `json:"startStationId"`
	EndStationID          string           `json:"endStationId"`
}

type GTITime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

type ScheduleElementLight struct {
	DepartureStationID string `json:"departureStationId"`
	ArrivalStationID   string `json:"arrivalStationId"`
	LineID             string `json:"lineId"`
}

type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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

type LineListEntry struct {
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	CarrierNameShort string             `json:"carrierNameShort"`
	CarrierNameLong  string             `json:"carrierNameLong"`
	Sublines         []SublineListEntry `json:"sublines"`
	Exists           bool               `json:"exists"`
	Type             ServiceType        `json:"type"`
}

type SDName struct {
	Name                  string        `json:"name,omitempty"`
	City                  string        `json:"city,omitempty"`
	CombinedName          string        `json:"combinedName,omitempty"`
	ID                    string        `json:"id,omitempty"`
	GlobalID              string        `json:"globalId,omitempty"`
	Type                  SDType        `json:"type,omitempty"`
	Coordinate            Coordniate    `json:"coordinate,omitempty"`
	Layer                 int32         `json:"layer,omitempty"`
	TariffDetails         TariffDetails `json:"tariffDetails,omitempty"`
	SerivceTypes          []string      `json:"serviceTypes,omitempty"`
	HasStationInformation bool          `json:"hasStationInformation,omitempty"`
}

type TariffDetails struct {
	InnerCity       bool     `json:"innerCity"`
	City            bool     `json:"city"`
	CityTraffic     bool     `json:"cityTraffic"`
	Gratis          bool     `json:"gratis"`
	GreaterArea     bool     `json:"greaterArea"`
	SHVillageID     int32    `json:"shVillageId"`
	SHTariffZone    int32    `json:"shTariffZone"`
	TariffZones     []int32  `json:"tariffZones"`
	Regions         []int32  `json:"regions"`
	Counties        []string `json:"counties"`
	Rings           []string `json:"rings"`
	FareStage       bool     `json:"fareStage"`
	FareStageNumber int32    `json:"fareStageNumber"`
	TariffNames     []string `json:"tariffNames"`
	UniqueValues    bool     `json:"uniqueValues"`
}

type TariffOptimizerTicket struct {
	tariffKindId     int32      `json:"tariffKindId,omitempty"`
	tariffKindLabel  string     `json:"tariffKindLabel,omitempty"`
	tariffLevelId    int32      `json:"tariffLevelId,omitempty"`
	tariffLevelLabel string     `json:"tariffLevelLabel,omitempty"`
	tariffRegions    []string   `json:"tariffRegions"`
	regionType       RegionType `json:"regionType"`
	count            int32      `json:"count,omitempty"`
	extraFare        bool       `json:"extraFare,omitempty"`
	personType       PersonType `json:"personType"`
	centPrice        int32      `json:"centPrice,omitempty"`
}

type SingleTicketOptimizerRequestRoute struct {
	Trip                      SingleTicketOptimizerRequestTrip `json:"trip"`
	Departure                 string                           `json:"departure"`
	Arrival                   string                           `json:"arrival"`
	TariffRegions             TariffOptimizerRegions           `json:"tariffRegions"`
	SingleTicketTariffLevelID int32                            `json:"singleTicketTariffLevelId,omitempty"`
	ExtraFare                 ExtraFareType                    `json:"extraFareType"`
}

type SingleTicketOptimizerRequestTrip struct {
	Start       SingleTicketOptimizerRequestStation `json:"start"`
	Destination SingleTicketOptimizerRequestStation `json:"destination"`
	Line        SingleTicketOptimizerRequestLine    `json:"line"`
	VehicleType string                              `json:"vehicleType"`
}

type SingleTicketOptimizerRequestStation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SingleTicketOptimizerRequestLine struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TariffOptimizerRegions struct {
	Zones     []TariffRegions `json:"zones"`
	Rings     []TariffRegions `json:"rings"`
	Countries []TariffRegions `json:"counties"`
}

type TariffRegions struct {
	Regions []string `json:"regions"`
}

type BoundingBox struct {
	LowerLeft  Coordniate `json:"lowerLeft"`
	UpperRight Coordniate `json:"upperRight"`
}

type Penalty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ContSearchByServiceId struct {
	ServiceID                   int32   `json:"serviceId,omitempty"`
	LineKey                     string  `json:"lineKey"`
	PlannedDepartureArrivalTime GTITime `json:"plannedDepArrTime"`
	AdditionalOffset            int32   `json:"additionalOffset,omitempty"`
}

type TariffInfoSelector struct {
	Tariff        string  `json:"tariff"`        //default: HVV
	TariffRegions bool    `json:"tariffRegions"` //default: true
	Kinds         []int32 `json:"kinds"`
	Groups        []int32 `json:"groups"`
	Blacklist     bool    `json:"blacklist"` //default: false
}

type TimeRange struct {
	Begin string `json:"begin,omitempty"`
	End   string `json:"end,omitempty"`
}

type DLFilterEntry struct {
	ServiceID   string   `json:"serviceID,omitempty"`
	StationIDs  []string `json:"stationIDs,omitempty"`
	Label       string   `json:"label,omitempty"`
	ServiceName string   `json:"serviceName,omitempty"`
}
