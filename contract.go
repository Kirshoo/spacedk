package spacedk

import (
	"github.com/Kirshoo/spacedk/requests"
)

type ContractService struct {
	client *Client
}

func NewContractService(c *Client) *ContractService {
	return &ContractService{
		client: c,
	}
}

func (s *ContractService) List(page, limit int) ([]Contract, *Metadata, error) {
	req := &requests.ListContractsEndpoint{
		Pagination: requests.Pagination{
			Page: page, 
			Limit: limit,
		},
	}

	reply, err := Do[[]Contract](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return reply.Data, &reply.Meta, nil
}

func (s *ContractService) Get(id string) (*Contract, error) {
	req := &requests.GetContractEndpoint{ContractId: id}

	reply, err := Do[Contract](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

type ContractAgentPayload struct {
	Contract Contract `json:"contract"`
	Agent Agent `json:"agent"`
}

func (s *ContractService) Accept(id string) (*Contract, *Agent, error) {
	req := &requests.AcceptContractEndpoint{ContractId: id}

	reply, err := Do[ContractAgentPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Contract, &reply.Data.Agent, nil
}

func (s *ContractService) Fulfill(id string) (*Contract, *Agent, error) {
	req := &requests.FulfillContractEndpoint{ContractId: id}

	reply, err := Do[ContractAgentPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Contract, &reply.Data.Agent, nil
}

type ContractCargoPayload struct {
	Contract Contract `json:"contract"`
	Cargo CargoInfo `json:"cargo"`
}

func (s *ContractService) Deliver(id, shipSymbol, itemSymbol string, amount int) (*Contract, *CargoInfo, error) {
	req := &requests.DeliverToContractEndpoint{
		ContractId: id,
		ShipSymbol: shipSymbol,
		ItemSymbol: itemSymbol,
		Units: amount,
	}

	reply, err := Do[ContractCargoPayload](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return &reply.Data.Contract, &reply.Data.Cargo, nil
}
