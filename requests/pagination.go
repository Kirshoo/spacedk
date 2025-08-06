package requests

import (
	"net/url"
	"strconv"
)

type Pagination struct {
	Page int
	Limit int
}

func (p *Pagination) Query() url.Values {
	query := url.Values{}

	query.Add("page", strconv.Itoa(p.Page))
	query.Add("limit", strconv.Itoa(p.Limit))

	return query
}
