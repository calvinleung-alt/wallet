package wallet

import (
	"github.com/calvinleung-alt/wallet/internal/invoker"
)

type Invoker interface {
	Get(opID string) (invoker.Command, error)
}

type Client struct {
	invoker Invoker
}

func NewClient(invoker Invoker) *Client {
	return &Client{invoker: invoker}
}

func (c *Client) invokeOperation(opID string, params interface{}) (interface{}, error) {
	cmd, err := c.invoker.Get(opID)
	if err != nil {
		return nil, err
	}
	return cmd.Execute(params)
}
