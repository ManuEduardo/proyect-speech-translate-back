package handlers

import (
	"net/http"

	"github.com/ManuEduardo/proyect-speech-translate-back/models"
	"github.com/ManuEduardo/proyect-speech-translate-back/services"
	"github.com/gin-gonic/gin"
)

func GetUsersHandler(ctx *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, users)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func PostUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := services.AddUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "No created"})
	}
	c.JSON(http.StatusCreated, user)
}
