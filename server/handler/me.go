package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It's me",
	})
}
