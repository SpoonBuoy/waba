package models

import (
	"time"
)

type Business struct {
	ID                       uint                 `gorm:"primaryKey;autoIncrement"`
	Name                     string               `gorm:"size:255"`
	Email                    string               `gorm:"size:255;unique"`
	Password                 string               `gorm:"size:255"`
	PhoneNumber              string               `gorm:"size:12;unique"`
	CreatedAt                time.Time            `gorm:"not null"`
	VerificationDocumentType string               `gorm:"type:text"`
	VerificationDocument     string               `gorm:"type:text"`
	Contexts                 []Context            `gorm:"constraint:OnDelete:CASCADE"`
	WhatsappCreds            []WhatsappCredential `gorm:"constraint:OnDelete:CASCADE"`
	ExcludedMappings         []ExcludedMapping    `gorm:"constraint:OnDelete:CASCADE"`
}

type Context struct {
	ContextID      uint      `gorm:"primaryKey;autoIncrement"`
	BusinessID     uint      `gorm:"index"`
	ContextContent string    `gorm:"type:text"`
	IsActive       bool      `gorm:"default:false"`
	CreatedAt      time.Time `gorm:"not null"`
}

type WhatsappCredential struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	BotToken       string `gorm:"size:255"`
	BusinessID     uint   `gorm:"index"`
	WhatsappNumber string `gorm:"size:12"`
}

type ExcludedMapping struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	BusinessID  uint   `gorm:"index"`
	PhoneNumber string `gorm:"size:12"`
}
