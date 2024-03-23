package wallet

import "github.com/calvinleung-alt/wallet/types"

func (c *Client) Send(params *types.SendInput) (*types.SendOutput, error) {
	rst, err := c.invokeOperation("Send", params)
	if err != nil {
		return nil, err
	}
	out := rst.(*types.SendOutput)
	return out, nil
}
