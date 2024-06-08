package models

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name    string
	Details string
}

type Context struct {
	gorm.Model
	Ctx      string
	IsActive bool
}

type ExcludedNos struct {
	gorm.Model
}
