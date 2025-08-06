package spacedk

import (
	"errors"
	"github.com/Kirshoo/spacedk/requests"
)

type FleetService struct {
	client *Client
}

func NewFleetService(c *Client) *FleetService {
	return &FleetService{
		client: c,
	}
}

func (s *FleetService) List(page, limit int) ([]Ship, *Metadata, error) {
	req := &requests.ListShipsEndpoint{
		Pagination: requests.Pagination{
			Page: page, 
			Limit: limit,
		},
	}

	reply, err := Do[[]Ship](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return reply.Data, &reply.Meta, nil
}

func (s *FleetService) Get(symbol string) (*Ship, error) {
	req := &requests.GetShipEndpoint{ShipSymbol: symbol}

	reply, err := Do[Ship](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

type ShipAgentTransactionPayload struct {
	Ship Ship `json:"ship"`
	Agent OwnAgent `json:"agent"`
	Transaction ShipyardTransactionEntry `json:"transaction"`
}

func (s *FleetService) PurchaseShip(shipType ShipType, waypoint WaypointSymbol) (*Ship, *OwnAgent, *ShipyardTransactionEntry, error) {
	req := &requests.PurchaseShipEndpoint{
		ShipType: string(shipType), 
		WaypointSymbol: waypoint.String(),
	}

	reply, err := Do[ShipAgentTransactionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, err
	}

	return &reply.Data.Ship, &reply.Data.Agent, &reply.Data.Transaction, nil
}

type ChartWaypointTransactionAgentPayload struct {
	Chart WaypointChart `json:"chart"`
	Waypoint WaypointDetailed `json:"waypoint"`
	Transaction GenericTransaction `json:"transaction"`
	Agent OwnAgent `json:"agent"`
}

func (s *FleetService) Chart(shipSymbol string) (*WaypointChart, *WaypointDetailed, *GenericTransaction, *OwnAgent, error) {
	req := &requests.CreateChartEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[ChartWaypointTransactionAgentPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &reply.Data.Chart, &reply.Data.Waypoint, &reply.Data.Transaction, &reply.Data.Agent, nil
}

func (s *FleetService) NegotiateContract(shipSymbol string) (*Contract, error) {
	req := &requests.NegotiateContractEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[Contract](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

func (s *FleetService) specialCooldownHandler(code int, body []byte, headers map[string][]string, v any) error {
	return NoCooldown
}

func (s *FleetService) GetCooldown(symbol string) (*CooldownInfo, bool, error) {
	req := &requests.GetCooldownEndpoint{ShipSymbol: symbol}

	reply, err := Do[CooldownInfo](s.client, req, 
		WithStatusHandler(204, s.specialCooldownHandler))
	if err != nil {
		if errors.Is(err, NoCooldown) {
			return nil, true, nil
		}

		return nil, false, err
	}

	return &reply.Data, false, nil
}

type AgentTransactionPayload struct {
	Agent OwnAgent `json:"agent"`
	Transaction GenericTransaction `json:"transaction"`
}

func (s *FleetService) ScrapShip(symbol string) (*OwnAgent, *GenericTransaction, error) {
	req := &requests.ScrapShipEndpoint{ShipSymbol: symbol}

	reply, err := Do[AgentTransactionPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Agent, &reply.Data.Transaction, nil
}

func (s *FleetService) GetScrapValue(symbol string) (*GenericTransaction, error) {
	req := &requests.GetScrapValueEndpoint{ShipSymbol: symbol}

	reply, err := Do[GenericTransaction](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

type AgentFuelCargoTransactionPayload struct {
	Agent OwnAgent `json:"agent"`
	Fuel FuelInfo `json:"fuel"`
	Cargo CargoInfo `json:"cargo"`
	Transaction MarketTransactionEntry `json:"transaction"`
}

func (s *FleetService) RefuelShip(symbol string, amount int, fromCargo bool) (*OwnAgent, *FuelInfo, *CargoInfo, *MarketTransactionEntry, error) {
	req := &requests.RefuelShipEndpoint{
		ShipSymbol: symbol,
		ShipFuelUnits: amount,
		FromCargo: fromCargo,
	}

	reply, err := Do[AgentFuelCargoTransactionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &reply.Data.Agent, &reply.Data.Fuel, &reply.Data.Cargo, &reply.Data.Transaction, nil
}

type AgentShipTransactionPayload struct {
	Agent OwnAgent `json:"agent"`
	Ship Ship `json:"ship"`
	Transaction GenericTransaction `json:"transaction"`
}

func (s *FleetService) RepairShip(symbol string) (*OwnAgent, *Ship, *GenericTransaction, error) {
	req := &requests.RepairShipEndpoint{ShipSymbol: symbol}

	reply, err := Do[AgentShipTransactionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, err
	}

	return &reply.Data.Agent, &reply.Data.Ship, &reply.Data.Transaction, nil
}

func (s *FleetService) GetRepairCost(symbol string) (*GenericTransaction, error) {
	req := &requests.GetRepairCostEndpoint{ShipSymbol: symbol}

	reply, err := Do[GenericTransaction](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}
