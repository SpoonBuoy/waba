package models

import (
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Name                     string               `gorm:"size:255"`
	Email                    string               `gorm:"size:255;unique"`
	Password                 string               `gorm:"size:255"`
	PhoneNumber              string               `gorm:"size:12;unique"`
	VerificationDocumentType string               `gorm:"type:text"`
	VerificationDocument     string               `gorm:"type:text"`
	Contexts                 []Context            `gorm:"constraint:OnDelete:CASCADE"`
	WhatsappCreds            []WhatsappCredential `gorm:"constraint:OnDelete:CASCADE"`
	ExcludedMappings         []ExcludedMapping    `gorm:"constraint:OnDelete:CASCADE"`
}

type Context struct {
	gorm.Model
	BusinessID uint   `gorm:"index"`
	Content    string `gorm:"type:text"`
	IsActive   bool   `gorm:"default:false"`
}

type WhatsappCredential struct {
	gorm.Model
	BotToken       string `gorm:"size:255"`
	BusinessID     uint   `gorm:"index"`
	WhatsappNumber string `gorm:"size:12"`
	WabaID         string
}

type ExcludedMapping struct {
	gorm.Model
	BusinessID  uint   `gorm:"index"`
	PhoneNumber string `gorm:"size:12"`
}
