package handler

import (
	"github.com/K0STYAa/AvitoTech/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		user := api.Group("/users")
		{
			user.GET("/:id", h.getUserById)
		}

		history := api.Group("/history")
		{
			history.GET("/:id", h.getHistoryById)
			history.GET("/count/:id", h.getHistoryCountById)
		}

		operation := api.Group("/")
		{
			operation.POST("accrual", h.accrual)
			operation.POST("write-downs", h.writedowns)
			operation.POST("transfer", h.transfer)
		}
	}

	return router
}