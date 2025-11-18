package paystack_go

import (
	"github.com/go-resty/resty/v2"
)

const baseURL = "https://api.paystack.co"

type Client struct {
	apiKey string
	resty  *resty.Client
}

func New(apiKey string) *Client {
	client := resty.New().SetBaseURL(baseURL).
		SetAuthToken(apiKey)

	return &Client{
		apiKey: apiKey,
		resty:  client,
	}
}
