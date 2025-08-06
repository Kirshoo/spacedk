package spacedk

import "time"

type Response[T any] struct {
	Data T
	Meta Metadata
}

type Metadata struct {
	Total int `json:"total"`
	Page int `json:"page"`
	Limit int `json:"limit"`
}

type Agent struct {
	Symbol string `json:"symbol"`
	HQ WaypointSymbol `json:"headquarters"`
	Credits Currency `json:"credits"`
	StartingFaction FactionSymbol `json:"startingFaction"`
	ShipCount int `json:"shipCount"`
}

type OwnAgent struct {
	Agent
	AccountId string `json:"accountId"`
}

type AgentEvent struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Message string `json:"message"`
	Data any `json:"data"`
	CreatedAt time.Time `json:"createdAt"`
}

type Contract struct {
	Id string `json:"id"`
	FactionSymbol FactionSymbol `json:"factionSymbol"`
	Type ContractType `json:"type"`
	IsAccepted bool `json:"accepted"`
	IsFulfilled bool `json:"fulfilled"`
	AcceptanceDeadline time.Time `json:"deadlineToAccept"`

	Terms ContractTermsInfo `json:"terms"`
}

type ContractTermsInfo struct {
	Deadline time.Time `json:"deadline"`

	Payment PaymentTermsInfo `json:"payment"`
	Deliver []Freight `json:"deliver"`
}

type PaymentTermsInfo struct {
	OnAccept Currency `json:"onAccepted"`
	OnFulfill Currency `json:"onFulfilled"`
}

type Freight struct {
	ItemSymbol ItemSymbol `json:"tradeSymbol"`
	DestinationSymbol WaypointSymbol `json:"destinationSymbol"`
	Required int `json:"unitsRequired"`
	Fulfilled int `json:"unitsFulfilled"`
}

type Faction struct {
	Symbol FactionSymbol `json:"symbol"`
	Name string `json:"name"`
	Description string `json:"description"`
	HQ WaypointSymbol `json:"headquarters"`
	IsRecruiting bool `json:"isRecruiting"`

	Traits []FactionTrait `json:"traits"`
}

type FactionTrait struct {
	Symbol FactionTraitSymbol `json:"symbol"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type FactionReputation struct {
	Symbol FactionTraitSymbol `json:"symbol"`
	Reputation int `json:"reputation"`
}

type FactionRef struct {
	Symbol FactionSymbol `json:"symbol"`
}

type Ship struct {
	Symbol string `json:"symbol"`

	Registration ShipRegistrationInfo `json:"registration"`
	Navigation NavigationInfo `json:"nav"`
	Crew CrewDetailedInfo `json:"crew"`
	Frame FrameInfo `json:"frame"`
	Reactor ReactorInfo `json:"reactor"`
	Engine EngineInfo `json:"engine"`
	Modules []Module `json:"modules"`
	Mounts []Mount `json:"mounts"`
	Cargo CargoInfo `json:"cargo"`
	Fuel FuelInfo `json:"fuel"`
	Cooldown CooldownInfo `json:"cooldown"`
}

type ShipScan struct {
	Symbol string `json:"symbol"`

	Registration ShipRegistrationInfo `json:"registration"`
	Navigation NavigationInfo `json:"nav"`
	Frame FrameInfo `json:"frame"`
	Reactor ReactorInfo `json:"reactor"`
	Engine EngineInfo `json:"engine"`
	Mounts []Mount `json:"mounts"`
}

type ShipRegistrationInfo struct {
	// TODO: Maybe an enum?
	Name string `json:"name"`

	FactionSymbol FactionSymbol `json:"factionSymbol"`
	Role ShipRole `json:"role"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type NavigationInfo struct {
	SystemSymbol SystemSymbol `json:"systemSymbol"`
	WaypointSymbol WaypointSymbol `json:"waypointSymbol"`
	Status NavStatus `json:"status"`
	FlightMode FlightMode `json:"flightMode"`

	Route RouteInfo `json:"route"`
}

type RouteInfo struct {
	DeparturedAt time.Time `json:"departureTime"`
	ArivesAt time.Time `json:"arrivalTime"`

	Destination RoutePoint `json:"destination"`
	Origin RoutePoint `json:"origin"`
}

type RoutePoint struct {
	SystemSymbol SystemSymbol `json:"systemSymbol"`
	WaypointSymbol WaypointSymbol `json:"symbol"`
	Type WaypointType `json:"type"`
	Coordinates
}

type CrewInfo struct {
	Current int `json:"current"`
	Required int `json:"required"`
}

type CrewDetailedInfo struct {
	CrewInfo
	Capacity int `json:"capacity"`
	Rotation CrewRotation `json:"rotation"`
	Morale int `json:"morale"`
	Wages int `json:"wages"`
}

// T defines, which symbol enum to use
type ShipPart[T any] struct {
	Symbol T `json:"symbol"`
	Name string `json:"name"`
	Description string `json:"description"`

	Requirements UsageRequirement `json:"requirements"`
}

type UsageRequirement struct {
	Power int `json:"power"`
	Crew int `json:"crew"`
	Slots int `json:"slots"`
}

type ShipPartDetailed[T any] struct {
	ShipPart[T]
	Quality int `json:"quality"`

	// Between 0 and 1
	Condition float32 `json:"condition"`
	// Between 0 and 1
	Integrity float32 `json:"integrity"`
}

type FrameInfo struct {
	ShipPartDetailed[FrameType]

	// Never negative
	ModuleSlots int `json:"moduleSlots"`
	// Never negative
	MountPoints int `json:"mountingPoints"`
	// Never negative
	FuelCapacity int `json:"fuelCapacity"`
}

type ReactorInfo struct {
	ShipPartDetailed[ReactorType]

	// At least 1
	PowerOutput int `json:"powerOutput"`
}

type EngineInfo struct {
	ShipPartDetailed[EngineType]

	// At least 1
	Speed int `json:"speed"`
}

type Module struct {
	ShipPart[ModuleType]

	// May not always be set, but never negative
	Capacity int `json:"capacity"`
	// May not always be set, but never negative
	Range int `json:"range"`
}

type ItemUnits struct {
	ItemSymbol ItemSymbol `json:"tradeSymbol"`
	Units int `json:"units"`
}

type Mount struct {
	ShipPart[MountType]

	// Never negative
	Strength int `json:"strength"`
	Deposits []RawMaterial `json:"deposits"`
}

type ExtractionInfo struct {
	ShipSymbol string `json:"shipSymbol"`
	Yield ItemUnits `json:"yield"`
}

type Survey struct {
	Signature string `json:"signature"`
	Symbol WaypointSymbol `json:"symbol"`
	Deposits []SurveyDeposit `json:"deposits"`
	Expiration time.Time `json:"expiration"`
	Size SurveySize `json:"size"`
}

type SurveyDeposit struct {
	Symbol ItemSymbol `json:"symbol"`
}

type CargoInfo struct {
	// Never negative
	Capacity int `json:"capacity"`
	// Never negative
	Stored int `json:"units"`
	Inventory []CargoItem `json:"inventory"`
}

type CargoItem struct {
	Symbol ItemSymbol `json:"symbol"`
	Name string `json:"name"`
	Description string `json:"description"`
	// Always at least 1
	Units int `json:"units"`
}

type FuelInfo struct {
	// Never negative
	Current int `json:"current"`
	// Never negative
	Capacity int `json:"capacity"`

	Consumed ConsumedFuel `json:"consumed"`
}

type ConsumedFuel struct {
	// Never negative
	Amount int `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type CooldownInfo struct {
	ShipSymbol string `json:"shipSymbol"`
	TotalDuration int `json:"totalSeconds"`
	RemainingDuration int `json:"remainingSeconds"`
	ExpiresAt time.Time `json:"expiration"`
}

type ShipConditionEvent struct {
	Symbol ConditionEventSymbol `json:"symbol"`
	Component ComponentSymbol `json:"component"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type System struct {
	SectorSymbol SectorSymbol `json:"sectorSymbol"`
	Symbol SystemSymbol `json:"symbol"`
	Type SystemType `json:"type"`
	Coordinates
}

type SystemDetailed struct {
	System
	Name string `json:"name"`
	Constellation string `json:"constellation"`

	Waypoints []Waypoint `json:"waypoints"`
	Factions []FactionRef `json:"factions"`
}

type SystemScan struct {
	System
	Distance int `json:"distance"`
}

type Waypoint struct {
	Symbol WaypointSymbol `json:"symbol"`
	Type WaypointType `json:"type"`
	Coordinates
	Orbitals []WaypointOrbital `json:"orbitals"`
	Orbits string `json:"orbits"`
}

func (w *Waypoint) IsOrbital() bool {
	return w.Orbits != ""
}

type WaypointOrbital struct {
	Symbol WaypointSymbol `json:"symbol"`
}

type WaypointDetailed struct {
	Waypoint
	SystemSymbol SystemSymbol `json:"systemSymbol"`
	IsUnderConstruction bool `json:"isUnderConstruction"`
	
	Faction FactionRef `json:"faction"`
	Traits []WaypointTrait `json:"traits"`
	Modifiers []WaypointModifier `json:"modifier"`
	Chart WaypointChart `json:"chart"`
}

type WaypointScan struct {
	Waypoint
	SystemSymbol SystemSymbol `json:"systemSymbol"`
	Faction FactionRef `json:"faction"`
	Traits []WaypointTrait `json:"traits"`
	Chart WaypointChart `json:"chart"`
}

func (w *WaypointDetailed) IsCharted() bool {
	return w.Chart.SubmittedAt.IsZero()
}

type WaypointTrait struct {
	Symbol WaypointTraitSymbol `json:"symbol"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type WaypointModifier struct {
	Symbol WaypointModifierSymbol `json:"symbol"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type WaypointChart struct {
	WaypointSymbol WaypointSymbol `json:"waypointSymbol"`
	SubmittedBy string `json:"submittedBy"`
	SubmittedAt time.Time `json:"submittedOn"`
}

type GenericTransaction struct {
	WaypointSymbol WaypointSymbol `json:"waypointSymbol"`
	ShipSymbol string `json:"shipSymbol"`
	TotalPrice int `json:"totalPrice"`
	Timestamp time.Time `json:"timestamp"`
}

type ConstructionSite struct {
	Symbol WaypointSymbol `json:"symbol"`
	IsCompleted bool `json:"isComplete"`

	Materials []ConstructionMaterial `json:"materials"`
}

type ConstructionMaterial struct {
	ItemSymbol ItemSymbol `json:"tradeSymbol"`
	Required int `json:"required"`
	Fulfilled int `json:"fulfilled"`
}

type Marketplace struct {
	Symbol WaypointSymbol `json:"symbol"`

	Exports []TradeItem `json:"exports"`
	Imports []TradeItem `json:"imports"`
	Exchange []TradeItem `json:"exchange"`
	Transactions []MarketTransactionEntry `json:"transaction"`
	Stock []MarketTradeItem `json:"tradeGoods"`
}

type TradeItem struct {
	Symbol ItemSymbol `json:"symbol"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type MarketTransactionEntry struct {
	WaypointSymbol WaypointSymbol `json:"waypointSymbol"`
	ShipSymbol string `json:"shipSymbol"`
	ItemSymbol ItemSymbol `json:"tradeSymbol"`
	Type MarketTransactionType `json:"type"`
	Units int `json:"units"`
	PricePerUnit Currency `json:"pricePerUnit"`
	TotalPrice Currency `json:"totalPrice"`
	Timestamp time.Time `json:"timestamp"`
}

type MarketTradeItem struct {
	Symbol ItemSymbol `json:"symbol"`
	Type MarketTradeType `json:"type"`
	TradeVolume int `json:"tradeVolume"`
	Supply SupplyLevel `json:"supply"`
	Activity ActivityLevel `json:"activity"`
	PurchasePrice Currency `json:"purchasePrice"`
	SellPrice Currency `json:"sellPrice"`
}

type JumpGate struct {
	Symbol WaypointSymbol `json:"symbol"`

	Connections []WaypointSymbol `json:"connections"`
}

type Shipyard struct {
	Symbol WaypointSymbol `json:"symbol"`
	ModificationsFee Currency `json:"modificationsFee"`

	AvailableTypes []AvailableShipType `json:"shipTypes"`
	Transactions []ShipyardTransactionEntry `json:"transactions"`
	Stock []ShipyardShip `json:"ships"`
}

type AvailableShipType struct {
	Type ShipType `json:"type"`
}

type ShipyardTransactionEntry struct {
	WaypointSymbol WaypointSymbol `json:"waypointSymbol"`
	ShipType ShipType `json:"shipType"`
	Price Currency `json:"price"`
	AgentSymbol string `json:"agentSymbol"`
	Timestamp time.Time `json:"timestamp"`
}

type ShipyardShip struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Type ShipType `json:"type"`
	Supply SupplyLevel `json:"supply"`
	Activity ActivityLevel `json:"activity"`
	PurchasePrice Currency `json:"purchasePrice"`

	Frame FrameInfo `json:"frame"`
	Reactor ReactorInfo `json:"reactor"`
	Engine EngineInfo `json:"engine"`
	Modules []Module `json:"modules"`
	Mounts []Mount `json:"mounts"`
	Crew CrewInfo `json:"crew"`
}

type ApiServer struct {
	Status string `json:"status"`
	Version string `json:"version"`
	ResetDate DateOnlyTime `json:"resetDate"`
	Description string `json:"description"`

	Stats ApiServerStats `json:"stats"`
	Health ApiServerHealth `json:"health"`
	Leaderboards ApiServerLeaderboards `json:"leaderboards"`
	ServerResets ApiServerResetInfo `json:"serverResets"`
	Announcements []Announcement `json:"announcements"`
	Links []Link `json:"links"`
}

type ApiServerStats struct {
	AccountCount int `json:"accounts"`
	AgentCount int `json:"agents"`
	ShipCount int `json:"ships"`
	SystemCount int `json:"systems"`
	WaypointCount int `json:"waypoints"`
}

type ApiServerHealth struct {
	MarketUpdatedAt time.Time `json:"lastMarketUpdate"`
}

type ApiServerLeaderboards struct {
	MostCredits []CreditsLeaderboardEntry `json:"mostCredits"`
	MostCharts []ChartsLeaderboardEntry `json:"mostCredits"`
}

type CreditsLeaderboardEntry struct {
	AgentSymbol string `json:"agentSymbol"`
	Credits Currency `json:"credits"`
}

type ChartsLeaderboardEntry struct {
	AgentSymbol string `json:"agentSymbol"`
	ChartCount int `json:"chartCount"`
}

type ApiServerResetInfo struct {
	Next time.Time `json:"next"`
	Frequency string `json:"frequency"`
}

type Announcement struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

type Link struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type ErrorCodesInfo struct {
	ErrorCodes []ErrorCode `json:"errorCodes"`
}

type ErrorCode struct {
	Code int `json:"code"`
	Name string `json:""name`
}
