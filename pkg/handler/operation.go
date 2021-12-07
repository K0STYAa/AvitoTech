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

	amount, amount_err := strconv.Atoi(c.Query("amount"))
	if amount_err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid amount param")
		return
	}

	accrual_err := h.services.Operation.Accrual(id, amount)
	if accrual_err != nil {
		newErrorResponse(c, http.StatusInternalServerError, accrual_err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) writedowns(c *gin.Context) {

	id, id_err := strconv.Atoi(c.Query("id"))
	if id_err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	amount, amount_err := strconv.Atoi(c.Query("amount"))
	if amount_err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid amount param")
		return
	}

	write_downs_err := h.services.Operation.WriteDowns(id, amount)
	if write_downs_err != nil {
		newErrorResponse(c, http.StatusInternalServerError, write_downs_err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
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

	amount, amount_err := strconv.Atoi(c.Query("amount"))
	if amount_err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid amount param")
		return
	}

	transfer_err := h.services.Operation.Transfer(sender_id, receiver_id, amount)
	if transfer_err != nil {
		newErrorResponse(c, http.StatusInternalServerError, transfer_err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}