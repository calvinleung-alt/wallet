package types

type CheckBalanceInput struct {
	WalletID string
}

type CheckBalanceOutput struct {
	WalletID string
	Balance  int64
}

type DepositInput struct {
	WalletID string
	Amount   int64
}

type DepositOutput struct {
	WalletID string
	Balance  int64
}

type WithdrawInput struct {
	WalletID string
	Amount   int64
}

type WithdrawOutput struct {
	WalletID string
	Balance  int64
}

type SendInput struct {
	SrcWalletID  string
	DestWalletID string
	Amount       int64
}

type SendOutput struct {
	DestWalletID      string
	DestWalletBalance int64
}
