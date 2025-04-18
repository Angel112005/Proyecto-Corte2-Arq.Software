package dependencies

import (
	"my-clinic-api/config"

	appAppointment "my-clinic-api/appointment/application"
	appointmentAPI "my-clinic-api/appointment/infrastructure/controllers"
	appointmentPersistence "my-clinic-api/appointment/infrastructure/persistence"
	rabbitAdapter "my-clinic-api/appointment/infrastructure/adapters"
)

func InitializeAppointmentController() (*appointmentAPI.Controller, error) {
	// 1. Conectar a la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}

	// 2. Inicializar RabbitMQ Adapter desde infraestructura (ahora acoplado a appointment)
	rabbit, err := rabbitAdapter.NewRabbitMQAdapter()
	if err != nil {
		return nil, err
	}

	// 3. Crear Repositorio
	repo := appointmentPersistence.NewAppointmentMySQL(db)

	// 4. Casos de uso con interfaces
	create := appAppointment.NewCreateAppointment(repo, rabbit) // ahora recibe repo y rabbit
	list := appAppointment.NewListAppointments(repo)
	update := appAppointment.NewUpdateAppointment(repo)
	delete := appAppointment.NewDeleteAppointment(repo)
	findByID := appAppointment.NewFindAppointmentByID(repo)

	// 5. Retornar el controlador con los casos de uso
	controller := appointmentAPI.NewController(create, list, update, delete, findByID)
	return controller, nil
}




















// package dependencies

// import (
// 	"my-clinic-api/config"
// 	appAppointment "my-clinic-api/appointment/application"
// 	// Corrección en la línea de abajo api -> controllers
// 	appointmentAPI "my-clinic-api/appointment/infrastructure/controllers"
// 	appointmentPersistence "my-clinic-api/appointment/infrastructure/persistence"
// 	rabbitmqInfra "my-clinic-api/services/rabbitmq/infrastructure"
// 	rabbitmqApp "my-clinic-api/services/rabbitmq/application"
// )

// func InitializeAppointmentController() (*appointmentAPI.Controller, error) {
// 	// Conexión a base de datos
// 	db, err := config.ConnectDB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Conexión a RabbitMQ
// 	rabbitMQ, err := rabbitmqInfra.NewRabbitMQ()
// 	if err != nil {
// 		return nil, err
// 	}
// 	publisher := rabbitmqApp.NewPublishMessageService(rabbitMQ)

// 	// Repositorio y casos de uso
// 	repo := appointmentPersistence.NewAppointmentMySQL(db)
// 	create := appAppointment.NewCreateAppointment(repo, publisher)
// 	list := appAppointment.NewListAppointments(repo)
// 	update := appAppointment.NewUpdateAppointment(repo)
// 	delete := appAppointment.NewDeleteAppointment(repo)
// 	findByID := appAppointment.NewFindAppointmentByID(repo)

// 	// Retornar controlador
// 	controller := appointmentAPI.NewController(create, list, update, delete, findByID)
// 	return controller, nil
// }
