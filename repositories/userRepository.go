package repositories

import models "github.com/brandomota/golang-api/models"

func GetAllUsers() []models.User {
	var users []models.User
	DBCon.Find(&users)
	return users
}
