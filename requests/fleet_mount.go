package requests

import (
	"time"
	"fmt"
	"net/http"
	"encoding/json"
)

type GetMountsEndpoint struct {
	ShipSymbol string
}

func (e *GetMountsEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetMountsEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/mounts", e.ShipSymbol)
}

func (e *GetMountsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetMountsEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetMountsEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetMountsEndpoint{}


type InstallMountEndpoint struct {
	ShipSymbol string `json:"-"`

	MountSymbol string `json:"symbol"`
}

func (e *InstallMountEndpoint) Method() string {
	return http.MethodPost
}

func (e *InstallMountEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/mounts/install", e.ShipSymbol)
}

func (e *InstallMountEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *InstallMountEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *InstallMountEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &InstallMountEndpoint{}


type RemoveMountEndpoint struct {
	ShipSymbol string `json:"-"`

	MountSymbol string `json:"symbol"`
}

func (e *RemoveMountEndpoint) Method() string {
	return http.MethodPost
}

func (e *RemoveMountEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/mounts/remove", e.ShipSymbol)
}

func (e *RemoveMountEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *RemoveMountEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *RemoveMountEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &RemoveMountEndpoint{}


type ExtractResourcesEndpoint struct {
	ShipSymbol string
}

func (e *ExtractResourcesEndpoint) Method() string {
	return http.MethodPost
}

func (e *ExtractResourcesEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/extract", e.ShipSymbol)
}

func (e *ExtractResourcesEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ExtractResourcesEndpoint) Headers() map[string]string {
	return nil
}

func (e *ExtractResourcesEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &ExtractResourcesEndpoint{}


type ExtractResourcesWithSurveyEndpoint struct {
	ShipSymbol string `json:"-"`

	Signature string `json:"signature"`
	Symbol string `json:"symbol"`
	Deposits []struct{
		Symbol string `json:"symbol"`
	} `json:"deposits"`
	Expiration time.Time `json:"expiration"`
	Size string `json:"size"`
}

func (e *ExtractResourcesWithSurveyEndpoint) Method() string {
	return http.MethodPost
}

func (e *ExtractResourcesWithSurveyEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/extract/survey", e.ShipSymbol)
}

func (e *ExtractResourcesWithSurveyEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *ExtractResourcesWithSurveyEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func (e *ExtractResourcesWithSurveyEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &ExtractResourcesWithSurveyEndpoint{} 


type SiphonResourcesEndpoint struct {
	ShipSymbol string
}

func (e *SiphonResourcesEndpoint) Method() string {
	return http.MethodPost
}

func (e *SiphonResourcesEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/siphon", e.ShipSymbol)
}

func (e *SiphonResourcesEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *SiphonResourcesEndpoint) Headers() map[string]string {
	return nil
}

func (e *SiphonResourcesEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &SiphonResourcesEndpoint{}


type ScanSystemsEndpoint struct {
	ShipSymbol string
}

func (e *ScanSystemsEndpoint) Method() string {
	return http.MethodPost
}

func (e *ScanSystemsEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/scan/systems", e.ShipSymbol)
}

func (e *ScanSystemsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ScanSystemsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ScanSystemsEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &ScanSystemsEndpoint{}


type ScanWaypointsEndpoint struct {
	ShipSymbol string
}

func (e *ScanWaypointsEndpoint) Method() string {
	return http.MethodPost
}

func (e *ScanWaypointsEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/scan/waypoints", e.ShipSymbol)
}

func (e *ScanWaypointsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ScanWaypointsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ScanWaypointsEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &ScanWaypointsEndpoint{}


type ScanShipsEndpoint struct {
	ShipSymbol string
}

func (e *ScanShipsEndpoint) Method() string {
	return http.MethodPost
}

func (e *ScanShipsEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/scan/ships", e.ShipSymbol)
}

func (e *ScanShipsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ScanShipsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ScanShipsEndpoint) IsTokenRequired() bool {
	return true
}


type CreateSurveyEndpoint struct {
	ShipSymbol string
}

func (e *CreateSurveyEndpoint) Method() string {
	return http.MethodPost
}

func (e *CreateSurveyEndpoint) Path() string {
	return fmt.Sprintf("/my/ships/%s/survey", e.ShipSymbol)
}

func (e *CreateSurveyEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *CreateSurveyEndpoint) Headers() map[string]string {
	return nil
}

func (e *CreateSurveyEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &CreateSurveyEndpoint{}
