package medical

import (
	"time"

	"gorm.io/gorm"
)

type Clinic struct {
	gorm.Model
	Name         string
	Appointments []Booking
	Doctors      []Doc
}

type Doc struct {
	gorm.Model
	Name         string
	Slots        []Slot
	Appointments []Booking
}

type Booking struct {
	gorm.Model
	Doc  Doc
	User string
	Slot Slot
}
type Slot struct {
	gorm.Model
	From        time.Time
	To          time.Time
	IsAvailable bool
}

// helper funcs
func docById(docs []Doc, id int) Doc {
	return docs[0]
}
