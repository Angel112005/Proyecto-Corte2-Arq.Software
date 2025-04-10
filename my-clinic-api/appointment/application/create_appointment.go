package appointment

import (
	"encoding/json"
	"log"
	"my-clinic-api/appointment/domain"
	"my-clinic-api/services/rabbitmq/application"
)

type CreateAppointment struct {
	repo          appointment.Repository
	messageBus    *application.PublishMessageService
}

func NewCreateAppointment(repo appointment.Repository, mb *application.PublishMessageService) *CreateAppointment {
	return &CreateAppointment{repo: repo, messageBus: mb}
}

func (uc *CreateAppointment) Execute(a appointment.Appointment) error {
	// Guardar en MySQL
	err := uc.repo.Save(a)
	if err != nil {
		log.Println("❌ Error al guardar la cita en MySQL:", err)
		return err
	}

	// Serializar evento
	event, _ := json.Marshal(a)

	// Enviar evento a RabbitMQ
	uc.messageBus.Execute("citas.nuevas", event)

	return nil
}

















// package appointment

// import (
// 	"encoding/json"
// 	"log"

// 	"my-clinic-api/appointment/domain"
// 	rabbitmq "my-clinic-api/services/rabbitmq/domain" // 👈 solo importa la interfaz
// )

// type CreateAppointment struct {
// 	repo       appointment.Repository
// 	messageBus rabbitmq.IMessageBus // ✅ usa la interfaz, no la implementación concreta
// }

// func NewCreateAppointment(repo appointment.Repository, bus rabbitmq.IMessageBus) *CreateAppointment {
// 	return &CreateAppointment{
// 		repo:       repo,
// 		messageBus: bus,
// 	}
// }

// func (uc *CreateAppointment) Execute(a appointment.Appointment) error {
// 	// 1. Guardar la cita
// 	err := uc.repo.Save(a)
// 	if err != nil {
// 		log.Println("❌ Error al guardar en la base de datos:", err)
// 		return err
// 	}

// 	// 2. Enviar el mensaje al bus
// 	event, _ := json.Marshal(a)
// 	uc.messageBus.Publish("citas.nuevas", event) // 👈 usando la interfaz, sin importar qué hay detrás

// 	return nil
// }
























// package appointment

// import (
//     "log"
//     "my-clinic-api/appointment/domain"
// )

// // CreateAppointment usa una interfaz en vez de infraestructura directa
// type CreateAppointment struct {
//     repo      appointment.Repository
//     publisher appointment.EventPublisher
// }

// // Constructor
// func NewCreateAppointment(repo appointment.Repository, publisher appointment.EventPublisher) *CreateAppointment {
//     return &CreateAppointment{repo: repo, publisher: publisher}
// }

// // Ejecutar la creación de la cita
// func (uc *CreateAppointment) Execute(a appointment.Appointment) error {
//     err := uc.repo.Save(a)
//     if err != nil {
//         log.Println("❌ Error al guardar la cita en MySQL:", err)
//         return err
//     }

//     // Enviar el evento usando la interfaz en lugar de una implementación concreta
//     evento := appointment.CitaEvento{
//         ID:              a.ID,
//         PatientName:     a.PatientName,
//         DoctorID:        a.DoctorID,
//         AppointmentDate: a.AppointmentDate,
//         Status:          a.Status,
//     }

//     err = uc.publisher.Publish(evento)
//     if err != nil {
//         log.Println("❌ Error al publicar evento en RabbitMQ:", err)
//     }

//     return nil
// }

