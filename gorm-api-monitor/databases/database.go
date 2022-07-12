package databases

import (
	"robot-monitoreo/config"
	"robot-monitoreo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

var DATABASE_URI string = config.DatabaseEnv("DBUser") + ":" + config.DatabaseEnv("DBPassword") + "@tcp(" + config.DatabaseEnv("DBHost") + ":" + config.DatabaseEnv("DBPort") + ")/" + config.DatabaseEnv("DBName") + "?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error
	switch config.DatabaseEnv("DBConnection") {
	case "mysql":
		Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
	case "postgres":
		// Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
		// 	SkipDefaultTransaction: true,
		// 	PrepareStmt:            true,
		// })
	case "sqlite":
		// Database, err = gorm.Open(sqlite.Open(DATABASE_URI), &gorm.Config{
		// 	SkipDefaultTransaction: true,
		// 	PrepareStmt:            true,
		// })
	}

	if err != nil {
		panic(err)
	}

	Database.Logger = logger.Default.LogMode(logger.Info)

	Database.AutoMigrate(&models.Dog{})

	return nil
}
