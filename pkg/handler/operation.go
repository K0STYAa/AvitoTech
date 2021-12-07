package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) accrual(c *gin.Context) {

	id, id_err := strconv.Atoi(c.Query("id"))
	if id_err != nil {
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
	if id_err != nil {
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

}