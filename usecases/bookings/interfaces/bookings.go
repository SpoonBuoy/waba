package interfaces

type IBookingsController interface {
	IBookerController
}

type IBookingsService interface {
	IBookerService
}

type IClinicsRepository interface {
	IBookingRepository
}
