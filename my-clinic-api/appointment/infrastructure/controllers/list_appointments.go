package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (c *Controller) List(ctx *gin.Context) {
	appointments, err := c.listAppointments.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments"})
		return
	}
	ctx.JSON(http.StatusOK, appointments)
}
