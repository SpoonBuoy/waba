package bookings

// manages all the operations for business
type BusinessRepo interface {
	AddActor(Actor, int)
	GetActor(int) Actor
	GetAllActors(int) []Actor
	GetAllAppointments(int) []Appointment
	GetAllServices(int) []ActorService
	AddService(ActorService, int)
	GetAllSlots(int) []Slot
	GetSlot(int) Slot
	GetService(int) ActorService
}
