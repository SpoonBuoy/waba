package db

import "gorm.io/gorm"

type Db struct {
	Client *gorm.DB
}

func NewDb() *Db {
	return &Db{}
}
