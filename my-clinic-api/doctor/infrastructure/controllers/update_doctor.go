package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domainDoctor "my-clinic-api/doctor/domain"
)

func (c *Controller) Update(ctx *gin.Context) {
	var d domainDoctor.Doctor
	if err := ctx.ShouldBindJSON(&d); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	d.ID = convertID(id, ctx)

	if err := c.UpdateDoctor.Execute(d); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully"})
}
