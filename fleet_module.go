package spacedk

import "github.com/Kirshoo/spacedk/requests"

func (s *FleetService) GetModules(shipSymbol string) ([]Module, error) {
	req := &requests.GetModulesEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[[]Module](s.client, req)
	if err != nil {
		return nil, err
	}

	return reply.Data, nil
}

type AgentModulesCargoTransactionPayload struct {
	Agent OwnAgent `json:"agent"`
	Modules []Module `json:"modules"`
	Cargo CargoInfo `json:"cargo"`
	Transaction GenericTransaction `json:"transaction"`
}

func (s *FleetService) InstallModule(shipSymbol string, moduleSymbol ModuleType) (*OwnAgent, []Module, *CargoInfo, *GenericTransaction, error) {
	req := &requests.InstallModuleEndpoint{
		ShipSymbol: shipSymbol,
		ModuleSymbol: string(moduleSymbol),
	}

	reply, err := Do[AgentModulesCargoTransactionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &reply.Data.Agent, reply.Data.Modules, &reply.Data.Cargo, &reply.Data.Transaction, nil
}

func (s *FleetService) RemoveModule(shipSymbol string, moduleSymbol ModuleType) (*OwnAgent, []Module, *CargoInfo, *GenericTransaction, error) {
	req := &requests.RemoveModuleEndpoint{
		ShipSymbol: shipSymbol,
		ModuleSymbol: string(moduleSymbol),
	}

	reply, err := Do[AgentModulesCargoTransactionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &reply.Data.Agent, reply.Data.Modules, &reply.Data.Cargo, &reply.Data.Transaction, nil
}

type CargoCooldownProducedConsumedPayload struct {
	Cargo CargoInfo `json:"cargo"`
	Cooldown CooldownInfo `json:"cooldown"`
	Produced []ItemUnits `json:"produced"`
	Consumed []ItemUnits `json:"consumed"`
}

// First ItemUnits returns produced items and their units
// Second - consumed items and their units
func (s *FleetService) Refine(shipSymbol string, productSymbol ProductItem) (*CargoInfo, *CooldownInfo, []ItemUnits, []ItemUnits, error) {
	req := &requests.RefineResourceEndpoint{
		ShipSymbol: shipSymbol,
		ProductSymbol: string(productSymbol),
	}

	reply, err := Do[CargoCooldownProducedConsumedPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &reply.Data.Cargo, &reply.Data.Cooldown, reply.Data.Produced, reply.Data.Consumed, nil
}
