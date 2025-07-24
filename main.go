package main

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"bzdev/handlers"
	"bzdev/middleware"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	r := gin.Default()

	r.POST("/login", middleware.AuthMiddleware(),handlers.LoginHandler)
	r.POST("/signup", middleware.AuthMiddleware(),handlers.SignupHandler)
	r.POST("/maketicket", middleware.AuthMiddleware(),handlers.MakeTicketHandler)
	r.Run(":" + port)
}
