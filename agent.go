package spacedk

import (
	"fmt"
	"github.com/Kirshoo/spacedk/requests"
)

type AgentService struct {
	client *Client
}

func NewAgentService(c *Client) *AgentService {
	return &AgentService{
		client: c,
	}
}

func (s *AgentService) List(page, limit int) ([]Agent, *Metadata, error) {
	req := &requests.ListAgentsEndpoint{
		Pagination: requests.Pagination{
			Page: page, 
			Limit: limit,
		},
	}

	reply, err := Do[[]Agent](s.client, req)
	if err != nil {
		return nil, nil, fmt.Errorf("list agents: %w", err)
	}

	return reply.Data, &reply.Meta, nil
}

func (s *AgentService) Get(symbol string) (*Agent, error) {
	req := &requests.GetAgentEndpoint{AgentSymbol: symbol}

	reply, err := Do[Agent](s.client, req)
	if err != nil {
		return nil, fmt.Errorf("get agent: %w", err)
	}

	return &reply.Data, nil
}

func (s *AgentService) GetOwn() (*OwnAgent, error) {
	req := &requests.GetOwnAgentEndpoint{}

	reply, err := Do[OwnAgent](s.client, req)
	if err != nil {
		return nil, fmt.Errorf("get own agent: %w", err)
	}

	return &reply.Data, nil
}

func (s *AgentService) GetEvents() ([]AgentEvent, error) {
	req := &requests.GetOwnAgentEventsEndpoint{}

	reply, err := Do[[]AgentEvent](s.client, req)
	if err != nil {
		return nil, fmt.Errorf("get recent agent events: %w", err)
	}

	return reply.Data, nil
}
