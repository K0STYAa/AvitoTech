package handler

import (
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
			user.POST("/", h.createUser)
			user.GET("/:id", h.getUserById)
			user.DELETE(":/id", h.deleteUser)
		}
		history := api.Group("/history")
		{
			history.GET("/:id", h.getHistoryById)
		}

		operation := api.Group("/")
		{
			operation.POST("accrual/:id", h.acrrual)
			operation.PUT("write-downs/:id", h.writedowns)
			operation.DELETE("transfer/:id", h.transfer)
		}
	}

	return router
}