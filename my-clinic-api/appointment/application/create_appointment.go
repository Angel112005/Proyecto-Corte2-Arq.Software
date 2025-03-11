package appointment

import (
	"log"
	"my-clinic-api/appointment/domain"
	persistence "my-clinic-api/appointment/infrastructure/persistence"
	// "time"
)

type CreateAppointment struct {
	repo appointment.Repository
}

func NewCreateAppointment(repo appointment.Repository) *CreateAppointment {
	return &CreateAppointment{repo: repo}
}

// func (uc *CreateAppointment) Execute(a appointment.Appointment) error {
// 	return uc.repo.Save(a)
// }

func (uc *CreateAppointment) Execute(a appointment.Appointment) error {
	// 1️⃣ Guardar la cita en la base de datos
	err := uc.repo.Save(a)
	if err != nil {
		log.Println("❌ Error al guardar la cita en MySQL:", err)
		return err
	}

	// 2️⃣ Enviar la cita a RabbitMQ
	evento := persistence.CitaEvento{
		ID:              a.ID,
		PatientName:     a.PatientName,
		DoctorID:        a.DoctorID,
		AppointmentDate: a.AppointmentDate,
		Status:          a.Status,
	}

	persistence.PublicarCitaEvento(evento)

	return nil
}
