package cmd

import (
	"github.com/calvinleung-alt/wallet/internal/data"
	"github.com/calvinleung-alt/wallet/types"
)

type SendCmd struct {
	*data.Repository
	withdrawCmd *WithdrawCmd
	depositCmd  *DepositCmd
}

func NewSendCmd(
	repository *data.Repository,
	withdrawCmd *WithdrawCmd,
	depositCmd *DepositCmd,
) *SendCmd {
	return &SendCmd{
		Repository:  repository,
		withdrawCmd: withdrawCmd,
		depositCmd:  depositCmd,
	}
}

func (s *SendCmd) Execute(params interface{}) (result interface{}, err error) {
	tx := s.BeginTx()

	if err := tx.Start(); err != nil {
		return nil, err
	}

	input := params.(*types.SendInput)

	if _, err := s.withdrawCmd.execWithTransaction(&types.WithdrawInput{
		WalletID: input.SrcWalletID,
		Amount:   input.Amount,
	}, tx); err != nil {
		return nil, err
	}

	out, err := s.depositCmd.execWithTransaction(&types.DepositInput{
		WalletID: input.DestWalletID,
		Amount:   input.Amount,
	}, tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	depositOutput := out.(*types.DepositOutput)
	return &types.SendOutput{
		DestWalletID:      depositOutput.WalletID,
		DestWalletBalance: depositOutput.Balance,
	}, nil
}
