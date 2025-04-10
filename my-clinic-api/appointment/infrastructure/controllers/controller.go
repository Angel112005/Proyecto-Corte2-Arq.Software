package controllers

import (
	appAppointment "my-clinic-api/appointment/application"
)

type Controller struct {
	createAppointment *appAppointment.CreateAppointment
	listAppointments  *appAppointment.ListAppointments
	updateAppointment *appAppointment.UpdateAppointment
	deleteAppointment *appAppointment.DeleteAppointment
	findByID          *appAppointment.FindAppointmentByID
}

func NewController(
	ca *appAppointment.CreateAppointment,
	la *appAppointment.ListAppointments,
	ua *appAppointment.UpdateAppointment,
	da *appAppointment.DeleteAppointment,
	findByID *appAppointment.FindAppointmentByID,
) *Controller {
	return &Controller{
		createAppointment: ca,
		listAppointments:  la,
		updateAppointment: ua,
		deleteAppointment: da,
		findByID:          findByID,
	}
}
