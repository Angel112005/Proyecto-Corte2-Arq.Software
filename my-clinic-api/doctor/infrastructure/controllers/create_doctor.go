package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domainDoctor "my-clinic-api/doctor/domain"
)

func (c *Controller) Create(ctx *gin.Context) {
	var d domainDoctor.Doctor
	if err := ctx.ShouldBindJSON(&d); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.CreateDoctor.Execute(d); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Doctor created successfully"})
}
