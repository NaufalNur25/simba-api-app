package database

import (
	"fmt"
	"log"

	"github.com/naufal/simba-api-app/config"
	"github.com/naufal/simba-api-app/models"
)

func SeedPaymentMethods() {
	paymentMethods := []models.PaymentMethod{
		{
			Name:     "Cash",
			IsActive: true,
		},
		{
			Name:     "BCA",
			IsActive: true,
		},
		{
			Name:     "BRI",
			IsActive: true,
		},
		{
			Name:     "Mandiri",
			IsActive: true,
		},
		{
			Name:     "ShopeePay",
			IsActive: true,
		},
		{
			Name:     "Gopay",
			IsActive: true,
		},
		{
			Name:     "Jago",
			IsActive: true,
		},
		{
			Name:     "Ovo",
			IsActive: true,
		},
		{
			Name:     "Dana",
			IsActive: true,
		},
		{
			Name:     "Bank Transfer Lainnya",
			IsActive: true,
		},
	}

	for _, pm := range paymentMethods {
		var existing models.PaymentMethod
		err := config.DB.First(&existing, "name = ?", pm.Name).Error
		if err != nil {
			if err := config.DB.Create(&pm).Error; err != nil {
				log.Printf("❌ Gagal seed %s: %v", pm.Name, err)
			}
		}
	}

	fmt.Println("🌱 Seeder PaymentMethod: done")
}
