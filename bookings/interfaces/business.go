package bookings

// manages all the operations for business
type Business interface {
	AddActor(Actor, int)
	GetActor(int) Actor
	GetAllActors(int) []Actor
	GetAllAppointments(int) []Appointment
	GetAllServices(int) []ActorService
	AddService(ActorService, int)
}
