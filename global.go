package spacedk

import "github.com/Kirshoo/spacedk/requests"

func (c *Client) Status() (*ApiServer, error) {
	req := &requests.GetStatusEndpoint{}

	reply, err := DoRaw[ApiServer](c, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (c *Client) ErrorCodes() (*ErrorCodesInfo, error) {
	req := &requests.GetErrorCodesEndpoint{}

	reply, err := DoRaw[ErrorCodesInfo](c, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
