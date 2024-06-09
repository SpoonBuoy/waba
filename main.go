package main

import (
	"fmt"
	"log"

	"github.com/SpoonBuoy/waba/controllers"
	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/repository"
	"github.com/SpoonBuoy/waba/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	dbUser      string
	dbName      string
	dbPass      string
	dbPort      int
	dbHost      string
	Db          *db.Db
	AutoMigrate bool

	busRepo        = repository.NewBusinessRepo(*Db)
	busServ        = service.NewBusinessService(*busRepo)
	llmServ        = service.NewLlmService()
	wabaService    = service.NewWabaService(*busServ, *llmServ)
	chatController = controllers.NewChatController(wabaService)
)

func ReadConfig() {
	viper.SetConfigName("app")  // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	viper.AddConfigPath("./")   // call multiple times to add many search paths

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbPort = viper.GetInt("db.port")
	dbHost = viper.GetString("db.host")
	dbUser = viper.GetString("db.user")
	dbName = viper.GetString("db.name")
	dbPass = viper.GetString("db.password")
	AutoMigrate = viper.GetBool("auto_migrate")
	fmt.Printf("DB config with host %s and port %s", dbHost, dbPort)

}
func init() {
	ReadConfig()
	Db = db.NewDb(dbHost, uint(dbPort), dbName, dbUser, dbPass)
	if AutoMigrate {
		Db.Migrate()
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/verify", chatController.Verify)
	r.POST("/verify", chatController.Listen)
	r.Run(":9000")
}
