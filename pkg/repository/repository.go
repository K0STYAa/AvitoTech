package repository

import "github.com/jmoiron/sqlx"

type User interface {

}

type History interface {

}

type Operation interface {

}

type Repository struct {
	User
	History
	Operation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}