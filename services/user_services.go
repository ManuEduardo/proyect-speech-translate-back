package services

import "github.com/ManuEduardo/proyect-speech-translate-back/models"

var users = []models.User{
	{
		ID:       1,
		Name:     "Juan",
		LastName: "Perez",
		Email:    "juan.perez@example.com",
	},
	{
		ID:       2,
		Name:     "Maria",
		LastName: "Gonzalez",
		Email:    "maria.gonzalez@example.com",
	},
	{
		ID:       3,
		Name:     "Carlos",
		LastName: "Rodriguez",
		Email:    "carlos.rodriguez@example.com",
	},
}

func GetUsers() ([]models.User, error) {
	return users, nil
}

func AddUser(user models.User) (models.User, error) {
	users = append(users, user)
	return user, nil
}
