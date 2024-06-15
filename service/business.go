package service

import (
	"fmt"

	"github.com/SpoonBuoy/waba/dto"
	"github.com/SpoonBuoy/waba/models"
	"github.com/SpoonBuoy/waba/repository"
)

type BusinessService struct {
	businessRepo repository.BusinessRepo
}

func NewBusinessService(br repository.BusinessRepo) *BusinessService {
	return &BusinessService{
		businessRepo: br,
	}
}
func (bs *BusinessService) CreateContext(c dto.CreateCtxReq, bid uint) models.Context {
	//get business id from token
	var bid uint = 1
	context := models.Context{
		Content:    c.Content,
		IsActive:   false,
		BusinessID: bid,
	}
	cCreated := bs.businessRepo.AddContext(context)

	return cCreated
}
func (bs *BusinessService) GetActiveContext(id uint) (*models.Context, error) {
	//gets the active context from business Id
	ctx := bs.businessRepo.GetActiveContext(id)
	if ctx != nil {
		return ctx, nil
	}
	return nil, fmt.Errorf("no active context found with id %d", id)
}
func (bs *BusinessService) SetActiveContext(c dto.SwitchActiveCtxReq, bid uint) models.Context {
	//get bid from token
	var bid uint = 1
	res := bs.businessRepo.UseContext(c.ContextId, bid)
	return *res
}
func (bs *BusinessService) GetAllContexts(id uint) []models.Context {
	//gets all the context from business ID
	b := bs.businessRepo.GetAllContexts(id)
	return b
}
func (bs *BusinessService) GetBusiness(id uint) (*models.Business, error) {
	//get business details
	bus := bs.businessRepo.GetBusiness(id)

	if bus != nil {
		return bus, nil
	}
	return nil, fmt.Errorf("no business found with id %d", id)
}
func (bs *BusinessService) CreateBusiness(bus dto.CreateBusinessReq) models.Business {
	//fill b with bus dto
	b := models.Business{
		Name:        bus.Name,
		Email:       bus.Email,
		Password:    bus.Password,
		PhoneNumber: bus.Phone,
	}
	//creates new business
	bCreated := bs.businessRepo.AddBusiness(b)
	return bCreated
}
func (bs *BusinessService) ValidateBusiness(creds dto.LoginReq) (*models.Business, error) {
	b := bs.businessRepo.ValidateBusiness(creds.Name, creds.Password)
	if b.ID < 0 {
		return nil, fmt.Errorf("invalid credentials")
	}
	return &b, nil
}

func (bs *BusinessService) GetExcludedNos(id uint) {
	//gets excluded nos
}
func (bs *BusinessService) SetExcludedNos() {
	//sets excluded nos
}
func (bs *BusinessService) GetWabaId() {
	//get waba id
}
