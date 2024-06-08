package models

import (
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	SenderNumber string `gorm:"size:12"`
	RecipientNum string `gorm:"size:12"`
	IsActive     bool   `gorm:"default:true"`
	IsSenderBus  bool   `gorm:"default:false"`
}

type Message struct {
	gorm.Model
	ConversationID uint   `gorm:"index"`
	Content        string `gorm:"type:text"`
}
