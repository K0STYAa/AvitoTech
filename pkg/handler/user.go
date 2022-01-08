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

	currency := c.Query("currency")
	user, err := h.services.User.GetById(id, currency)
	if err != nil {
		switch err.Error() {
		case `sql: no rows in result set`:
			newErrorResponse(c, http.StatusInternalServerError, "user not found")
		default:
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
