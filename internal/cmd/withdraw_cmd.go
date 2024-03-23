package cmd

import (
	"errors"
	"github.com/calvinleung-alt/wallet/internal/data"
	"github.com/calvinleung-alt/wallet/internal/models"
	"github.com/calvinleung-alt/wallet/types"
)

type WithdrawCmd struct {
	*data.Repository
}

func NewWithdrawCmd(repository *data.Repository) *WithdrawCmd {
	return &WithdrawCmd{Repository: repository}
}

func (w *WithdrawCmd) Execute(params interface{}) (result interface{}, err error) {
	tx := w.BeginTx()
	if err := tx.Start(); err != nil {
		return nil, err
	}
	out, err := w.execWithTransaction(params, tx)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *WithdrawCmd) execWithTransaction(params interface{}, tx *data.Transaction) (interface{}, error) {
	input := params.(*types.WithdrawInput)
	rst, err := tx.FindByID(input.WalletID)
	if err != nil {
		return nil, err
	}
	wallet := rst.(*models.Wallet)
	wallet.Balance -= input.Amount
	if wallet.Balance < 0 {
		return nil, errors.New("insufficient balance")
	}
	if err := tx.Save(wallet); err != nil {
		return nil, err
	}
	return &types.WithdrawOutput{
		WalletID: wallet.ObjectID(),
		Balance:  wallet.Balance,
	}, nil
}
