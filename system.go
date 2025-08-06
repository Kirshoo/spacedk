package spacedk

import "github.com/Kirshoo/spacedk/requests"

type SystemService struct {
	client *Client
}

func NewSystemService(c *Client) *SystemService {
	return &SystemService{
		client: c,
	}
}

func (s *SystemService) ListSystems(page, limit int) ([]SystemDetailed, *Metadata, error) {
	req := &requests.ListSystemsEndpoint{
		Pagination: requests.Pagination{
			Page: page,
			Limit: limit,
		},
	}

	reply, err := Do[[]SystemDetailed](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return reply.Data, &reply.Meta, nil
}

func (s *SystemService) GetSystem(symbol SystemSymbol) (*SystemDetailed, error) {
	req := &requests.GetSystemEndpoint{SystemSymbol: symbol.String()}

	reply, err := Do[SystemDetailed](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

// TODO: Add trait and type optionals
func (s *SystemService) ListWaypoints(system SystemSymbol, page, limit int) ([]WaypointDetailed, *Metadata, error) {
	req := &requests.ListWaypointsEndpoint{
		SystemSymbol: system.String(),
		Pagination: requests.Pagination{
			Page: page,
			Limit: limit,
		},
	}

	reply, err := Do[[]WaypointDetailed](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return reply.Data, &reply.Meta, nil
}

func (s *SystemService) GetWaypoint(symbol WaypointSymbol) (*WaypointDetailed, error) {
	req := &requests.GetWaypointEndpoint{
		SystemSymbol: symbol.ToSystemSymbol().String(),
		WaypointSymbol: symbol.String(),
	}

	reply, err := Do[WaypointDetailed](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

func (s *SystemService) GetConstructionSite(symbol WaypointSymbol) (*ConstructionSite, error) {
	req := &requests.GetConstructionEndpoint{
		SystemSymbol: symbol.ToSystemSymbol().String(),
		WaypointSymbol: symbol.String(),
	}

	reply, err := Do[ConstructionSite](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

type ConstructionCargoPayload struct {
	Construction ConstructionSite `json:"construction"`
	Cargo CargoInfo `json:"cargo"`
}

func (s *SystemService) SupplyConstruction(waypoint WaypointSymbol, shipSymbol string, itemSymbol ItemSymbol, amount int) (*ConstructionSite, *CargoInfo, error) {
	req := &requests.SupplyConstructionEndpoint{
		SystemSymbol: waypoint.ToSystemSymbol().String(),
		WaypointSymbol: waypoint.String(),

		ShipSymbol: shipSymbol,
		ItemSymbol: string(itemSymbol),
		Units: amount,
	}

	reply, err := Do[ConstructionCargoPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Construction, &reply.Data.Cargo, nil
}

func (s *SystemService) GetMarket(symbol WaypointSymbol) (*Marketplace, error) {
	req := &requests.GetMarketEndpoint{
		SystemSymbol: symbol.ToSystemSymbol().String(),
		WaypointSymbol: symbol.String(),
	}

	reply, err := Do[Marketplace](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

func (s *SystemService) GetJumpGate(symbol WaypointSymbol) (*JumpGate, error) {
	req := &requests.GetJumpGateEndpoint{
		SystemSymbol: symbol.ToSystemSymbol().String(),
		WaypointSymbol: symbol.String(),
	}

	reply, err := Do[JumpGate](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

func (s *SystemService) GetShipyard(symbol WaypointSymbol) (*Shipyard, error) {
	req := &requests.GetShipyardEndpoint{
		SystemSymbol: symbol.ToSystemSymbol().String(),
		WaypointSymbol: symbol.String(),
	}

	reply, err := Do[Shipyard](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}
