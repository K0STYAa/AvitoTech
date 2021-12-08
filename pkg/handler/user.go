package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	currency, currency_err := strconv.Atoi(c.DefaultQuery("currency", "0"))
	if currency_err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid currency param")
		return
	}

	user, err := h.services.User.GetById(id, currency)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
