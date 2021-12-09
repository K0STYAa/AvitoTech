package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) accrual(c *gin.Context) {

	id, id_err := strconv.Atoi(c.Query("id"))
	if id_err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}


	amount, amount_err := strconv.ParseFloat(c.Query("amount"), 64)
	if amount_err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid amount param")
		return
	}

	accrual_err := h.services.Operation.Accrual(id, float64(amount))
	if accrual_err != nil {
		switch accrual_err.Error() {
		case `pq: new row for relation "history" violates check constraint "history_amount_check"`:
			newErrorResponse(c, http.StatusInternalServerError, "invalid amount")
		case `pq: new row for relation "users" violates check constraint "users_id_check"`:
			newErrorResponse(c, http.StatusInternalServerError, "invalid id")
		default:
			newErrorResponse(c, http.StatusInternalServerError, accrual_err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})

}

func (h *Handler) writedowns(c *gin.Context) {

	id, id_err := strconv.Atoi(c.Query("id"))
	if id_err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	amount, amount_err := strconv.ParseFloat(c.Query("amount"), 64)
	if amount_err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid amount param")
		return
	}

	write_downs_err := h.services.Operation.WriteDowns(id, float64(amount))
	if write_downs_err != nil {
		switch write_downs_err.Error() {
		case `pq: new row for relation "history" violates check constraint "history_amount_check"`:
			newErrorResponse(c, http.StatusInternalServerError, "invalid amount")
		case `pq: new row for relation "users" violates check constraint "users_balance_check"`:
			newErrorResponse(c, http.StatusInternalServerError, "not enough money")
		case `pq: insert or update on table "history" violates foreign key constraint "history_sender_id_fkey"`:
			newErrorResponse(c, http.StatusInternalServerError, "id not found")
		case `pq: new row for relation "history" violates check constraint "history_sender_id_check"`:
			newErrorResponse(c, http.StatusInternalServerError, "id not found")
		default:
			newErrorResponse(c, http.StatusInternalServerError, write_downs_err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})

}

func (h *Handler) transfer(c *gin.Context) {

	sender_id, sender_err := strconv.Atoi(c.Query("sender_id"))
	if sender_err != nil || sender_id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid sender_id param")
		return
	}

	receiver_id, receiver_err := strconv.Atoi(c.Query("receiver_id"))
	if receiver_err != nil || receiver_id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid receiver_id param")
		return
	}

	if sender_id == receiver_id {
		newErrorResponse(c, http.StatusBadRequest, "equal sender_id and receiver_id values")
		return
	}

	amount, amount_err := strconv.ParseFloat(c.Query("amount"), 64)
	if amount_err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid amount param")
		return
	}

	transfer_err := h.services.Operation.Transfer(sender_id, receiver_id, float64(amount))
	if transfer_err != nil {
		switch transfer_err.Error() {
		case `pq: new row for relation "history" violates check constraint "history_amount_check"`:
			newErrorResponse(c, http.StatusInternalServerError, "invalid amount")
		case `pq: insert or update on table "history" violates foreign key constraint "history_sender_id_fkey"`:
			newErrorResponse(c, http.StatusInternalServerError, "sender_id not found")
		case `pq: new row for relation "users" violates check constraint "users_balance_check"`:
			newErrorResponse(c, http.StatusInternalServerError, "not enough money")
		case `pq: new row for relation "users" violates check constraint "users_id_check"`:
			newErrorResponse(c, http.StatusInternalServerError, "invalid id")
		default:
			newErrorResponse(c, http.StatusInternalServerError, transfer_err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})

}