package repositories

import "github.com/brandomota/golang-api/models"

func GetAllOrders() []models.Order {
	var orders []models.Order
	DBCon.First(&orders)
	return orders
}

func GetOrderById(id int64) models.Order {
	var order models.Order
	DBCon.Find(&order, id)
	return order
}

func GetOrdersByUserId(id int64) []models.Order {
	var orders []models.Order
	DBCon.Find(&orders).Where("userId = ?", id)
	return orders
}

func CreateOrder(order *models.Order) error {
	err := DBCon.Create(order).Error
	return err
}

func UpdateStatus(status string, id int64) error {
	err := DBCon.Update(&models.User{}).Where(id).Set("status", status).Error
	return err
}
