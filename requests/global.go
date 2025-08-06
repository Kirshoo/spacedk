package requests

import (
	"net/http"
)

type GetStatusEndpoint struct {}

func (e *GetStatusEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetStatusEndpoint) Path() string {
	return "/"
}

func (e *GetStatusEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetStatusEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetStatusEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetStatusEndpoint)(nil)


type GetErrorCodesEndpoint struct {}

func (e *GetErrorCodesEndpoint) Method() string {
	return http.MethodGet
}

func (e *GetErrorCodesEndpoint) Path() string {
	return "/error-codes"
}

func (e *GetErrorCodesEndpoint) Body() ([]byte, error) {
	return nil, nil
}

func (e *GetErrorCodesEndpoint) Headers() map[string]string {
	return nil
}

func (e *GetErrorCodesEndpoint) IsTokenRequired() bool {
	return false
}

var _ ApiEndpoint = (*GetErrorCodesEndpoint)(nil)
