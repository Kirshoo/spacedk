package requests

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type GetModulesEndpoint struct {
	ShipSymbol string
}

func (e *GetModulesEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetModulesEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/modules", e.ShipSymbol)
}

func (e *GetModulesEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetModulesEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetModulesEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetModulesEndpoint{}


type InstallModuleEndpoint struct {
	ShipSymbol string `json:"-"`

	ModuleSymbol string `json:"symbol"`
}

func (e *InstallModuleEndpoint) Method() string {
	return http.MethodPost
}

func (e *InstallModuleEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/modules/install", e.ShipSymbol)
}

func (e *InstallModuleEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *InstallModuleEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *InstallModuleEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &InstallModuleEndpoint{}


type RemoveModuleEndpoint struct {
	ShipSymbol string `json:"-"`

	ModuleSymbol string `json:"symbol"`
}

func (e *RemoveModuleEndpoint) Method() string {
	return http.MethodPost
}

func (e *RemoveModuleEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/modules/remove", e.ShipSymbol)
}

func (e *RemoveModuleEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *RemoveModuleEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *RemoveModuleEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &RemoveModuleEndpoint{}


type RefineResourceEndpoint struct {
	ShipSymbol string `json:"-"`

	ProductSymbol string `json:"produce"`
}

func (e *RefineResourceEndpoint) Method() string {
	return http.MethodPost
}

func (e *RefineResourceEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/refine", e.ShipSymbol)
}

func (e *RefineResourceEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *RefineResourceEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *RefineResourceEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &RefineResourceEndpoint{}
