package spacedk

import "github.com/Kirshoo/spacedk/requests"

func (s *FleetService) GetShipNavigation(symbol string) (*NavigationInfo, error) {
	req := &requests.GetNavigationEndpoint{ShipSymbol: symbol}

	reply, err := Do[NavigationInfo](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

type NavigationPayload struct {
	Navigation NavigationInfo `json:"nav"`
}

func (s *FleetService) DockShip(symbol string) (*NavigationInfo, error) {
	req := &requests.DockShipEndpoint{ShipSymbol: symbol}

	reply, err := Do[NavigationPayload](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data.Navigation, nil
}

func (s *FleetService) OrbitShip(symbol string) (*NavigationInfo, error) {
	req := &requests.OrbitShipEndpoint{ShipSymbol: symbol}

	reply, err := Do[NavigationPayload](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data.Navigation, nil
}

type NavigationCooldownTransactionAgentPayload struct {
	Navigation NavigationInfo `json:"nav"`
	Cooldown CooldownInfo `json:"cooldown"`
	Transaction MarketTransactionEntry `json:"transaction"`
	Agent OwnAgent `json:"agent"`
}

func (s *FleetService) JumpShip(symbol string, destination WaypointSymbol) (*NavigationInfo, *CooldownInfo, *MarketTransactionEntry, *OwnAgent, error) {
	req := &requests.JumpShipEndpoint{
		ShipSymbol: symbol,
		DestinationSymbol: destination.String(),
	}

	reply, err := Do[NavigationCooldownTransactionAgentPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &reply.Data.Navigation, &reply.Data.Cooldown, &reply.Data.Transaction, &reply.Data.Agent, nil
}

type NavigationFuelConditionPayload struct {
	Navigation NavigationInfo `json:"nav"`
	Fuel FuelInfo `json:"fuel"`
	Events []ShipConditionEvent `json:"events"`
}

func (s *FleetService) NavigateShip(symbol string, destination WaypointSymbol) (*NavigationInfo, *FuelInfo, []ShipConditionEvent, error) {
	req := &requests.NavigateShipEndpoint{
		ShipSymbol: symbol,
		DestinationSymbol: destination.String(),
	}

	reply, err := Do[NavigationFuelConditionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, err
	}

	return &reply.Data.Navigation, &reply.Data.Fuel, reply.Data.Events, nil
}

func (s *FleetService) WarpShip(symbol string, destination WaypointSymbol) (*NavigationInfo, *FuelInfo, []ShipConditionEvent, error) {
	req := &requests.WarpShipEndpoint{
		ShipSymbol: symbol,
		DestinationSymbol: destination.String(),
	}

	reply, err := Do[NavigationFuelConditionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, err
	}

	return &reply.Data.Navigation, &reply.Data.Fuel, reply.Data.Events, nil
}

func (s *FleetService) ChangeShipNavigationMode(symbol string, mode FlightMode) (*NavigationInfo, *FuelInfo, []ShipConditionEvent, error) {
	req := &requests.ChangeNavigationModeEndpoint{
		ShipSymbol: symbol,
		FlightMode: string(mode),
	}

	reply, err := Do[NavigationFuelConditionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, err
	}

	return &reply.Data.Navigation, &reply.Data.Fuel, reply.Data.Events, nil
}
