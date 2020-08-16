package migrations

import (
	"fmt"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/models"
)

func InitMigrate() {
	fmt.Println("1")
	config.LoadEnv()
	_, err := config.SetupDB()
	if err != nil {
		panic(err)
	}
	Migrate()
}

func Migrate() {
	fmt.Println("Migrating...")
	Log.Info().Msg("Migrating")
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.PaymentMethod{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.Transaction{})
	DB.AutoMigrate(&models.UserTransactionLog{})
	DB.AutoMigrate(&models.File{})
	DB.AutoMigrate(&models.UserFile{})
	Log.Info().Msg("Migrated")
	fmt.Println("Migrated...")
}
