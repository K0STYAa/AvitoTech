package handler

import (
	"github.com/K0STYAa/AvitoTech"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getHistResponce struct {
	Data []AvitoTech.History `json:"data"`
}

func (h *Handler) getHistoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	hist, err := h.services.History.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getHistResponce{
		Data: hist,
	})

}