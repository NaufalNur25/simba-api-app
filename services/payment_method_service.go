package services

import (
	"github.com/naufal/simba-api-app/models"
	"github.com/naufal/simba-api-app/repository"
)

func GetAllPaymentMethod(isActive *bool) ([]models.PaymentMethod, error) {
	return repository.FetchPaymentMethods(isActive)
}

func GetPaymentMethodByID(id string) (models.PaymentMethod, error) {
	return repository.GetPaymentMethodByID(id)
}

func CreatePaymentMethod(input models.PaymentMethod) (models.PaymentMethod, error) {
	return repository.CreatePaymentMethod(input)
}

func UpdatePaymentMethod(id string, input models.PaymentMethod) (models.PaymentMethod, error) {
	return repository.UpdatePaymentMethod(id, input)
}

func DeletePaymentMethod(id string) error {
	return repository.DeletePaymentMethod(id)
}
