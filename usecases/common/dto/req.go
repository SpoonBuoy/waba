package dto

import "github.com/SpoonBuoy/waba/usecases/factory"

type CreateActorReq struct {
	Name         string
	Details      string
	Service      string
	BusinessType factory.BusinessType
}
type CreateAppointmentReq struct{}
type CreateBusinessReq struct{}
