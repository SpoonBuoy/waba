package interfaces

type IBookingsController interface {
	IBookerController
	IActorController
}

type IBookingsService interface {
	IBookerService
	IActorService
}

type IClinicsRepository interface {
	IBookingRepository
	IClinicRepository
}
