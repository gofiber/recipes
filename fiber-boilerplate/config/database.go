package config

import (
	"fmt"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/jinzhu/gorm"
	"time"
)

type Database struct {
	DB_Driver string //nolint:goimports,gofmt
	DB_Host   string
	DB_Port   int
	DB_User   string
	DB_Pass   string
	DB_Name   string
}

var DBConfig *Database //nolint:gochecknoglobals

func LoadDBConfig() {
	loadDBDefaultConfig()
	ViperConfig.Unmarshal(&DBConfig) //nolint:errcheck
}

func loadDBDefaultConfig() {
	ViperConfig.SetDefault("DB_DRIVER", "mysql")
	ViperConfig.SetDefault("DB_HOST", "localhost")
	ViperConfig.SetDefault("DB_PORT", 3306)
	ViperConfig.SetDefault("DB_USER", "root")
	ViperConfig.SetDefault("DB_PASS", "root")
	ViperConfig.SetDefault("DB_NAME", "casbin")
}

func SetupDB() (*gorm.DB, error) {
	LoadDBConfig()
	connectionString := ""
	if DB != nil {
		return DB, nil
	}
	switch DBConfig.DB_Driver {
	case "postgres":
		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", DBConfig.DB_Host, DBConfig.DB_Port, DBConfig.DB_User, DBConfig.DB_Name, DBConfig.DB_Pass)
	default:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DBConfig.DB_User, DBConfig.DB_Pass, DBConfig.DB_Host, DBConfig.DB_Port, DBConfig.DB_Name)
	}
	//nolint:wsl,lll
	var err error //nolint:wsl

	// Connect again with DB name.
	DB, err = gorm.Open(DBConfig.DB_Driver, connectionString)
	if err != nil {
		panic(err)
	}

	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetMaxIdleConns(100)
	DB.DB().SetConnMaxLifetime(5 * time.Minute)
	return DB, nil //nolint:wsl
}
