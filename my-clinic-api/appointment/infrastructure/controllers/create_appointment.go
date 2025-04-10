package controllers

import (
	"fmt"
	"net/http"
	// "strconv"

	domainAppointment "my-clinic-api/appointment/domain"
	"github.com/gin-gonic/gin"
)

func (c *Controller) Create(ctx *gin.Context) {
	var a domainAppointment.Appointment
	if err := ctx.ShouldBindJSON(&a); err != nil {
		fmt.Println("❌ Error en ShouldBindJSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("📌 Datos recibidos:", a)

	if err := c.createAppointment.Execute(a); err != nil {
		fmt.Println("❌ Error en createAppointment.Execute:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "✅ Appointment created and sent to RabbitMQ"})
}
