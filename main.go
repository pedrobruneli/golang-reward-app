package main

import (
	"log"
	"rewards/docs"
	"rewards/internal/database"
	"rewards/internal/usecases/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func main() {
	initEnv()
	router := gin.Default()
	generateSwagger(router)
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	registerUserRoutes(router, db)
	router.Run(":8081")
}

func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func generateSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Rewards API"
	docs.SwaggerInfo.Description = "This is a simple API for rewards"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8081/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))
}

func registerUserRoutes(router *gin.Engine, db *gorm.DB) {
	userController := users.NewUserController(users.NewUserService(users.NewUserRepository(db)))
	userController.RegisterUserRoutes(router, db)
}
