package doctor

import "my-clinic-api/doctor/domain"

type CreateDoctor struct {
	repo doctor.Repository
}

func NewCreateDoctor(repo doctor.Repository) *CreateDoctor {
	return &CreateDoctor{repo: repo}
}

func (uc *CreateDoctor) Execute(d doctor.Doctor) error {
	return uc.repo.Save(d)
}
