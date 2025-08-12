package requests

import (
	"net/http"
	"encoding/json"
)

type GetAccountEndpoint struct {}

func (e *GetAccountEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetAccountEndpoint) Path() string {
	return "/my/account"
}

func (e *GetAccountEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetAccountEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetAccountEndpoint) IsTokenRequired() bool {
	return true
}

var _ ApiEndpoint = &GetAccountEndpoint{}


type RegisterAgentEndpoint struct {
	AccountToken string `json:"-"`

	AgentSymbol string `json:"symbol"`
	FactionSymbol string `json:"faction"`
}

func (e *RegisterAgentEndpoint) Method() string {
	return http.MethodPost
}

func (e *RegisterAgentEndpoint) Path() string {
	return "/register"
}

func (e *RegisterAgentEndpoint) Body() ([]byte, error) {
	return json.Marshal(*e)
}

func (e *RegisterAgentEndpoint) Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"Authorization": "Bearer " + e.AccountToken,
	}
}

func (e *RegisterAgentEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = &RegisterAgentEndpoint{}
