package repository

import (
	"errors"
	"fmt"
	"github.com/K0STYAa/AvitoTech"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type HistoryPostgres struct {
	db *sqlx.DB
}

func NewHistoryPostgres(db *sqlx.DB) *HistoryPostgres {
	return &HistoryPostgres{db: db}
}

func (r *HistoryPostgres) GetById(historyId int, sort string, typeSort string, limit string, offset int) ([]AvitoTech.History, error) {
	var hist []AvitoTech.History

	if sort == "date" {
		sort = "departure_time"
	} else if sort == "sum" {
		sort = "amount"
	} else {
		return hist, errors.New("invalid sort param")
	}

	if typeSort == "inc" {
		typeSort = "ASC"
	} else if typeSort == "dec" {
		typeSort = "DESC"
	} else {
		return hist, errors.New("invalid type param")
	}

	limitInt, err_lim := strconv.Atoi(limit)
	if limit != "ALL" && (err_lim != nil || limitInt <= 0) {
		return hist, errors.New("invalid limit param")
	}

	query := fmt.Sprintf(`SELECT * FROM %s t1 WHERE t1.sender_id = $1 or t1.receiver_id = $1 ORDER BY %v %v LIMIT %v OFFSET %v`, historyTable, sort, typeSort, limit, offset)
	err := r.db.Select(&hist, query, historyId)

	return hist, err
}

func (r *HistoryPostgres) GetCountById(historyId int) (int, error) {
	var count int
	query := fmt.Sprintf(`SELECT count(*) FROM %s t1 WHERE t1.sender_id = $1 or t1.receiver_id = $1`, historyTable)
	err := r.db.Get(&count, query, historyId)
	return count, err
}
