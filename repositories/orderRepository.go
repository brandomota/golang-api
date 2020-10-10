package repositories

import "github.com/brandomota/golang-api/models"

func GetAllOrders()[]models.Order {
	var orders []models.Order
	DBCon.First(&orders)
	return orders
}

func GetOrderById(id int64) models.Order {
	var order models.Order
	DBCon.Find(&order, id)
	return order
}

func CreateOrder(order *models.Order) error {
	var err error
	err = DBCon.Create(order).Error
	return err
}

func UpdateStatus(status string, id int64) error {
	var err error
	err = DBCon.Update(&models.User{}).Where(id).Set("status", status).Error
	return err
}
