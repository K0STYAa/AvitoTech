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

	var count_zero int
	query_zero := fmt.Sprintf(`SELECT count(*) FROM %s t1 WHERE t1.id = 0`, usersTable)
	r.db.Get(&count_zero, query_zero)
	if count_zero == 0 {
		insert_zero := fmt.Sprintf(`INSERT INTO %s (id, balance) VALUES (0, 0)`, usersTable)
		_, err = tx.Exec(insert_zero)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	var user AvitoTech.Users
	query := fmt.Sprintf(`SELECT * FROM %s t1 WHERE t1.id = $1`, usersTable)
	r.db.Get(&user, query, userId)

	var user_query string
	if user.Id == userId {
		user_query = fmt.Sprintf(`UPDATE %s t1 SET balance = balance + $2 WHERE t1.id = $1`, usersTable)
	} else {
		user_query = fmt.Sprintf(`INSERT INTO %s (id, balance) VALUES ($1, $2)`, usersTable)
	}
	_, err = tx.Exec(user_query, userId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	hist_query := fmt.Sprintf("INSERT INTO %s (sender_id, receiver_id, amount, departure_time) VALUES (0, $1, $2, now())", historyTable)
	_, err = tx.Exec(hist_query, userId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}


	return tx.Commit()
}

func (r *OperationPostgres) WriteDowns(userId int, amount int) (error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var count_zero int
	query_zero := fmt.Sprintf(`SELECT count(*) FROM %s t1 WHERE t1.id = 0`, usersTable)
	r.db.Get(&count_zero, query_zero)
	if count_zero == 0 {
		insert_zero := fmt.Sprintf(`INSERT INTO %s (id, balance) VALUES (0, 0)`, usersTable)
		_, err = tx.Exec(insert_zero)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	user_query := fmt.Sprintf(`UPDATE %s t1 SET balance = balance - $2 WHERE t1.id = $1`, usersTable)
	_, err = tx.Exec(user_query, userId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	hist_query := fmt.Sprintf("INSERT INTO %s (sender_id, receiver_id, amount, departure_time) VALUES ($1, 0, $2, now())", historyTable)
	_, err = tx.Exec(hist_query, userId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *OperationPostgres) Transfer(senderId int, receiver_id int, amount int) (error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var receiver AvitoTech.Users
	query := fmt.Sprintf(`SELECT * FROM %s t1 WHERE t1.id = $1`, usersTable)
	err = r.db.Get(&receiver, query, receiver_id)

	var receiver_query string
	if receiver.Id == receiver_id {
		receiver_query = fmt.Sprintf(`UPDATE %s t1 SET balance = balance + $2 WHERE t1.id = $1`, usersTable)
	} else {
		receiver_query = fmt.Sprintf(`INSERT INTO %s (id, balance) VALUES ($1, $2)`, usersTable)
	}
	_, err = tx.Exec(receiver_query, receiver_id, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	sender_query := fmt.Sprintf(`UPDATE %s t1 SET balance = balance - $2 WHERE t1.id = $1`, usersTable)
	_, err = tx.Exec(sender_query, senderId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}