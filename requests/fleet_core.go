package requests

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type ListShipsEndpoint struct {
	Pagination
}

func (e *ListShipsEndpoint) Method() string {
	return http.MethodGet
}

func (e *ListShipsEndpoint) Path() string {
	return "/my/ships"
}

func (e *ListShipsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ListShipsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ListShipsEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &ListShipsEndpoint{}
var _ WithQuery = &ListShipsEndpoint{}


// Will accept any string as ShipType, but prefer to use enum instead
type PurchaseShipEndpoint struct {
	WaypointSymbol string `json:"waypointSymbol"`
	ShipType string `json:"shipType"`
}

func (e *PurchaseShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *PurchaseShipEndpoint) Path() string {
	return "/my/ships"
}

func (e *PurchaseShipEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *PurchaseShipEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *PurchaseShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &PurchaseShipEndpoint{}


type GetShipEndpoint struct {
	ShipSymbol string
}

func (e *GetShipEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s", e.ShipSymbol)
}

func (e *GetShipEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetShipEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetShipEndpoint{}


type CreateChartEndpoint struct {
	ShipSymbol string
}

func (e *CreateChartEndpoint) Method() string {
	return http.MethodPost
}

func (e *CreateChartEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/chart", e.ShipSymbol)
}

func (e *CreateChartEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *CreateChartEndpoint) Headers() map[string]string {
	return nil
}

func (e *CreateChartEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &CreateChartEndpoint{}


type NegotiateContractEndpoint struct {
	ShipSymbol string
}

func (e *NegotiateContractEndpoint) Method() string {
	return http.MethodPost
}

func (e *NegotiateContractEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/negotiate/contract", e.ShipSymbol)
}

func (e *NegotiateContractEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *NegotiateContractEndpoint) Headers() map[string]string {
	return nil
}

func (e *NegotiateContractEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &NegotiateContractEndpoint{}


type GetCooldownEndpoint struct {
	ShipSymbol string
}

func (e *GetCooldownEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetCooldownEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/cooldown", e.ShipSymbol)
}

func (e *GetCooldownEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetCooldownEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetCooldownEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetCooldownEndpoint{}


type ScrapShipEndpoint struct {
	ShipSymbol string
}

func (e *ScrapShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *ScrapShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/scrap", e.ShipSymbol)
}

func (e *ScrapShipEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ScrapShipEndpoint) Headers() map[string]string {
	return nil
}

func (e *ScrapShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &ScrapShipEndpoint{}


type GetScrapValueEndpoint struct {
	ShipSymbol string
}

func (e *GetScrapValueEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetScrapValueEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/scrap", e.ShipSymbol)
}

func (e *GetScrapValueEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetScrapValueEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetScrapValueEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetScrapValueEndpoint{}


type RefuelShipEndpoint struct {
	ShipSymbol string `json:"-"`

	ShipFuelUnits int `json:"units"`
	FromCargo bool `json:"fromCargo"`
}

func (e *RefuelShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *RefuelShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/refuel", e.ShipSymbol)
}

func (e *RefuelShipEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *RefuelShipEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *RefuelShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &RefuelShipEndpoint{}


type RepairShipEndpoint struct {
	ShipSymbol string
}

func (e *RepairShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *RepairShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/repair", e.ShipSymbol)
}

func (e *RepairShipEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *RepairShipEndpoint) Headers() map[string]string {
	return nil
}

func (e *RepairShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &RepairShipEndpoint{}


type GetRepairCostEndpoint struct {
	ShipSymbol string
}

func (e *GetRepairCostEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetRepairCostEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/repair", e.ShipSymbol)
}

func (e *GetRepairCostEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetRepairCostEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetRepairCostEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetRepairCostEndpoint{}
