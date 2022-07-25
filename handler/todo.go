package handler

import (
	"log"
	"net/http"
	"rest_api_golang_crud_sqlx/model"
	"strconv"

	"github.com/gin-gonic/gin"
)
//get Category
func (h *Handler) getAllCategoryH(c *gin.Context) {
	lists, err := h.service.GetAllCategoryService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, lists)
}

//get id Category
func (h *Handler) getCategoryByIdH(c *gin.Context) {
	id := c.Param("id")
	lists, err := h.service.GetByIdC(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, lists)
}

//post Category
func (h *Handler) CreateCategoryH(c *gin.Context) {
	var input model.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCategory, err := h.service.CreateC(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error",
		})
		return
	}
	c.JSON(http.StatusOK, newCategory)
}


//put Category

func (h *Handler) updateCategoryH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id param",
		})
		return
	}

	var input model.UpdateCategoryinput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id param",
		})
		return
	}

	if err := h.service.UpdateC(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}
//DELETE Category
func (h *Handler) deleteCategoryH(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteC(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}


//GET BOOK
func (h *Handler) getAllBookH(c *gin.Context) {
	lists, err := h.service.GetAll()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, lists)
}

//GET ID BOOK
func (h *Handler) getBookByIdH(c *gin.Context) {
	id := c.Param("id")
	lists, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, lists)
}

//PUT BOOK
func (h *Handler) updateBookH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id param",
		})
		return
	}

	var input model.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id param",
		})
		return
	}

	if err := h.service.Update(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}

//POST BOOK
func (h *Handler) createBookH(c *gin.Context) {
	var input model.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oshybka",
		})
		return
	}
	c.JSON(http.StatusOK, newBook)
}
//DELETE BOOK
func (h *Handler) deleteBookH(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}
