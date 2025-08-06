package requests

import (
	"fmt"
	"net/http"
)

type ListFactionsEndpoint struct {
	Pagination
}

func (e *ListFactionsEndpoint) Method() string {
	return http.MethodGet
}

func (e *ListFactionsEndpoint) Path() string {
	return "/factions"
}

func (e *ListFactionsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ListFactionsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ListFactionsEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*ListFactionsEndpoint)(nil)
var _ WithQuery = (*ListFactionsEndpoint)(nil)


type GetFactionEndpoint struct {
	FactionSymbol string
}

func (e *GetFactionEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetFactionEndpoint) Path() string {
	return fmt.Sprintf("/factions/%s", e.FactionSymbol)
}

func (e *GetFactionEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetFactionEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetFactionEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetFactionEndpoint)(nil)


type GetFactionReputationsEndpoint struct {
	Pagination
}

func (e *GetFactionReputationsEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetFactionReputationsEndpoint) Path() string {
	return "/my/factions"
}

func (e *GetFactionReputationsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetFactionReputationsEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetFactionReputationsEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = (*GetFactionReputationsEndpoint)(nil)
var _ WithQuery = (*GetFactionReputationsEndpoint)(nil)
