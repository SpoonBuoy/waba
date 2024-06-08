package repository

import (
	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/models"
)

type BusinessRepo struct {
	db db.Db
}

func NewBusinessRepo() *BusinessRepo {
	return &BusinessRepo{}
}

func (br *BusinessRepo) GetBusiness(id uint) models.Business {
	var bus models.Business
	err := br.db.Client.Where("ID = ?", id).First(&bus).Error

	if err != nil {
		//handle db err
	}
	return bus
}

func (br *BusinessRepo) GetAllContexts(bid uint) []models.Context {
	var ctxs []models.Context
	err := br.db.Client.Where("BusinessID = ?", bid).Find(&ctxs).Error
	if err != nil {
		//handle db err
	}
	return ctxs
}

func (br *BusinessRepo) AddBusiness(bus models.Business) uint {
	err := br.db.Client.Create(&bus).Error

	if err != nil {
		//handle err
	}
	return bus.ID
}

func (br *BusinessRepo) GetWabaId(bid uint) string {
	var cred models.WhatsappCredential
	err := br.db.Client.Where("BusinessID = ?", bid).First(&cred).Error

	if err != nil {
		//handle err
	}
	return cred.WabaID
}

func (br *BusinessRepo) AddWabaCreds(cred models.WhatsappCredential) uint {
	err := br.db.Client.Create(&cred).Error

	if err != nil {
		//handle err
	}
	return cred.ID
}
