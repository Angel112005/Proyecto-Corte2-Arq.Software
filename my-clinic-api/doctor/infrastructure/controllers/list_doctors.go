package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) List(ctx *gin.Context) {
	doctors, err := c.ListDoctors.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doctors"})
		return
	}
	ctx.JSON(http.StatusOK, doctors)
}
