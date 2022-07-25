package handler

import (
	"rest_api_golang_crud_sqlx/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.Default()

	book := router.Group("/")
	category := router.Group(("/"))
	{
		//BOOK
		book.POST("book", h.createBookH)
		book.GET("books", h.getAllBookH)
		book.GET("book/:id", h.getBookByIdH)
		book.DELETE("book/:id", h.deleteBookH)
		book.PUT("book/:id", h.updateBookH)

		//CATEGORY
		category.GET("categories", h.getAllCategoryH)
		category.POST("category", h.CreateCategoryH)
		category.GET("category/:id", h.getCategoryByIdH)
		category.PUT("category/:id", h.updateCategoryH)
		category.DELETE("category/:id", h.deleteBookH)

	}

	return router
}
