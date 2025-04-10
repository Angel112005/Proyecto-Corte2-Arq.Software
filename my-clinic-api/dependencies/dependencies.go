package dependencies

// import (
// 	"my-clinic-api/config"

// 	appAppointment "my-clinic-api/appointment/application"
// 	appointmentAPI "my-clinic-api/appointment/infrastructure/api"
// 	appointmentPersistence "my-clinic-api/appointment/infrastructure/persistence"
	
// 	appDoctor "my-clinic-api/doctor/application"
// 	doctorAPI "my-clinic-api/doctor/infrastructure/api"
// 	doctorPersistence "my-clinic-api/doctor/infrastructure/persistence"

// 	// rabbitmqDomain "my-clinic-api/services/rabbitmq/domain"
// 	rabbitmqApp "my-clinic-api/services/rabbitmq/application"
// 	rabbitmqInfra "my-clinic-api/services/rabbitmq/infrastructure"
// )

// type AppDependencies struct {
// 	DoctorController     *doctorAPI.Controller
// 	AppointmentController *appointmentAPI.Controller
// }

// func InitializeDependencies() (*AppDependencies, error) {
// 	// Conexión a la base de datos
// 	db, err := config.ConnectDB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Conectar a RabbitMQ desde el servicio
// 	rabbitMQ, err := rabbitmqInfra.NewRabbitMQ()
// 	if err != nil {
// 		return nil, err
// 	}
// 	publishMessageService := rabbitmqApp.NewPublishMessageService(rabbitMQ)


// 	// Inicializar repositorios
// 	doctorRepo := doctorPersistence.NewDoctorMySQL(db)
// 	appointmentRepo := appointmentPersistence.NewAppointmentMySQL(db)

// 	// Inicializar casos de uso
// 	createDoctorUseCase := appDoctor.NewCreateDoctor(doctorRepo)
// 	listDoctorsUseCase := appDoctor.NewListDoctors(doctorRepo)
// 	updateDoctorUseCase := appDoctor.NewUpdateDoctor(doctorRepo)
// 	deleteDoctorUseCase := appDoctor.NewDeleteDoctor(doctorRepo)
// 	listDoctorByIDUseCase := appDoctor.NewListDoctorByID(doctorRepo)


// 	createAppointmentUseCase := appAppointment.NewCreateAppointment(appointmentRepo, publishMessageService)
// 	listAppointmentsUseCase := appAppointment.NewListAppointments(appointmentRepo)
// 	updateAppointmentUseCase := appAppointment.NewUpdateAppointment(appointmentRepo)
// 	deleteAppointmentUseCase := appAppointment.NewDeleteAppointment(appointmentRepo)
// 	findAppointmentByIDUseCase := appAppointment.NewFindAppointmentByID(appointmentRepo) // Nuevo caso de uso


// 	// Inicializar controladores
// 	doctorController := doctorAPI.NewController(
// 		createDoctorUseCase,
// 		listDoctorsUseCase,
// 		updateDoctorUseCase,
// 		deleteDoctorUseCase,
// 		listDoctorByIDUseCase,
// 	)

// 	appointmentController := appointmentAPI.NewController(
// 		createAppointmentUseCase,
// 		listAppointmentsUseCase,
// 		updateAppointmentUseCase,
// 		deleteAppointmentUseCase,
// 		findAppointmentByIDUseCase, // Pasar el nuevo caso de uso

// 	)

// 	return &AppDependencies{
// 		DoctorController:     doctorController,
// 		AppointmentController: appointmentController,
// 	}, nil
// }
