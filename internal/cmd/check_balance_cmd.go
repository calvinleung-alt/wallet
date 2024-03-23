package cmd

import (
	"github.com/calvinleung-alt/wallet/internal/data"
	"github.com/calvinleung-alt/wallet/internal/models"
	"github.com/calvinleung-alt/wallet/types"
)

type CheckBalanceCmd struct {
	*data.Repository
}

func NewCheckBalanceCmd(repository *data.Repository) *CheckBalanceCmd {
	return &CheckBalanceCmd{
		Repository: repository,
	}
}

func (c *CheckBalanceCmd) Execute(payload interface{}) (result interface{}, err error) {
	params := payload.(*types.CheckBalanceInput)
	rst, err := c.FindByID(params.WalletID)
	if err != nil {
		return nil, err
	}
	wallet := rst.(*models.Wallet)
	return &types.CheckBalanceOutput{
		WalletID: wallet.ObjectID(),
		Balance:  wallet.Balance,
	}, nil
}
