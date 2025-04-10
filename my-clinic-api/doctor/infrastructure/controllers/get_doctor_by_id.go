package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	doctor, err := c.ListDoctorByID.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doctor"})
		return
	}

	if doctor == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	ctx.JSON(http.StatusOK, doctor)
}
