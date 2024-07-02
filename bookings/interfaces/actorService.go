package bookings

// type of service an actor provides in business
type ActorService interface {
	Type() string
	GetDetails() *ServiceDetails
}
type ServiceDetails struct {
	ServiceType string
	Name        string
}
