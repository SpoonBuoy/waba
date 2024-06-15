package main

import (
	"fmt"
	"log"

	"github.com/SpoonBuoy/waba/controllers"
	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/middleware"
	"github.com/SpoonBuoy/waba/repository"
	"github.com/SpoonBuoy/waba/service"
	"github.com/gin-contrib/cors"
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

	busRepo        *repository.BusinessRepo
	busServ        *service.BusinessService
	llmServ        *service.LlmService
	wabaService    *service.WabaService
	chatController *controllers.ChatController
	businessCtrl   *controllers.BusinessController
	authMiddleware *middleware.Auth
	jwtSecret      string
	llmChatUri     string
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
	AutoMigrate = viper.GetBool("db.auto_migrate")
	llmChatUri = viper.GetString("service.llm.chat_uri")
	jwtSecret = viper.GetString("jwt.secret")
	fmt.Printf("DB config with host %s and port %s", dbHost, dbPort)

}
func init() {
	ReadConfig()
	Db = db.NewDb(dbHost, uint(dbPort), dbName, dbUser, dbPass)
	if AutoMigrate {
		Db.Migrate()
	}
	busRepo = repository.NewBusinessRepo(*Db)
	busServ = service.NewBusinessService(*busRepo)
	llmServ = service.NewLlmService(llmChatUri)
	wabaService = service.NewWabaService(*busServ, *llmServ)
	chatController = controllers.NewChatController(wabaService)
	businessCtrl = controllers.NewBusinessController(*busServ, jwtSecret)
	authMiddleware = middleware.NewAuthMiddleware(jwtSecret)
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "https://ai.gasha.live", "https://api.gasha.live", "http://localhost", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},  // Including more methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Including more headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{

		account := api.Group("/account")

		account.POST("/onboard", businessCtrl.AddBusiness)
		account.POST("/login", businessCtrl.Login)

		business := api.Group("/business")
		business.Use(authMiddleware.Verify)

		business.GET("/verify", businessCtrl.VerifyAuth)
		business.GET("/contexts", businessCtrl.GetContexts)
		business.POST("/context", businessCtrl.AddContext)
		business.POST("/context/set", businessCtrl.SetActiveCtx)

		waba := api.Group("/waba")
		waba.GET("/event", chatController.Verify)
		waba.POST("/event", chatController.Listen)

	}
	r.Run(":9000")
}
