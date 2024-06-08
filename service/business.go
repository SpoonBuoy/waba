package service

import "github.com/SpoonBuoy/waba/repository"

type BusinessService struct {
	businessRepo repository.BusinessRepo
}

func NewBusinessService() *BusinessService {
	return &BusinessService{}
}

func (bs *BusinessService) GetActiveContext(id uint) {
	//gets the active context from business Id
	// ctx := bs.businessRepo.GetActiveContext(id)
}
func (bs *BusinessService) GetAllContexts(id uint) {
	//gets all the context from business ID
	// b := bs.businessRepo.GetAllContexts(id)
}
func (bs *BusinessService) GetBusiness(id uint) {
	//get business details
	// bus := bs.businessRepo.GetBusiness(id)
}
func (bs *BusinessService) CreateBusiness() {
	//creates new business

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
