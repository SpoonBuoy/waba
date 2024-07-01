package factory

import (
	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
	"github.com/SpoonBuoy/waba/usecases/medical"
)

type BusinessType string

const (
	MED   BusinessType = "medical"
	SALON BusinessType = "salon"
)

func NewActor(name string, details string, svc string, bus BusinessType) bookings.Actor {
	switch bus {
	case MED:
		return medical.NewDoc(name, NewService(svc, bus).(medical.MedicalService), details)
	default:
		return nil
	}
}

func NewService(name string, bus BusinessType) bookings.ActorService {
	switch bus {
	case MED:
		return medical.NewMedicalService(name)
	default:
		return nil
	}
}
