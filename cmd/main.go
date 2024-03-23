package main

import (
	"github.com/calvinleung-alt/wallet/internal/cmd"
	"github.com/calvinleung-alt/wallet/internal/data"
	iinvoker "github.com/calvinleung-alt/wallet/internal/invoker"
	"github.com/calvinleung-alt/wallet/internal/models"
	"github.com/calvinleung-alt/wallet/pkg/wallet"
	"github.com/calvinleung-alt/wallet/types"
	"log"
)

func main() {
	repository := data.NewRepository(
		&models.Wallet{ID: "0", Balance: 100},
		&models.Wallet{ID: "1", Balance: 100},
		&models.Wallet{ID: "2", Balance: 100},
		&models.Wallet{ID: "3", Balance: 100},
		&models.Wallet{ID: "4", Balance: 100},
		&models.Wallet{ID: "5", Balance: 100},
	)

	invoker := iinvoker.NewInvoker(map[string]iinvoker.Command{
		"CheckBalance": cmd.NewCheckBalanceCmd(repository),
		"Deposit":      cmd.NewDepositCmd(repository),
		"Withdraw":     cmd.NewWithdrawCmd(repository),
		"Send":         cmd.NewSendCmd(repository, cmd.NewWithdrawCmd(repository), cmd.NewDepositCmd(repository)),
	})

	client := wallet.NewClient(invoker)

	log.Println(client.Deposit(&types.DepositInput{
		WalletID: "1",
		Amount:   500,
	}))

	log.Println(client.Withdraw(&types.WithdrawInput{
		WalletID: "2",
		Amount:   100,
	}))

	log.Println(client.Send(&types.SendInput{
		SrcWalletID:  "1",
		DestWalletID: "2",
		Amount:       200,
	}))

	log.Println(client.CheckBalance(&types.CheckBalanceInput{
		WalletID: "2",
	}))
}
