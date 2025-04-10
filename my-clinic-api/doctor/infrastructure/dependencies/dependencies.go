package dependencies

import (
	"my-clinic-api/config"
	appDoctor "my-clinic-api/doctor/application"
	//Correción en la línea de abajo para la nueva ruta de la carpeta api -> controllers
	doctorAPI "my-clinic-api/doctor/infrastructure/controllers"
	doctorPersistence "my-clinic-api/doctor/infrastructure/persistence"
)

func InitializeDoctorController() (*doctorAPI.Controller, error) {
	// Conexión a base de datos
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}

	// Repositorio y casos de uso
	repo := doctorPersistence.NewDoctorMySQL(db)
	create := appDoctor.NewCreateDoctor(repo)
	list := appDoctor.NewListDoctors(repo)
	update := appDoctor.NewUpdateDoctor(repo)
	delete := appDoctor.NewDeleteDoctor(repo)
	findByID := appDoctor.NewListDoctorByID(repo)

	// Retornar controlador
	controller := doctorAPI.NewController(create, list, update, delete, findByID)
	return controller, nil
}
