package repositories

import "github.com/brandomota/golang-api/models"

func GetAllProducts() []models.Product {
	var products []models.Product
	DBCon.Find(&products)

	return products
}

func GetProductById(id int64) models.Product  {
	var product models.Product
	DBCon.First(&product,id)

	return product
}

func CreateProduct(newProduct *models.Product) error {
	var err error
	err = DBCon.Create(&newProduct).Error

	return err
}

func UpdateProduct(product *models.Product) error {
	var err error
	err = DBCon.Save(&product).Error

	return err
}

func DeleteProduct(product *models.Product) error {
	var err error
	err = DBCon.Unscoped().Delete(&product).Error

	return err
}
