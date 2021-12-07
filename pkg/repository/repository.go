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
	Accrual(userId int, amount int) (error)
	WriteDowns(userId int, amount int) (error)
	Transfer(senderId int, receiver_id int, amount int) (error)
}

type Repository struct {
	User
	History
	Operation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
		History: NewHistoryPostgres(db),
		Operation: NewOperationPostgres(db),
	}
}