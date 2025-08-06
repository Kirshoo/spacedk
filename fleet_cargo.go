package spacedk

import "github.com/Kirshoo/spacedk/requests"

func (s *FleetService) GetCargo(shipSymbol string) (*CargoInfo, error) {
	req := &requests.GetCargoEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[CargoInfo](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

type CargoMarketTransactionAgentPayload struct {
	Cargo CargoInfo `json:"cargo"`
	Transaction MarketTransactionEntry `json:"transaction"`
	Agent OwnAgent `json:"agent"`
}

func (s *FleetService) PurchaseCargo(shipSymbol string, itemSymbol ItemSymbol, amount int) (*CargoInfo, *MarketTransactionEntry, *OwnAgent, error) {
	req := &requests.PurchaseCargoEndpoint{
		ShipSymbol: shipSymbol,
		ItemSymbol: string(itemSymbol),
		Units: amount,
	}

	reply, err := Do[CargoMarketTransactionAgentPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, err
	}

	return &reply.Data.Cargo, &reply.Data.Transaction, &reply.Data.Agent, nil
}

func (s *FleetService) SellCargo(shipSymbol string, itemSymbol ItemSymbol, amount int) (*CargoInfo, *MarketTransactionEntry, *OwnAgent, error) {
	req := &requests.SellCargoEndpoint{
		ShipSymbol: shipSymbol,
		ItemSymbol: string(itemSymbol),
		Units: amount,
	}

	reply, err := Do[CargoMarketTransactionAgentPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, err
	}

	return &reply.Data.Cargo, &reply.Data.Transaction, &reply.Data.Agent, nil
}

type CargoCargoPayload struct {
	MyCargo CargoInfo `json:"cargo"`
	TargetCargo CargoInfo `json:"targetCargo"`
}

// First CargoInfo corresponds to own cargo,
// Second to the cargo of the target ship
func (s *FleetService) TransferCargo(shipSymbol, targetSymbol string, itemSymbol ItemSymbol, amount int) (*CargoInfo, *CargoInfo, error) {
	req := &requests.TransferCargoEndpoint{
		ShipSymbol: shipSymbol,
		ItemSymbol: string(itemSymbol),
		Units: amount,
		TargetSymbol: targetSymbol,
	}

	reply, err := Do[CargoCargoPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.MyCargo, &reply.Data.TargetCargo, nil
}

func (s *FleetService) JettisonCargo(shipSymbol string, itemSymbol ItemSymbol, amount int) (*CargoInfo, error) {
	req := &requests.JettisonCargoEndpoint{
		ShipSymbol: shipSymbol,
		ItemSymbol: string(itemSymbol),
		Units: amount,
	}

	reply, err := Do[CargoInfo](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}
