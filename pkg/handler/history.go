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

	sort := c.DefaultQuery("sort", "date")
	typeSort := c.DefaultQuery("type", "inc")
	limit := c.DefaultQuery("limit", "ALL")

	offset, err5 := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err5 != nil || offset < 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid offset param")
		return
	}

	hist, err6 := h.services.History.GetById(id, sort, typeSort, limit, offset)
	if err6 != nil {
		newErrorResponse(c, http.StatusInternalServerError, err6.Error())
		return
	}

	c.JSON(http.StatusOK, getHistResponce{
		Data: hist,
	})

}