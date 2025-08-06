package requests

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type ListContractsEndpoint struct {
	Pagination
}

func (e *ListContractsEndpoint) Method() string {
	return http.MethodGet
}

func (e *ListContractsEndpoint) Path() string {
	return "/my/contracts"
}

func (e *ListContractsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ListContractsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ListContractsEndpoint) IsTokenRequired() bool {
	return true
}

// Compile time check for correct implementation of interfaces
var _ ApiEndpoint = (*ListContractsEndpoint)(nil)
var _ WithQuery = (*ListContractsEndpoint)(nil)


type GetContractEndpoint struct {
	ContractId string
}

func (e *GetContractEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetContractEndpoint) Path() string {
	return fmt.Sprintf("/my/contracts/%s", e.ContractId)
}

func (e *GetContractEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetContractEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetContractEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = (*GetContractEndpoint)(nil)


type AcceptContractEndpoint struct {
	ContractId string
}

func (e *AcceptContractEndpoint) Method() string {
	return http.MethodPost
}

func (e *AcceptContractEndpoint) Path() string {
	return fmt.Sprintf("/my/contracts/%s/accept", e.ContractId)
}

func (e *AcceptContractEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *AcceptContractEndpoint) Headers() map[string]string {
	return nil
}

func (e *AcceptContractEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = (*AcceptContractEndpoint)(nil)


type FulfillContractEndpoint struct {
	ContractId string
}

func (e *FulfillContractEndpoint) Method() string {
	return http.MethodPost
}

func (e *FulfillContractEndpoint) Path() string {
	return fmt.Sprintf("/my/contracts/%s/fulfill", e.ContractId)
}

func (e *FulfillContractEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *FulfillContractEndpoint) Headers() map[string]string {
	return nil
}

func (e *FulfillContractEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = (*FulfillContractEndpoint)(nil)


type DeliverToContractEndpoint struct {
	ContractId string `json:"-"`
	ShipSymbol string `json:"shipSymbol"`
	ItemSymbol string `json:"tradeSymbol"`
	Units int `json:"units"`
}

func (e *DeliverToContractEndpoint) Method() string {
	return http.MethodPost
}

func (e *DeliverToContractEndpoint) Path() string {
	return fmt.Sprintf("/my/contracts/%s/deliver", e.ContractId)
}

func (e *DeliverToContractEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *DeliverToContractEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *DeliverToContractEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = (*DeliverToContractEndpoint)(nil)
