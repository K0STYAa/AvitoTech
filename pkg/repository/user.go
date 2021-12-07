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
	var user AvitoTech.Users

	query := fmt.Sprintf(`SELECT * FROM %s t1 WHERE t1.id = $1`, usersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}