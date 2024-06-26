package repository

import (
	"log"

	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/models"
	"github.com/ahsmha/gashtools/logger"
)

type BusinessRepo struct {
	db     db.Db
	logger logger.ILogWriter
}

func NewBusinessRepo(gdb db.Db, logger logger.ILogWriter) *BusinessRepo {
	return &BusinessRepo{
		db:     gdb,
		logger: logger,
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

func (br *BusinessRepo) UpdateActiveContext(bid uint, ctxId uint, setActive bool) *models.Context {
	var ctx *models.Context = nil
	if ctxId > 0 {
		err := br.db.Client.Model(&models.Context{}).Where("business_id = ? AND id = ?", bid, ctxId).Update("is_active", setActive).First(&ctx).Error
		HanldeDbErr(err, "SetActiveContext")
	} else {
		err := br.db.Client.Model(&models.Context{}).Where("business_id = ? AND is_active = ?", bid, !setActive).Update("is_active", setActive).First(&ctx).Error
		HanldeDbErr(err, "SetActiveContext")
	}
	return ctx

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
	ctx := br.UpdateActiveContext(bid, 0, false)
	ctx = br.UpdateActiveContext(bid, cid, true)
	//
	return ctx
}
