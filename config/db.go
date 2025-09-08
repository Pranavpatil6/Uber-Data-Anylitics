package config

import (
	"github.com/HOA-Share/hoashare-go-common/config"
	"github.com/HOA-Share/hoashare-go-common/db"
	"github.com/HOA-Share/hoashare-go-common/logger"
	"github.com/pranavpatil6/models"
	"gorm.io/gorm"
	dblogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDb() {
	dbConn, err := db.ConnectDB(config.DatabaseConfig{
		Host:     settings.DBConfig.Host,
		Port:     settings.DBConfig.Port,
		DBName:   settings.DBConfig.DBName,
		User:     settings.DBConfig.User,
		Password: settings.DBConfig.Password,
		SSLMode:  settings.DBConfig.SSLMode,
	}, dblogger.Info)
	if err != nil {
		panic(err)
	}
	DB = dbConn.Raw()

	if settings.DBConfig.Migrate {
		DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
		logger.GetLogger().Info("migrating database")

		err := DB.AutoMigrate(
			&models.Ride{},
		)
		if err != nil {
			panic(err)
		}
		logger.GetLogger().Info("migration complete")

	}
}
