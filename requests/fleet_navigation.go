package requests

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type DockShipEndpoint struct {
	ShipSymbol string
}

func (e *DockShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *DockShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/dock", e.ShipSymbol)
}

func (e *DockShipEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *DockShipEndpoint) Headers() map[string]string {
	return nil
}

func (e *DockShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &DockShipEndpoint{}


type JumpShipEndpoint struct {
	ShipSymbol string `json:"-"`

	DestinationSymbol string `json:"waypointSymbol"`
}

func (e *JumpShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *JumpShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/jump", e.ShipSymbol)
}

func (e *JumpShipEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *JumpShipEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *JumpShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &JumpShipEndpoint{}


type NavigateShipEndpoint struct {
	ShipSymbol string `json:"-"`

	DestinationSymbol string `json:"waypointSymbol"`
}

func (e *NavigateShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *NavigateShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/navigate", e.ShipSymbol)
}

func (e *NavigateShipEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *NavigateShipEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *NavigateShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &NavigateShipEndpoint{}


type WarpShipEndpoint struct {
	ShipSymbol string `json:"-"`

	DestinationSymbol string `json:"waypointSymbol"`
}

func (e *WarpShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *WarpShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/warp", e.ShipSymbol)
}

func (e *WarpShipEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *WarpShipEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *WarpShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &WarpShipEndpoint{}


type OrbitShipEndpoint struct {
	ShipSymbol string
}

func (e *OrbitShipEndpoint) Method() string {
	return http.MethodPost
}

func (e *OrbitShipEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/orbit", e.ShipSymbol)
}

func (e *OrbitShipEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *OrbitShipEndpoint) Headers() map[string]string {
	return nil
}

func (e *OrbitShipEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &OrbitShipEndpoint{}


type GetNavigationEndpoint struct {
	ShipSymbol string
}

func (e *GetNavigationEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetNavigationEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/nav", e.ShipSymbol)
}

func (e *GetNavigationEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetNavigationEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetNavigationEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetNavigationEndpoint{}


type ChangeNavigationModeEndpoint struct {
	ShipSymbol string `json:"-"`

	FlightMode string `json:"flightMode"`
}

func (e *ChangeNavigationModeEndpoint) Method() string {
	return http.MethodPatch
}

func (e *ChangeNavigationModeEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/nav", e.ShipSymbol)
}

func (e *ChangeNavigationModeEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *ChangeNavigationModeEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *ChangeNavigationModeEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &ChangeNavigationModeEndpoint{}
