package data

import (
	"errors"
)

type Data interface {
	ObjectID() string
}

type Repository struct {
	data []Data
}

func NewRepository(data ...Data) *Repository {
	return &Repository{data: data}
}

func (r *Repository) BeginTx() *Transaction {
	return NewTransaction(r)
}

func (r *Repository) FindByID(id string) (interface{}, error) {
	for _, dat := range r.data {
		if dat.ObjectID() == id {
			return dat, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *Repository) Save(update Data) error {
	for i, dat := range r.data {
		if dat.ObjectID() == update.ObjectID() {
			r.data[i] = update
			return nil
		}
	}
	return errors.New("not found")
}
