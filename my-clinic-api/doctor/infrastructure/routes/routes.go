package routes

import (
	"github.com/gin-gonic/gin"
	doctorAPI "my-clinic-api/doctor/infrastructure/controllers"
)

// RegisterDoctorRoutes configura las rutas para la entidad "doctor"
func RegisterDoctorRoutes(router *gin.Engine, doctorController *doctorAPI.Controller) {
	doctors := router.Group("/doctors")
	{
		doctors.POST("", doctorController.Create)
		doctors.GET("", doctorController.List)
		doctors.GET("/:id", doctorController.GetByID)
		doctors.PUT("/:id", doctorController.Update)
		doctors.DELETE("/:id", doctorController.Delete)
	}
}
