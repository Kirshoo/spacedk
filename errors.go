package spacedk

import (
	"errors"
	"fmt"
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
