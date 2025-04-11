package main

import (
	"log"
	"my-notification-api/application"
	"my-notification-api/config"
	"my-notification-api/infrastructure/adapters"
	"my-notification-api/infrastructure/api"
	"my-notification-api/infrastructure/persistence"
	"my-notification-api/infrastructure/routes"
)

func main() {
	config.ConnectDatabase()

	// Conectar a RabbitMQ
	rabbit, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("❌ No se pudo inicializar RabbitMQAdapter: %v", err)
	}

	// Inyección de dependencias
	repo := persistence.NewNotificationMySQL(config.DB)
	useCase := application.NewCreateNotification(repo, rabbit)
	controller := api.NewNotificationController(useCase)

	
	r := server.SetupRouter(controller)
	r.Run(":8081")
}









// package main

// import (
// 	"my-notification-api/config"
// 	"my-notification-api/infrastructure/api"
// 	"my-notification-api/infrastructure/persistence"
// 	"my-notification-api/application"
// 	"my-notification-api/server"
// )

// func main() {
// 	config.ConnectDatabase()

// 	repo := persistence.NewNotificationMySQL(config.DB)
// 	useCase := application.NewCreateNotification(repo)
// 	controller := api.NewNotificationController(useCase)

// 	r := server.SetupRouter(controller)
// 	r.Run(":8081")
// }
