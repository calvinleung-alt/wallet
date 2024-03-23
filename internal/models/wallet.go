package models

type Wallet struct {
	ID      string
	Balance int64
}

func (w *Wallet) ObjectID() string {
	return w.ID
}
