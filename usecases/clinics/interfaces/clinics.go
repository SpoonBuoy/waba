package interfaces

type IClinicController interface {
	IBookerController
}

type IClinicService interface {
	IBookerService
}

type IClinicRepository interface {
	IBookerRepository
}
