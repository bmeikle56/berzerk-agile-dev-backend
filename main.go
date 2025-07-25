package main

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{
        "http://localhost:3000",  // dev
        "https://myfrontend.app", // production
    },
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
}))

	r.POST("/login", middleware.AuthMiddleware(),handlers.LoginHandler)
	r.POST("/signup", middleware.AuthMiddleware(),handlers.SignupHandler)
	r.POST("/make", middleware.AuthMiddleware(),handlers.MakeTicketHandler)
	r.POST("/update", middleware.AuthMiddleware(),handlers.UpdateStatusHandler)
	r.POST("/fetch", middleware.AuthMiddleware(),handlers.FetchTicketsHandler)
	r.Run(":" + port)
}
