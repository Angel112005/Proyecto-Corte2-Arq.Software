package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	// Configuración
	"my-clinic-api/config"

	// Rutas
	appointmentRoutes "my-clinic-api/appointment/infrastructure/routes"
	doctorRoutes "my-clinic-api/doctor/infrastructure/routes"

	// Dependencias por entidad
	appointmentDeps "my-clinic-api/appointment/infrastructure/dependencies"
	doctorDeps "my-clinic-api/doctor/infrastructure/dependencies"
)

func main() {
	// 📦 Inicializar controlador de Doctor
	doctorController, err := doctorDeps.InitializeDoctorController()
	if err != nil {
		fmt.Println("❌ Error al inicializar DoctorController:", err)
		return
	}

	// 📦 Inicializar controlador de Appointment
	appointmentController, err := appointmentDeps.InitializeAppointmentController()
	if err != nil {
		fmt.Println("❌ Error al inicializar AppointmentController:", err)
		return
	}

	// 🚀 Crear router de Gin y configurar CORS
	router := gin.Default()
	router.Use(config.ConfigureCORS())

	// 🧭 Registrar rutas de cada módulo
	doctorRoutes.RegisterDoctorRoutes(router, doctorController)
	appointmentRoutes.RegisterAppointmentRoutes(router, appointmentController)

	// 🌐 Iniciar servidor
	fmt.Println("✅ Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}


















// package main

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"

// 	"my-clinic-api/config"
// 	appointmentRoutes "my-clinic-api/appointment/infrastructure/routes"
// 	doctorRoutes "my-clinic-api/doctor/infrastructure/routes"

// 	appointmentDeps "my-clinic-api/appointment/infrastructure/dependencies"
// 	doctorDeps "my-clinic-api/doctor/infrastructure/dependencies"
// )

// func main() {
// 	// Inicializar controladores
// 	doctorController, err := doctorDeps.InitializeDoctorController()
// 	if err != nil {
// 		fmt.Println("❌ Error inicializando DoctorController:", err)
// 		return
// 	}

// 	appointmentController, err := appointmentDeps.InitializeAppointmentController()
// 	if err != nil {
// 		fmt.Println("❌ Error inicializando AppointmentController:", err)
// 		return
// 	}

// 	// Configurar router
// 	router := gin.Default()
// 	router.Use(config.ConfigureCORS())

// 	// Registrar rutas
// 	appointmentRoutes.RegisterAppointmentRoutes(router, appointmentController)
// 	doctorRoutes.RegisterDoctorRoutes(router, doctorController)

// 	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
// 	router.Run(":8080")
// }






