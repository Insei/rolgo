package rolgo

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

// Client the base API Client
type Client struct {
	resty *resty.Client

	Rents RentsService
}

func NewClient(baseUrl string, apiKey string) *Client {
	var c = new(Client)

	c.resty = resty.New()
	c.resty.BaseURL = baseUrl
	c.resty.SetHeader("X-API-Key", apiKey)
	c.resty.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		if resp.IsError() {
			return errors.New(resp.Status())
		}

		apiResp := new(ApiResponse)
		c.JSONUnmarshal(resp.Body(), apiResp)
		if apiResp.Status.Code != 0 {
			return errors.New(apiResp.Status.Message)
		}

		return nil
	})


	c.Rents = &RentsServiceOp{client: c}

	return c
}
