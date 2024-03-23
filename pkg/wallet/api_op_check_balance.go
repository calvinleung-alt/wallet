package wallet

import "github.com/calvinleung-alt/wallet/types"

func (c *Client) CheckBalance(params *types.CheckBalanceInput) (*types.CheckBalanceOutput, error) {
	rst, err := c.invokeOperation("CheckBalance", params)
	if err != nil {
		return nil, err
	}
	out := rst.(*types.CheckBalanceOutput)
	return out, nil
}
