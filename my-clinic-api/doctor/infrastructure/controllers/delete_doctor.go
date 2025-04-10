package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedID := convertID(id, ctx)

	if err := c.DeleteDoctor.Execute(parsedID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete doctor"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}

func convertID(id string, ctx *gin.Context) int {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		ctx.Abort()
	}
	return parsedID
}