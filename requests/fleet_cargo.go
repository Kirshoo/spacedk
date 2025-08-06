package requests

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type GetCargoEndpoint struct {
	ShipSymbol string
}

func (e *GetCargoEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetCargoEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/cargo", e.ShipSymbol)
}

func (e *GetCargoEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetCargoEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetCargoEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetCargoEndpoint{}


type PurchaseCargoEndpoint struct {
	ShipSymbol string `json:"-"`

	ItemSymbol string `json:"symbol"`
	Units int `json:"units"`
}

func (e *PurchaseCargoEndpoint) Method() string {
	return http.MethodPost
}

func (e *PurchaseCargoEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/purchase", e.ShipSymbol)
}

func (e *PurchaseCargoEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *PurchaseCargoEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *PurchaseCargoEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &PurchaseCargoEndpoint{}


type SellCargoEndpoint struct {
	ShipSymbol string `json:"-"`

	ItemSymbol string `json:"symbol"`
	Units int `json:"units"`
}

func (e *SellCargoEndpoint) Method() string {
	return http.MethodPost
}

func (e *SellCargoEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/sell", e.ShipSymbol)
}

func (e *SellCargoEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *SellCargoEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *SellCargoEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &SellCargoEndpoint{}


type TransferCargoEndpoint struct {
	ShipSymbol string `json:"-"`

	ItemSymbol string `json:"tradeSymbol"`
	Units int `json:"units"`
	TargetSymbol string `json:"shipSymbol"`
}

func (e *TransferCargoEndpoint) Method() string {
	return http.MethodPost
}

func (e *TransferCargoEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/transfer", e.ShipSymbol)
}

func (e *TransferCargoEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *TransferCargoEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *TransferCargoEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &TransferCargoEndpoint{}


type JettisonCargoEndpoint struct {
	ShipSymbol string `json:"-"`

	ItemSymbol string `json:"symbol"`
	Units int `json:"units"`
}

func (e *JettisonCargoEndpoint) Method() string {
	return http.MethodPost
}

func (e *JettisonCargoEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/jettison", e.ShipSymbol)
}

func (e *JettisonCargoEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *JettisonCargoEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *JettisonCargoEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &JettisonCargoEndpoint{}
