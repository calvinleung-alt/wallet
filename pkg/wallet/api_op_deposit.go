package wallet

import "github.com/calvinleung-alt/wallet/types"

func (c *Client) Deposit(params *types.DepositInput) (*types.DepositOutput, error) {
	rst, err := c.invokeOperation("Deposit", params)
	if err != nil {
		return nil, err
	}
	out := rst.(*types.DepositOutput)
	return out, nil
}
