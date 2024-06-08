package service

type BusinessService struct {
}

func NewBusinessService() *BusinessService {
	return &BusinessService{}
}

func (bs *BusinessService) GetActiveContext(id uint) {
	//gets the active context from business Id
}
func (bs *BusinessService) GetAllContexts(id uint) {
	//gets all the context from business ID
}
func (bs *BusinessService) GetBusiness(id uint) {
	//get business details
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
