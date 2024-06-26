package interfaces

import "context"

type IBookerController interface {
	GetBookings(ctx context.Context)
}

type IBookerService interface {
	GetBookings(ctx context.Context)
}

type IBookerRepository interface {
	GetBookings(ctx context.Context)
}
