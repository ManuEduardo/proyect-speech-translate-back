package main

import (
	"github.com/ManuEduardo/proyect-speech-translate-back/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define routes
	r.GET("/users", handlers.GetUsersHandler)
	r.POST("/user", handlers.PostUserHandler)

	r.Run(":8080") // Run the server on port 8080
}
