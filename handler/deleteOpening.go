package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletepeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Opening API",
	})
}
