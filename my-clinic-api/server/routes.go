package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	doctorAPI "my-clinic-api/doctor/infrastructure/api" 
	appointmentAPI "my-clinic-api/appointment/infrastructure/api" 
)


func SetupRouter(
	doctorController *doctorAPI.Controller,
	appointmentController *appointmentAPI.Controller,
) *gin.Engine {
	r := gin.Default()

	// Configurar CORS para permitir solicitudes del frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	
	r.POST("/doctors", doctorController.Create)
	r.GET("/doctors", doctorController.List)
	r.GET("/doctors/:id", doctorController.GetByID)
	r.PUT("/doctors/:id", doctorController.Update)
	r.DELETE("/doctors/:id", doctorController.Delete)

	
	r.POST("/appointments", appointmentController.Create)
	r.GET("/appointments", appointmentController.List)
	r.GET("/appointments/:id", appointmentController.GetByID)
	r.PUT("/appointments/:id", appointmentController.Update)
	r.DELETE("/appointments/:id", appointmentController.Delete)


	return r
}
