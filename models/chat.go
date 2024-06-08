package models

import "time"

type Conversation struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	SenderNumber string    `gorm:"size:12"`
	RecipientNum string    `gorm:"size:12"`
	IsActive     bool      `gorm:"default:true"`
	IsSenderBus  bool      `gorm:"default:false"`
	CreatedAt    time.Time `gorm:"not null"`
}

type Message struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	ConversationID uint      `gorm:"index"`
	Content        string    `gorm:"type:text"`
	CreatedAt      time.Time `gorm:"not null"`
}
