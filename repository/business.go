package repository

import (
	"log"

	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/models"
)

type BusinessRepo struct {
	db db.Db
}

func NewBusinessRepo(gdb db.Db) *BusinessRepo {
	return &BusinessRepo{
		db: gdb,
	}
}

func HanldeDbErr(err error, where string) {
	if err != nil {
		log.Printf("layer repo : error in %s with err : %v", where, err.Error())
	}
}

func (br *BusinessRepo) GetBusiness(id uint) *models.Business {
	var bus *models.Business = nil
	err := br.db.Client.Where("ID = ?", id).First(&bus).Error

	HanldeDbErr(err, "GetBusiness")
	return bus
}

func (br *BusinessRepo) GetAllContexts(bid uint) []models.Context {
	var ctxs []models.Context
	err := br.db.Client.Where("business_id = ?", bid).Find(&ctxs).Error
	HanldeDbErr(err, "GetAllContexts")
	return ctxs
}
func (br *BusinessRepo) AddContext(c models.Context) models.Context {
	err := br.db.Client.Create(&c).Error

	HanldeDbErr(err, "AddContext")
	return c
}

func (br *BusinessRepo) AddBusiness(bus models.Business) models.Business {
	err := br.db.Client.Create(&bus).Error

	HanldeDbErr(err, "AddBusiness")
	return bus
}
func (br *BusinessRepo) ValidateBusiness(name string, password string) models.Business {
	var bus models.Business
	err := br.db.Client.Where("email= ? AND password = ?", name, password).First(&bus).Error

	HanldeDbErr(err, "AddBusiness")
	return bus
}
func (br *BusinessRepo) GetWabaId(bid uint) string {
	var cred models.WhatsappCredential
	err := br.db.Client.Where("business_id = ?", bid).First(&cred).Error

	HanldeDbErr(err, "GetWabaID")
	return cred.WabaID
}

func (br *BusinessRepo) AddWabaCreds(cred models.WhatsappCredential) uint {
	err := br.db.Client.Create(&cred).Error

	HanldeDbErr(err, "AddWabaCreds")
	return cred.ID
}

func (br *BusinessRepo) GetActiveContext(bid uint) *models.Context {
	var ctx *models.Context = nil
	err := br.db.Client.Where("business_id = ? AND is_active = ?", bid, true).First(&ctx).Error

	HanldeDbErr(err, "GetActiveContext")
	return ctx
}
func (br *BusinessRepo) GetContext(cid uint, bid uint) *models.Context {
	var ctx *models.Context = nil
	err := br.db.Client.Where("business_id = ? AND ID = ?", bid, cid).First(&ctx).Error

	HanldeDbErr(err, "GetContext")
	return ctx
}
func (br *BusinessRepo) UseContext(cid uint, bid uint) *models.Context {
	alreadyActive := br.GetActiveContext(bid)

	if alreadyActive != nil {
		//update it
		alreadyActive.IsActive = false
		br.db.Client.Save(&alreadyActive)
	}
	//make it active
	ctx := br.GetContext(cid, bid)

	if ctx != nil {
		//found ctx
		ctx.IsActive = true
		br.db.Client.Save(&ctx)
	}

	return ctx

}
