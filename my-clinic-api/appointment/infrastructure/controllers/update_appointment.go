package controllers

import (
	"net/http"
	"strconv"

	domainAppointment "my-clinic-api/appointment/domain"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Update(ctx *gin.Context) {
	var a domainAppointment.Appointment
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	a.ID = parsedID

	if err := c.updateAppointment.Execute(a); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update appointment"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"})
}
