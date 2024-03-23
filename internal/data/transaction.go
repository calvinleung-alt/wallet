package data

type Transaction struct {
	*Repository
}

func NewTransaction(repository *Repository) *Transaction {
	return &Transaction{Repository: repository}
}

func (t *Transaction) Start() error {
	return nil
}

func (t *Transaction) Commit() error {
	return nil
}
