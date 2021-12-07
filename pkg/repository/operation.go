package repository

import (
	"fmt"
	"github.com/K0STYAa/AvitoTech"
	"github.com/jmoiron/sqlx"
)

type OperationPostgres struct {
	db *sqlx.DB
}

func NewOperationPostgres(db *sqlx.DB) *OperationPostgres {
	return &OperationPostgres{db: db}
}

func (r *OperationPostgres) Accrual(userId int, amount int) (error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	fmt.Printf("INSERT INTO %s (sender_id, receiver_id, amount, departure_time) VALUES (0, %v, %v, now())", historyTable, userId, amount)
	hist_query := fmt.Sprintf("INSERT INTO %s (sender_id, receiver_id, amount, departure_time) VALUES (0, $1, $2, now())", historyTable)
	_, err = tx.Exec(hist_query, userId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}


	var user AvitoTech.Users
	query := fmt.Sprintf(`SELECT * FROM %s t1 WHERE t1.id = $1`, usersTable)
	err_2 := r.db.Get(&user, query, userId)
	if err_2 != nil {
		tx.Rollback()
		return err_2
	}

	if user.Id == userId {
		user_query := fmt.Sprintf(`UPDATE %s t1 SET balance = balance + $2 WHERE t1.id = $1`, usersTable)
		_, err = tx.Exec(user_query, userId, amount)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		user_query := fmt.Sprintf(`INSERT INTO %s (id, balance) VALUES ($1, $2)`, usersTable)
		_, err = tx.Exec(user_query, userId, amount)
		if err != nil {
			tx.Rollback()
			return err
		}
	}


	return tx.Commit()
}

func (r *OperationPostgres) WriteDowns(userId int, amount int) (error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	hist_query := fmt.Sprintf("INSERT INTO %s (sender_id, receiver_id, amount, departure_time) VALUES ($1, 0, $2, now())", historyTable)
	_, err = tx.Exec(hist_query, userId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	user_query := fmt.Sprintf(`UPDATE %s t1 SET balance = balance - $2 WHERE t1.id = $1`, usersTable)
	_, err = tx.Exec(user_query, userId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}