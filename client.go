package spacedk

import (
	"fmt"
	"io"
	"bytes"
	"net/http"
	"encoding/json"

	endpoints "github.com/Kirshoo/spacedk/requests"
)

type Client struct {
	network http.Client
	baseURL string
	token string
	userAgent string

	Agents *AgentService
	Contracts *ContractService
	Factions *FactionService
	Fleet *FleetService
	Systems *SystemService
}

const (
	DefaultBaseURL string = "https://api.spacetraders.io/v2"
	defaultUserAgent string = "spacedk/0.1.0 (+https://github.com/Kirshoo/spacedk)"
)

var (
	DefaultClient = &Client{baseURL: DefaultBaseURL}
)

type ClientConfig struct {
	BaseURL string
	Token string
	UserAgent string
}

type ClientOption func(*ClientConfig)
func WithBaseURL(url string) ClientOption {
	return func(cfg *ClientConfig) {
		cfg.BaseURL = url
	}
}

func WithToken(token string) ClientOption {
	return func(cfg *ClientConfig) {
		cfg.Token = token
	}
}

func WithUserAgent(ua string) ClientOption {
	return func(cfg *ClientConfig) {
		cfg.UserAgent = ua
	}
}

func NewClient(opts ...ClientOption) *Client {
	config := ClientConfig{
		BaseURL: DefaultBaseURL,
		UserAgent: defaultUserAgent,
	}
	
	for _, opt := range opts {
		opt(&config)
	}

	c := &Client{
		network: http.Client{},
		baseURL: config.BaseURL,
		token: config.Token,
		userAgent: config.UserAgent,
	}

	c.Agents = NewAgentService(c)
	c.Contracts = NewContractService(c)
	c.Factions = NewFactionService(c)
	c.Fleet = NewFleetService(c)
	c.Systems = NewSystemService(c)

	return c
}

// May be used to change between different agents
func (c *Client) SetToken(newToken string) {
	c.token = newToken
}

func (c *Client) Execute(request endpoints.ApiEndpoint) (*http.Response, error) {
	if request.IsTokenRequired() && c.token == "" {
		return nil, fmt.Errorf("token is required, client is missing a token")
	}

	url := c.baseURL + request.Path()

	var buf io.Reader
	byteArray, err := request.Body()
	if err != nil {
		return nil, fmt.Errorf("encoding body: %w", err)
	}
	buf = bytes.NewReader(byteArray)

	var req *http.Request
	if ctx, ok := request.(endpoints.WithContext); ok {
		req, err = http.NewRequestWithContext(ctx.Context(), request.Method(), url, buf)
	} else {
		req, err = http.NewRequest(request.Method(), url, buf)
	}
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("User-Agent", c.userAgent)

	for key, val := range request.Headers() {
		req.Header.Add(key, val)
	}

	if request.IsTokenRequired() {
		req.Header.Add("Authorization", "Bearer " + c.token)
	}

	if query, ok := request.(endpoints.WithQuery); ok {
		req.URL.RawQuery = query.Query().Encode()
	}

	return c.network.Do(req)
}

type StatusHandler func(statusCode int, body []byte, headers map[string][]string, v any) error

type HandlerOptions struct {
	CustomHandlers map[int]StatusHandler
}

type HandlerOption func(*HandlerOptions)
func WithStatusHandler(code int, handler StatusHandler) HandlerOption {
	return func(options *HandlerOptions) {
		options.CustomHandlers[code] = handler
	}
}

// TODO: Check rate limiter headers
func (c *Client) Handle(resp *http.Response, v any, opts ...HandlerOption) error {
	defer resp.Body.Close()

	var options HandlerOptions
	for _, opt := range opts {
		opt(&options)
	}

	if handler, ok := options.CustomHandlers[resp.StatusCode]; ok {
		body, _ := io.ReadAll(resp.Body)
		return handler(resp.StatusCode, body, resp.Header, v)
	}
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Some error occurred
		var apiErr ApiError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return fmt.Errorf("error response: decoding body: %w", err)
		}

		return fmt.Errorf("error response: %w", apiErr)
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}

	return nil
}

// Most of the server responses are in {data: ..., meta: ...} format
// This function allows for easy wrapping
// Wraps DoRaw() to return *Response[T] instead of *T
func Do[T any](c *Client, request endpoints.ApiEndpoint, opts ...HandlerOption) (*Response[T], error) {
	return DoRaw[Response[T]](c, request, opts...)
}

// Only used when server doesnt respond in Response struct format
// i.e. GetStatus, ErrorCodes
func DoRaw[T any](c *Client, request endpoints.ApiEndpoint, opts ...HandlerOption) (*T, error) {
	resp, err := c.Execute(request)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	var reply T
	if err = c.Handle(resp, &reply, opts...); err != nil {
		return nil, fmt.Errorf("handling reply: %w", err)
	}

	return &reply, nil
}
