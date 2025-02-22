package geofox

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

type TicketListTicketInfos struct {
	TariffKindID         int32                     `json:"tariffKindID,omitempty"`
	TariffKindLabel      string                    `json:"tariffKindLabel"`
	TariffLevelID        int32                     `json:"tariffLevelID,omitempty"`
	TariffLevelLabel     string                    `json:"tariffLevelLabel"`
	TariffGroupID        int32                     `json:"tariffGroupID,omitempty"`
	TariffGroupLabel     string                    `json:"tariffGroupLabel,omitempty"`
	RegionType           TariffRegionType          `json:"regionType,omitempty"`
	SelectableRegions    int32                     `json:"selectableRegions,omitempty"`    //default: 0
	RequiredStartStation bool                      `json:"requiredStartStation,omitempty"` //default: false
	PersonInfos          []PersonInfo              `json:"personInfos,omitempty"`
	ValidityPeriods      []ValidityPeriod          `json:"validityPeriods,omitempty"`
	Variants             []TicketListTicketVariant `json:"variants,omitempty"`
}

type PersonInfo struct {
	PersonType  PersonType `json:"personType,omitempty"`
	PersonCount int32      `json:"personCount,omitempty"`
}

type ValidityPeriod struct {
	Day            DayType      `json:"day"`
	TimeValidities []TimePeriod `json:"timeValidities,omitempty"`
}

type TimePeriod struct {
	Begin string `json:"begin"`
	End   string `json:"end"`
}

type TicketListTicketVariant struct {
	TicketID      int32           `json:"ticketId,omitempty"`
	KaNummer      int32           `json:"kaNummer,omitempty"`
	Price         float64         `json:"price,omitempty"`
	Currency      string          `json:"currency,omitempty"` //default: EUR
	TicketClass   TicketClassType `json:"ticketClass"`
	Discont       DiscountType    `json:"discount"`
	ValidityBegin string          `json:"validityBegin"`
	ValidityEnd   string          `json:"validityEnd"`
}

type TariffZone struct {
	Zone       string   `json:"zone"`
	Ring       string   `json:"ring"`
	Neighbours []string `json:"neighbours"`
}

type TariffKind struct {
	ID                 int32      `json:"id,omitempty"`
	Label              string     `json:"label"`
	RequiresPersonType bool       `json:"requiresPersonType,omitempty"` //default: false
	TicketType         TicketType `json:"ticketType,omitempty"`
	LevelCombinations  []int32    `json:"levelCombinations,omitempty"`
}

type TariffLevel struct {
	ID                 int32              `json:"id,omitempty"`
	Label              string             `json:"label"`
	RequiredRegionType RequiredRegionType `json:"requiredRegionType"`
}

type RequiredRegionType struct {
	Type  TariffRegionType `json:"type"`
	Count int32            `json:"count,omitempty"`
}

type TariffCounty struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type Journey struct {
	JourneyID   string        `json:"journeyID"`
	Line        Service       `json:"line"`
	VehicleType VehicleType   `json:"vehicleType"`
	Realtime    bool          `json:"realtime,omitempty"`
	Segments    []Pathsegment `json:"segments,omitempty"`
}

type Service struct {
	Name             string      `json:"name"`
	Direction        string      `json:"direction,omitempty"`
	DirectionID      int32       `json:"directionId,omitempty"`
	Origin           string      `json:"origin,omitempty"`
	Type             ServiceType `json:"type"`
	ID               string      `json:"id,omitempty"`
	DLID             string      `json:"dlid,omitempty"`
	CarrierNameShort string      `json:"carrierNameShort,omitempty"`
	CarrierNameLong  string      `json:"carrierNameLong,omitempty"`
}

type Pathsegment struct {
	StartStopPointKey string         `json:"startStopPointKey"`
	EndStopPointKey   string         `json:"endStopPointKey"`
	StartStationName  string         `json:"startStationName"`
	StartStationKey   string         `json:"startStationKey"`
	StartDateTime     int64          `json:"startDateTime,omitempty"`
	EndStationName    string         `json:"endStationName"`
	EndStationKey     string         `json:"endStationKey"`
	EndDateTime       int64          `json:"endDateTime,omitempty"`
	Track             VehicleMapPath `json:"track"`
	Destination       string         `json:"destination"`
	RealtimeDelay     int32          `json:"realtimeDelay,omitempty"`
	IsFirst           bool           `json:"isFirst,omitempty"`
	IsLast            bool           `json:"isLast,omitempty"`
}

type VehicleMapPath struct {
	Track          []float64      `json:"track,omitempty"`
	CoordinateType CoordinateType `json:"coordinateType"`
}

type PartialStation struct {
	Lines          []string   `json:"lines,omitempty"`
	StationOutline string     `json:"stationOutline,omitempty"`
	Elevators      []Elevator `json:"elevators,omitempty"`
}

type Elevator struct {
	Lines        []string           `json:"lines,omitempty"`
	Label        string             `json:"label,omitempty"`
	CabinWidth   int32              `json:"cabinWidth,omitempty"`
	CabinLength  int32              `json:"cabinLength,omitempty"`
	DoorWidth    int32              `json:"doorWidth,omitempty"`
	Description  string             `json:"description,omitempty"`
	ElevatorType string             `json:"elevatorType,omitempty"`
	ButtonType   ElevatorButtonType `json:"buttonType,omitempty"`
	State        ElevatorStateType  `json:"state,omitempty"`
	Cause        string             `json:"cause,omitempty"`
}

type Schedule struct {
	RouteID              int32                 `json:"routeId,omitempty"`
	Start                SDName                `json:"start"`
	Destination          SDName                `json:"dest"`
	Time                 int32                 `json:"time,omitempty"`
	FootPathTime         int32                 `json:"footpathTime,omitempty"`
	PlannedDepartureTime string                `json:"plannedDepartureTime,omitempty"`
	RealDepartureTime    string                `json:"realDepartureTime,omitempty"`
	PlannedArrivalTime   string                `json:"plannedArrivalTime,omitempty"`
	RealArrivalTime      string                `json:"realArrivalTime,omitempty"`
	Tickets              []Ticket              `json:"tickets,omitempty"`
	TariffInfos          []TariffInfo          `json:"tariffInfos,omitempty"`
	ScheduleElements     []ScheduleElement     `json:"scheduleElements,omitempty"`
	ContSearchBefore     ContSearchByServiceId `json:"contSearchBefore,omitzero"`
	ContSearchAfter      ContSearchByServiceId `json:"contSearchAfter,omitzero"`
}

type Ticket struct {
	Price         float64 `json:"price,omitempty"`
	ReducedPrice  float64 `json:"reducedPrice,omitempty"`
	Currency      string  `json:"currency,omitempty"` //default: EUR
	Type          string  `json:"type"`
	Level         string  `json:"level"`
	Tariff        string  `json:"tariff"`
	Range         string  `json:"range,omitempty"`
	TicketRemarks string  `json:"ticketRemarks,omitempty"`
}

type ScheduleElement struct {
	DepartureStationID string   `json:"departureStationId"`
	ArrivalStationID   string   `json:"arrivalStationId"`
	IntermediateStops  []string `json:"intermediateStops,omitempty"`
	LineID             string   `json:"lineId"`
}

type IndividualTrack struct {
	Time   int32         `json:"time,omitempty"`
	Length int32         `json:"length,omitempty"`
	Type   TransportMode `json:"type"`
}

type IndividualRoute struct {
	Start       SDName        `json:"start"`
	Destination SDName        `json:"dest"`
	Path        Path          `json:"path,omitzero"`
	Paths       []Path        `json:"paths,omitempty"`
	Length      int32         `json:"length,omitempty"`
	Time        int32         `json:"time,omitempty"`
	ServiceType TransportMode `json:"serviceType"`
}

type Path struct {
	Track      []Coordniate `json:"track"`
	Attributes []string     `json:"attributes,omitempty"`
	Tags       []MapEntry   `json:"tags,omitempty"`
}

type MapEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Announcement struct {
	ID                string      `json:"id,omitempty"`
	Version           int32       `json:"version,omitempty"`
	Summary           string      `json:"summary,omitempty"`
	Locations         []Location  `json:"locations,omitempty"`
	Description       string      `json:"description"`
	Links             []Link      `json:"links,omitempty"`
	Publication       TimeRange   `json:"publication"`
	Validities        []TimeRange `json:"validities"`
	LastModified      string      `json:"lastModified"`
	Planned           bool        `json:"planned,omitempty"`
	Reason            string      `json:"reason,omitempty"`
	BroadcastRelevant bool        `json:"broadcastRelevant,omitempty"`
}

type Location struct {
	Type           LocationType `json:"type,omitempty"`
	Name           string       `json:"name,omitempty"`
	Line           Service      `json:"line,omitzero"`
	Begin          SDName       `json:"begin,omitzero"`
	End            SDName       `json:"end,omitzero"`
	BothDirections bool         `json:"bothDirections,omitempty"` //default: true
}

type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Departure struct {
	Line             Service     `json:"line,omitzero"`
	DirectionID      int32       `json:"directionId,omitempty"`
	TimeOffset       int32       `json:"timeOffset,omitempty"`
	Delay            int32       `json:"delay,omitempty"`
	Extra            bool        `json:"extra,omitempty"`     //default: false
	Cancelled        bool        `json:"cancelled,omitempty"` //default: false
	ServiceID        int32       `json:"serviceId,omitempty"`
	Station          SDName      `json:"station,omitzero"`
	Platform         string      `json:"platform,omitempty"`
	RealtimePlatform string      `json:"realtimePlatform,omitempty"`
	Atrributes       []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Title     string   `json:"title,omitempty"`
	IsPlanned bool     `json:"isPlanned,omitempty"`
	Value     string   `json:"value"`
	Types     []string `json:"types,omitempty"`
	ID        string   `json:"id,omitempty"`
}

type CourseElement struct {
	FromStation          SDName      `json:"fromStation"`
	FromPlatform         string      `json:"fromPlatform,omitempty"`
	FromRealtimePlatform string      `json:"fromRealtimePlatform,omitempty"`
	ToStation            SDName      `json:"toStation"`
	ToPlatform           string      `json:"toPlatform,omitempty"`
	ToRealtimePlatform   string      `json:"toRealtimePlatform,omitempty"`
	Model                string      `json:"model,omitempty"`
	DepartureTime        string      `json:"depTime"`
	ArrivalTime          string      `json:"arrTime"`
	DepartureDelay       int32       `json:"depDelay,omitempty"`
	ArrivalDelay         int32       `json:"arrDelay,omitempty"`
	FromExtra            bool        `json:"fromExtra,omitempty"`     //default: false
	FromCancelled        bool        `json:"fromCancelled,omitempty"` //default: false
	ToExtra              bool        `json:"toExtra,omitempty"`       //default: false
	ToCancelled          bool        `json:"toCancelled,omitempty"`   //default: false
	Attributes           []Attribute `json:"attributes,omitempty"`
	Path                 Path        `json:"path,omitzero"`
}
