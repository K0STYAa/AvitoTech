package repository

import (
	"github.com/K0STYAa/AvitoTech"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetById(userId int) (AvitoTech.Users, error)
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