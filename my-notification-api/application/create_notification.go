package application

import (
	"encoding/json"
	"log"
	"my-notification-api/domain"
	repositoryapplication "my-notification-api/application/repository_application"
)

type CreateNotification struct {
	repo   domain.NotificationRepository
	rabbit repositoryapplication.IRabbit // ✅ Interfaz para desacoplar de RabbitMQ
}

func NewCreateNotification(repo domain.NotificationRepository, rabbit repositoryapplication.IRabbit) *CreateNotification {
	return &CreateNotification{
		repo:   repo,
		rabbit: rabbit,
	}
}

func (uc *CreateNotification) Execute(notification domain.Notification) error {
	// Guardar en la BD
	err := uc.repo.Save(notification)
	if err != nil {
		log.Println("❌ Error al guardar en MySQL:", err)
		return err
	}

	// Publicar evento en RabbitMQ
	event, _ := json.Marshal(notification)
	err = uc.rabbit.Publish("notificaciones.nuevas", event)
	if err != nil {
		log.Println("❌ Error al publicar en RabbitMQ:", err)
	}

	return nil
}








// package application

// import (
// 	"encoding/json"
// 	"log"
// 	"my-notification-api/domain"
// 	"my-notification-api/application/repository_application"
// )

// type CreateNotification struct {
// 	repo   domain.NotificationRepository
// 	rabbit repository_application.IRabbit
// }

// func NewCreateNotification(repo domain.NotificationRepository, rabbit repository_application.IRabbit) *CreateNotification {
// 	return &CreateNotification{repo: repo, rabbit: rabbit}
// }

// func (uc *CreateNotification) Execute(notification domain.Notification) error {
// 	if err := uc.repo.Save(notification); err != nil {
// 		return err
// 	}

// 	// Serializar y publicar evento
// 	event, _ := json.Marshal(notification)
// 	if err := uc.rabbit.Publish("notificaciones.nuevas", event); err != nil {
// 		log.Println("❌ Error publicando a RabbitMQ:", err)
// 	}

// 	return nil
// }







// package application

// import "my-notification-api/domain"

// type CreateNotification struct {
// 	repo domain.NotificationRepository
// }

// func NewCreateNotification(repo domain.NotificationRepository) *CreateNotification {
// 	return &CreateNotification{repo: repo}
// }

// func (uc *CreateNotification) Execute(notification domain.Notification) error {
// 	return uc.repo.Save(notification)
// }
