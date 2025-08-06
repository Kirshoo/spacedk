package requests

import (
	"fmt"
	"net/http"
)

type ListAgentsEndpoint struct {
	Pagination
}

func (e *ListAgentsEndpoint) Method() string {
	return http.MethodGet
}

func (e *ListAgentsEndpoint) Path() string {
	return "/agents"
}

func (e *ListAgentsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *ListAgentsEndpoint) Headers() map[string]string {
	return nil
}

func (e *ListAgentsEndpoint) IsTokenRequired() bool {
	return false
}

// Making sure endpoint implements all required interfaces
var _ ApiEndpoint = (*ListAgentsEndpoint)(nil)
var _ WithQuery = (*ListAgentsEndpoint)(nil)


type GetAgentEndpoint struct {
	AgentSymbol string
}

func (e *GetAgentEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetAgentEndpoint) Path() string {
	return fmt.Sprintf("/agents/%s", e.AgentSymbol)
}

func (e *GetAgentEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetAgentEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetAgentEndpoint) IsTokenRequired() bool {
	return false
}

// Making sure endpoint implements all required interfaces
var _ ApiEndpoint = (*GetAgentEndpoint)(nil)


type GetOwnAgentEndpoint struct {}

func (e *GetOwnAgentEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetOwnAgentEndpoint) Path() string {
	return "/my/agent"
}

func (e *GetOwnAgentEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetOwnAgentEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetOwnAgentEndpoint) IsTokenRequired() bool {
	return true
}

// Making sure endpoint implements all required interfaces
var _ ApiEndpoint = (*GetOwnAgentEndpoint)(nil)


type GetOwnAgentEventsEndpoint struct {}

func (e *GetOwnAgentEventsEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetOwnAgentEventsEndpoint) Path() string {
	return "/my/agent/events"
}

func (e *GetOwnAgentEventsEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetOwnAgentEventsEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetOwnAgentEventsEndpoint) IsTokenRequired() bool {
	return true
}

// Making sure endpoint implements all required interfaces
var _ ApiEndpoint = (*GetOwnAgentEventsEndpoint)(nil)
