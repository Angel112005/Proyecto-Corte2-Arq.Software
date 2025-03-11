package appointment

import "my-clinic-api/appointment/domain"

type FindAppointmentByID struct {
	repo appointment.Repository
}

func NewFindAppointmentByID(repo appointment.Repository) *FindAppointmentByID {
	return &FindAppointmentByID{repo: repo}
}

func (f *FindAppointmentByID) Execute(id int) (*appointment.Appointment, error) {
	return f.repo.FindByID(id)
}
