package models

import (
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Name                     string
	Email                    string `gorm:"unique"`
	Password                 string
	PhoneNumber              string `gorm:"unique"`
	VerificationDocumentType string
	VerificationDocument     string
	Contexts                 []Context            `gorm:"constraint:OnDelete:CASCADE"`
	WhatsappCreds            []WhatsappCredential `gorm:"constraint:OnDelete:CASCADE"`
	ExcludedMappings         []ExcludedMapping    `gorm:"constraint:OnDelete:CASCADE"`
}

type Context struct {
	gorm.Model
	BusinessID uint
	Content    string
	IsActive   bool `gorm:"default:false"`
}

type WhatsappCredential struct {
	gorm.Model
	BotToken       string
	BusinessID     uint
	WhatsappNumber string
	WabaID         string
}

type ExcludedMapping struct {
	gorm.Model
	BusinessID  uint   `gorm:"index"`
	PhoneNumber string `gorm:"size:12"`
}
