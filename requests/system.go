package requests

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
)

type ListSystemsEndpoint struct {
	Pagination
}

func (e *ListSystemsEndpoint) Method() string {
	return http.MethodGet
}

func (e *ListSystemsEndpoint) Path() string {
	return "/systems"
}

func (e *ListSystemsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ListSystemsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ListSystemsEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*ListSystemsEndpoint)(nil)
var _ WithQuery = (*ListSystemsEndpoint)(nil)


type GetSystemEndpoint struct {
	SystemSymbol string
}

func (e *GetSystemEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetSystemEndpoint) Path() string {
	return fmt.Sprintf("/systems/%s", e.SystemSymbol)
}

func (e *GetSystemEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetSystemEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetSystemEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetSystemEndpoint)(nil)


type ListWaypointsEndpoint struct {
	SystemSymbol string

	Pagination
	Traits []string
}

func (e *ListWaypointsEndpoint) Method() string {
	return http.MethodGet
}

func (e *ListWaypointsEndpoint) Path() string {
	return fmt.Sprintf("/systems/%w/waypoints", e.SystemSymbol)
}

func (e *ListWaypointsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ListWaypointsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ListWaypointsEndpoint) IsTokenRequired() bool {
	return false
}

func (e *ListWaypointsEndpoint) Query() url.Values {
	query := url.Values{}

	pagination := e.Pagination.Query()
	for key, valArray := range pagination {
		for _, value := range valArray {
			query.Add(key, value)
		}
	}

	for _, trait := range e.Traits {
		query.Add("traits", trait)
	}

	return query
}

var _ ApiEndpoint = (*ListWaypointsEndpoint)(nil)
var _ WithQuery = (*ListWaypointsEndpoint)(nil)


type GetWaypointEndpoint struct {
	SystemSymbol string
	WaypointSymbol string
}

func (e *GetWaypointEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetWaypointEndpoint) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s", 
		e.SystemSymbol, e.WaypointSymbol)
}

func (e *GetWaypointEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetWaypointEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetWaypointEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetWaypointEndpoint)(nil)


type GetConstructionEndpoint struct {
	SystemSymbol string
	WaypointSymbol string
}

func (e *GetConstructionEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetConstructionEndpoint) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s/construction", 
		e.SystemSymbol, e.WaypointSymbol)
}

func (e *GetConstructionEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetConstructionEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetConstructionEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetConstructionEndpoint)(nil)


type SupplyConstructionEndpoint struct {
	SystemSymbol string `json:"-"`
	WaypointSymbol string `json:"-"`

	ShipSymbol string `json:"shipSymbol"`
	ItemSymbol string `json:"tradeSymbol"`
	Units int `json:"units"`
}

func (e *SupplyConstructionEndpoint) Method() string {
	return http.MethodPost
}

func (e *SupplyConstructionEndpoint) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s/construction/supply", 
		e.SystemSymbol, e.WaypointSymbol)
}

func (e *SupplyConstructionEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *SupplyConstructionEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *SupplyConstructionEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*SupplyConstructionEndpoint)(nil)


type GetMarketEndpoint struct {
	SystemSymbol string
	WaypointSymbol string
}

func (e *GetMarketEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetMarketEndpoint) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s/market", 
		e.SystemSymbol, e.WaypointSymbol)
}

func (e *GetMarketEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetMarketEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetMarketEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetMarketEndpoint)(nil)


type GetJumpGateEndpoint struct {
	SystemSymbol string
	WaypointSymbol string
}

func (e *GetJumpGateEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetJumpGateEndpoint) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s/jump-gate", 
		e.SystemSymbol, e.WaypointSymbol)
}

func (e *GetJumpGateEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetJumpGateEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetJumpGateEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetJumpGateEndpoint)(nil)


type GetShipyardEndpoint struct {
	SystemSymbol string
	WaypointSymbol string
}

func (e *GetShipyardEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetShipyardEndpoint) Path() string {
	return fmt.Sprintf("/systems/%s/waypoints/%s/shipyard", 
		e.SystemSymbol, e.WaypointSymbol)
}

func (e *GetShipyardEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetShipyardEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetShipyardEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetShipyardEndpoint)(nil)
