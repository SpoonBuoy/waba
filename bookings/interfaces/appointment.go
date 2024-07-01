package bookings

// appointment is a contract between actor and user
type Appointment interface {
	Book() Appointment
}
