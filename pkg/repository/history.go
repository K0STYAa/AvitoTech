package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/K0STYAa/AvitoTech"
)

type HistoryPostgres struct {
	db *sqlx.DB
}

func NewHistoryPostgres(db *sqlx.DB) *HistoryPostgres {
	return &HistoryPostgres{db: db}
}

func (r *HistoryPostgres) GetById(historyId int) (AvitoTech.History, error) {
	var list AvitoTech.History

	query := fmt.Sprintf(`SELECT * FROM %s t1 WHERE t1.id = $1`, historyTable)
	err := r.db.Get(&list, query, historyId)

	return list, err
}
