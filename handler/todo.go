package handler

import (
	"net/http"
	"strconv"
	"test_project/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllLists(c *gin.Context) {
	lists, err := h.services.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, lists)
}
func (h *Handler) updateList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id param",
		})
		return
	}

	var input model.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id param",
		})
		return
	}

	if err := h.services.UpdateList(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated",
	})
}

func (h *Handler) getListById(c *gin.Context) {
	id := c.Param("id")
	lists, err := h.services.GetListById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "StatusInternalServerError",
		})
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (h *Handler) createList(c *gin.Context) {
	var input model.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo, err := h.services.CreateList(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oshybj",
		})
		return
	}
	c.JSON(http.StatusOK, newTodo)
}

func (h *Handler) deleteList(c *gin.Context) {
	id := c.Param("id")
	err := h.services.DeleteList(id)
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
