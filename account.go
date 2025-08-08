package spacedk

import (
	"fmt"
	"github.com/Kirshoo/spacedk/requests"
)

type (
	AccountServiceInterface interface {
		Get() (*Account, error)
		RegisterAgent(cfg RegistrationConfig, accountToken string) (*OwnAgent, *Faction, *Contract, []Ship, error)
	}

	AccountService struct {
		client *Client
	}

	AccountPayload struct {
		Account Account `json:"account"`
	}

	TokenAgentFactionContractShipsPayload struct {
		Token string `json:"token"`
		Agent OwnAgent `json:"agent"`
		Faction Faction `json:"faction"`
		Contract Contract `json:"contract"`
		Ships []Ship `json:"ships"`
	}
)

func NewAccountService(c *Client) *AccountService {
	return &AccountService{
		client: c,
	}
}

// Initiates request to server to get account of current agent
func (s *AccountService) Get() (*Account, error) {
	req := &requests.GetAccountEndpoint{}

	reply, err := Do[AccountPayload](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data.Account, nil
}

// Registers agent on the server and changes current agent token
// to token of newly created agent
// Saving token beforehand is highly advised
func (s *AccountService) RegisterAgent(cfg RegistrationConfig, accToken string) (*OwnAgent, *Faction, *Contract, []Ship, error) {
	if err := cfg.Validate(); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("validating registration information: %w", err)
	}

	req := &requests.RegisterAgentEndpoint{
		AccountToken: accToken,

		AgentSymbol: cfg.AgentSymbol,
		FactionSymbol: string(cfg.FactionSymbol),
	}

	reply, err := Do[TokenAgentFactionContractShipsPayload](s.client, req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	s.client.SetToken(reply.Data.Token)

	return &reply.Data.Agent, &reply.Data.Faction, &reply.Data.Contract, reply.Data.Ships, nil
}
