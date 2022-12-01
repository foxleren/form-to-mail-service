package handler

import (
	"form-to-mail-service/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	api := router.Group("/api")
	{
		api.POST("/", h.sendEmail)
	}

	return router
}
