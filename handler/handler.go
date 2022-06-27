package handler

import (
	"test_project/service"

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

	todo := router.Group("/")
	{
		todo.GET("todos", h.getAllLists)
		todo.POST("todo", h.createList)
		todo.GET("todo/:id", h.getListById)
		todo.DELETE("todo/:id", h.deleteList)
		todo.PUT("todo/:id", h.updateList)
	}
	return router
}
type UpdateListInput struct {
	Description *string `json:"description"`
}

