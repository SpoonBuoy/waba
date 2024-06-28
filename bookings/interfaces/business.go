package bookings

// manages all the operations for business
type BusinessController interface {
	AddActor(Actor)
	GetActor(int) Actor
	GetAllActors() []Actor
	GetAllAppointments() []Appointment
	GetAllServices() []ActorService
	AddService(ActorService)
}
