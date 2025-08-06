package spacedk

import "github.com/Kirshoo/spacedk/requests"

type FactionService struct {
	client *Client
}

func NewFactionService(c *Client) *FactionService {
	return &FactionService{
		client: c,
	}
}

func (s *FactionService) List(page, limit int) ([]Faction, *Metadata, error) {
	req := &requests.ListFactionsEndpoint{
		Pagination: requests.Pagination{
			Page: page, 
			Limit: limit,
		},
	}

	reply, err := Do[[]Faction](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return reply.Data, &reply.Meta, nil
}

func (s *FactionService) Get(symbol string) (*Faction, error) {
	req := &requests.GetFactionEndpoint{
		FactionSymbol: symbol,
	}

	reply, err := Do[Faction](s.client, req)
	if err != nil {
		return nil, err
	}

	return &reply.Data, nil
}

func (s *FactionService) GetReputations(page, limit int) ([]FactionReputation, *Metadata, error) {
	req := &requests.GetFactionReputationsEndpoint{
		Pagination: requests.Pagination{
			Page: page, 
			Limit: limit,
		},
	}
	reply, err := Do[[]FactionReputation](s.client, req)
	if err != nil {
		return nil, nil, err
	}

	return reply.Data, &reply.Meta, nil
}
