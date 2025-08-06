package requests

import (
	"context"
	"net/url"
)

type ApiEndpoint interface {
	Method() string
	Path() string
	Body() ([]byte, error) // Marshaled body
	Headers() map[string]string
	IsTokenRequired() bool
}

type WithContext interface {
	Context() context.Context
}

type WithQuery interface {
	Query() url.Values
}
