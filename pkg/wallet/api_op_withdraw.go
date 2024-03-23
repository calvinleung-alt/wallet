package wallet

import "github.com/calvinleung-alt/wallet/types"

func (c *Client) Withdraw(params *types.WithdrawInput) (*types.WithdrawOutput, error) {
	rst, err := c.invokeOperation("Withdraw", params)
	if err != nil {
		return nil, err
	}
	out := rst.(*types.WithdrawOutput)
	return out, nil
}
