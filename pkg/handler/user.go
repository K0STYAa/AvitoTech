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
	user, err2 := h.services.User.GetById(id, currency)
	if err2 != nil {
		newErrorResponse(c, http.StatusInternalServerError, err2.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
