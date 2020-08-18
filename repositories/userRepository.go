package repositories

import models "github.com/brandomota/golang-api/models"

func GetAllUsers() []models.User {
	var users []models.User
	DBCon.Find(&users)
	return users
}

func GetUserById(id int64) models.User {
	var user models.User
	DBCon.First(&user, id)

	return user
}

func CreateUser(newUser *models.User) error {
	var err error
	err = DBCon.Create(&newUser).Error

	return err
}

func UpdateUser(user *models.User) error {
	var err error
	err = DBCon.Save(&user).Error

	return err
}

func DeleteUser(user *models.User) error {
	var err error
	err = DBCon.Unscoped().Delete(&user).Error

	return err
}
