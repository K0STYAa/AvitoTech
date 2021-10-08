package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/K0STYAa/AvitoTech"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetById(userId int) (AvitoTech.Users, error) {
	var list AvitoTech.Users

	query := fmt.Sprintf(`SELECT * FROM %s tl WHERE ul.user_id = $1`, usersTable)
	err := r.db.Get(&list, query, userId)

	return list, err
}