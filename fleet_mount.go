package spacedk

import "github.com/Kirshoo/spacedk/requests"

func (s *FleetService) GetMounts(shipSymbol string) ([]Mount, error) {
	req := &requests.GetMountsEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[[]Mount](s.client, req)
	if err != nil {
		return nil, err
	}

	return reply.Data, nil
}

type AgentMountsCargoTransactionPayload struct {
	Agent OwnAgent `json:"agent"`
	Mounts []Mount `json:"mounts"`
	Cargo CargoInfo `json:"cargo"`
	Transaction GenericTransaction `json:"transaction"`
}

func (s *FleetService) InstallMount(shipSymbol string, mountSymbol MountType) (*OwnAgent, []Mount, *CargoInfo, *GenericTransaction, error) {
	req := &requests.InstallMountEndpoint{
		ShipSymbol: shipSymbol,
		MountSymbol: string(mountSymbol),
	}

	reply, err := Do[AgentMountsCargoTransactionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &reply.Data.Agent, reply.Data.Mounts, &reply.Data.Cargo, &reply.Data.Transaction, nil
}

func (s *FleetService) RemoveMount(shipSymbol string, mountSymbol MountType) (*OwnAgent, []Mount, *CargoInfo, *GenericTransaction, error) {
	req := &requests.RemoveMountEndpoint{
		ShipSymbol: shipSymbol,
		MountSymbol: string(mountSymbol),
	}

	reply, err := Do[AgentMountsCargoTransactionPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &reply.Data.Agent, reply.Data.Mounts, &reply.Data.Cargo, &reply.Data.Transaction, nil
}

type ExtractionCooldownCargoModifiersEventsPayload struct {
	Extraction ExtractionInfo `json:"extraction"`
	Cooldown CooldownInfo `json:"cooldown"`
	Cargo CargoInfo `json:"cargo"`
	Modifiers []WaypointModifier `json:"modifiers"`
	Events []ShipConditionEvent `json:"event"`
}

func (s *FleetService) ExtractResource(shipSymbol string) (*ExtractionInfo, *CooldownInfo, *CargoInfo, []WaypointModifier, []ShipConditionEvent, error) {
	req := &requests.ExtractResourcesEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[ExtractionCooldownCargoModifiersEventsPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return &reply.Data.Extraction, &reply.Data.Cooldown, &reply.Data.Cargo, reply.Data.Modifiers, reply.Data.Events, nil
}

func (s *FleetService) ExtractResourceWithSurvey(shipSymbol string, survey *Survey) (*ExtractionInfo, *CooldownInfo, *CargoInfo, []WaypointModifier, []ShipConditionEvent, error) {
	req := &requests.ExtractResourcesWithSurveyEndpoint{
		ShipSymbol: shipSymbol,
		Signature: survey.Signature,
		Symbol: survey.Symbol.String(),
		Expiration: survey.Expiration,
		Size: string(survey.Size),
	}

	// Because endpoint requires an array of structs
	// containing strings and similar field of survey struct
	// has ItemSymbol, we need to perform the transformation
	// and include each deposit one by one
	for _, deposit := range survey.Deposits {
		transformed := struct {
			Symbol string `json:"symbol"`
		}{Symbol: string(deposit.Symbol)}

		req.Deposits = append(req.Deposits, transformed)
	}

	reply, err := Do[ExtractionCooldownCargoModifiersEventsPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return &reply.Data.Extraction, &reply.Data.Cooldown, &reply.Data.Cargo, reply.Data.Modifiers, reply.Data.Events, nil
}

func (s *FleetService) SiphonResource(shipSymbol string) (*ExtractionInfo, *CooldownInfo, *CargoInfo, []WaypointModifier, []ShipConditionEvent, error) {
	req := &requests.ExtractResourcesEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[ExtractionCooldownCargoModifiersEventsPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return &reply.Data.Extraction, &reply.Data.Cooldown, &reply.Data.Cargo, reply.Data.Modifiers, reply.Data.Events, nil
}

type CooldownSystemsPayload struct {
	Cooldown CooldownInfo `json:"cooldown"`
	Systems []SystemScan `json:"systems"`
}

func (s *FleetService) ScanSystem(shipSymbol string) (*CooldownInfo, []SystemScan, error) {
	req := &requests.ScanSystemsEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[CooldownSystemsPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Cooldown, reply.Data.Systems, nil
}

type CooldownWaypointsPayload struct {
	Cooldown CooldownInfo `json:"cooldown"`
	Waypoints []WaypointScan `json:"waypoints"`
}

func (s *FleetService) ScanWaypoints(shipSymbol string) (*CooldownInfo, []WaypointScan, error) {
	req := &requests.ScanWaypointsEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[CooldownWaypointsPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Cooldown, reply.Data.Waypoints, nil
}

type CooldownShipsPayload struct {
	Cooldown CooldownInfo `json:"cooldown"`
	Ships []ShipScan `json:"ships"`
}

func (s *FleetService) ScanShips(shipSymbol string) (*CooldownInfo, []ShipScan, error) {
	req := &requests.ScanShipsEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[CooldownShipsPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Cooldown, reply.Data.Ships, nil
}

type CooldownSurveysPayload struct {
	Cooldown CooldownInfo `json:"cooldown"`
	Surveys []Survey `json:"surveys"`
}

func (s *FleetService) Survey(shipSymbol string) (*CooldownInfo, []Survey, error) {
	req := &requests.CreateSurveyEndpoint{ShipSymbol: shipSymbol}

	reply, err := Do[CooldownSurveysPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Cooldown, reply.Data.Surveys, nil
}
