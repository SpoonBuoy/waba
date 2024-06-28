package interfaces

import (
	typeinterfaces "github.com/SpoonBuoy/waba/usecases/bookings/type-interfaces"
)

type IBookingsController interface {
	IBookerController
	IActorController
}

type IBookingsService interface {
	IBookerService
	IActorService
}

type IBookingsRepository interface {
	typeinterfaces.Actor
	typeinterfaces.Slot
	typeinterfaces.Appointment
	typeinterfaces.ActorServices
	typeinterfaces.BusinessController
}
