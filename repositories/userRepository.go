package repositories

import models "github.com/brandomota/golang-api/models"

func GetAllUsers() {
	var users = DBCon.Find(&models.User)
	return users
}
