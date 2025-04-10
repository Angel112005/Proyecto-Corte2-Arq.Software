package controllers

import (
	appDoctor "my-clinic-api/doctor/application"
)

type Controller struct {
	CreateDoctor    *appDoctor.CreateDoctor
	ListDoctors     *appDoctor.ListDoctors
	UpdateDoctor    *appDoctor.UpdateDoctor
	DeleteDoctor    *appDoctor.DeleteDoctor
	ListDoctorByID  *appDoctor.ListDoctorByID
}

func NewController(
	cd *appDoctor.CreateDoctor,
	ld *appDoctor.ListDoctors,
	ud *appDoctor.UpdateDoctor,
	dd *appDoctor.DeleteDoctor,
	ldb *appDoctor.ListDoctorByID,
) *Controller {
	return &Controller{
		CreateDoctor:   cd,
		ListDoctors:    ld,
		UpdateDoctor:   ud,
		DeleteDoctor:   dd,
		ListDoctorByID: ldb,
	}
}
