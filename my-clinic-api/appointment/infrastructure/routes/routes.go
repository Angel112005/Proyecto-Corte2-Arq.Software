package routes

import (
	"github.com/gin-gonic/gin"
	appointmentAPI "my-clinic-api/appointment/infrastructure/controllers"
)

// RegisterAppointmentRoutes configura las rutas para la entidad "appointment"
func RegisterAppointmentRoutes(router *gin.Engine, controller *appointmentAPI.Controller) {
	appointments := router.Group("/appointments")
	{
		appointments.POST("", controller.Create)
		appointments.GET("", controller.List)
		appointments.GET("/:id", controller.GetByID)
		appointments.PUT("/:id", controller.Update)
		appointments.DELETE("/:id", controller.Delete)
	}
}
