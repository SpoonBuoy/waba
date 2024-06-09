package db

import (
	"fmt"
	"os"

	"github.com/SpoonBuoy/waba/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Client *gorm.DB
}

func NewDb(host string, port uint, dbname string, user string, password string) *Db {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect to db %v", err.Error())
		os.Exit(1)
	}
	return &Db{Client: db}
}

func (db *Db) Migrate() {
	err := db.Client.AutoMigrate(
		&models.Business{},
		&models.Context{},
		&models.ExcludedMapping{},
		&models.Message{},
		&models.WhatsappCredential{},
	)
	if err != nil {
		fmt.Printf("migrations failed %v", err.Error())
		os.Exit(1)
	}
}
