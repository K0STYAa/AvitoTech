package repository

import (
	"github.com/K0STYAa/AvitoTech"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetById(userId int) (AvitoTech.Users, error)
}

type History interface {
	GetById(historyId int) (AvitoTech.History, error)
}

type Operation interface {

}

type Repository struct {
	User
	History
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
		History: NewHistoryPostgres(db),
	}
}