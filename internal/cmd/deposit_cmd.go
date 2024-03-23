package cmd

import (
	"github.com/calvinleung-alt/wallet/internal/data"
	"github.com/calvinleung-alt/wallet/internal/models"
	"github.com/calvinleung-alt/wallet/types"
)

type DepositCmd struct {
	*data.Repository
}

func NewDepositCmd(repository *data.Repository) *DepositCmd {
	return &DepositCmd{Repository: repository}
}

func (d *DepositCmd) Execute(params interface{}) (interface{}, error) {
	tx := d.BeginTx()
	if err := tx.Start(); err != nil {
		return nil, err
	}
	out, err := d.execWithTransaction(params, tx)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return out, nil
}

func (d *DepositCmd) execWithTransaction(params interface{}, tx *data.Transaction) (interface{}, error) {
	input := params.(*types.DepositInput)
	rst, err := tx.Repository.FindByID(input.WalletID)
	if err != nil {
		return nil, err
	}
	wallet := rst.(*models.Wallet)
	wallet.Balance += input.Amount
	if err = tx.Repository.Save(wallet); err != nil {
		return nil, err
	}
	return &types.DepositOutput{
		WalletID: wallet.ObjectID(),
		Balance:  wallet.Balance,
	}, nil
}
