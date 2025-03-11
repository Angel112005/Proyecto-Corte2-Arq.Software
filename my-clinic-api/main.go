package main

import (
	"fmt"
	"my-clinic-api/dependencies"
	"my-clinic-api/server"
)

func main() {
	// Inicializamos las dependencias
	appDeps, err := dependencies.InitializeDependencies()
	if err != nil {
		fmt.Println("Error initializing dependencies:", err)
		return
	}

	// Configuramos el router y para poder correr el servidor
	router := server.SetupRouter(appDeps.DoctorController, appDeps.AppointmentController)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
