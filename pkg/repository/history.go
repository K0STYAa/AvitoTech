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

func (r *HistoryPostgres) GetById(historyId int) ([]AvitoTech.History, error) {
	var hist []AvitoTech.History

	query := fmt.Sprintf(`SELECT * FROM %s t1 WHERE t1.sender_id = $1 or t1.receiver_id = $1`, historyTable)
	err := r.db.Select(&hist, query, historyId)

	return hist, err
}
