package database

import (
	"fmt"

	"github.com/naufal/simba-api-app/config"
	"github.com/naufal/simba-api-app/models"
)

func RunMigration() {
	err := config.DB.AutoMigrate(
		&models.PaymentMethod{},
	)

	if err != nil {
		panic("❌ Migration Failed: " + err.Error())
	}

	fmt.Println("✅ Database Migrated Successfully")

	runSeeder()
}

func runSeeder() {
	SeedPaymentMethods()
}
