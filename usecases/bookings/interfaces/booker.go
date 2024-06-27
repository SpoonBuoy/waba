package interfaces

import "context"

type IBookerController interface {
	GetBookings(ctx context.Context)
}

type IBookerService interface {
	GetBookings(ctx context.Context)
}

type IBookingRepository interface {
	GetBookings(ctx context.Context)
}
