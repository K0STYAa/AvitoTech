package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {

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
		}

		operation := api.Group("/")
		{
			operation.POST("accrual/:id", h.accrual)
			operation.POST("write-downs/:id", h.writedowns)
			operation.POST("transfer/:id", h.transfer)
		}
	}

	return router
}