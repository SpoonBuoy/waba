package dto

import "github.com/SpoonBuoy/waba/usecases/factory"

type CreateActorReq struct {
	Name         string
	Details      string
	Service      string
	BusinessType factory.BusinessType
}
type BookSlotReq struct {
	SlotId       int
	BusinessType factory.BusinessType
}
type CreateServiceReq struct {
	Name         string
	BusinessType factory.BusinessType
}
type CreateAppointmentReq struct{}
type CreateBusinessReq struct{}
