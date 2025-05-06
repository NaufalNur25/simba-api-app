package repository

import (
	"github.com/naufal/simba-api-app/config"
	"github.com/naufal/simba-api-app/models"
)

func FetchPaymentMethods(isActive *bool) ([]models.PaymentMethod, error) {
	var paymentMethods []models.PaymentMethod
	query := config.DB

	if isActive != nil {
		query = query.Where("is_active = ?", *isActive)
	}

	err := query.Find(&paymentMethods).Error
	return paymentMethods, err
}

func GetPaymentMethodByID(id string) (models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	err := config.DB.First(&paymentMethod, "id = ?", id).Error

	return paymentMethod, err
}

func CreatePaymentMethod(paymentMethod models.PaymentMethod) (models.PaymentMethod, error) {
	err := config.DB.Create(&paymentMethod).Error
	return paymentMethod, err
}

func UpdatePaymentMethod(id string, updatedData models.PaymentMethod) (models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod
	if err := config.DB.First(&paymentMethod, "id = ?", id).Error; err != nil {
		return paymentMethod, err
	}

	paymentMethod.Name = updatedData.Name
	paymentMethod.IsActive = updatedData.IsActive

	err := config.DB.Save(&paymentMethod).Error
	return paymentMethod, err
}

func DeletePaymentMethod(id string) (bool error) {
	return config.DB.Delete(&models.PaymentMethod{}, "id = ?", id).Error
}
