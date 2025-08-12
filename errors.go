package spacedk

import (
	"errors"
	"fmt"
	"strings"
)

type ErrorInfo struct {
	Message string `json:"Message"`
	Code int `json:"code"`
	Data any `json:"data"`
}

type ApiError struct {
	Info ErrorInfo `json:"error"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%d: %s", e.Info.Code, e.Info.Message)
}

func (e *ApiError) Data() any {
	return e.Info.Data
}

var NoCooldown error = errors.New("there is no cooldown")


type MultiError struct {
	errorList []error
}

func NewMultiError() *MultiError {
	return &MultiError{
		errorList: make([]error, 0, 1),
	}
}

func (m *MultiError) HasErrors() bool {
	return len(m.errorList) > 0
}

func (m *MultiError) Add(err error) {
	m.errorList = append(m.errorList, err)
}

func (m *MultiError) Error() string {
	var errMsgs []string
	for _, err := range m.errorList {
		errMsgs = append(errMsgs, err.Error())
	}

	return strings.Join(errMsgs, "; ")
}


type AgentNotTrackedError struct {
	Agent string
}

func (e AgentNotTrackedError) Error() string {
	return fmt.Sprintf("agent '%s' is not tracked", e.Agent)
}
