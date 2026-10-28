package configs

import (
	"airun/brain-tracker/server/src/models"
	"airun/brain-tracker/server/src/utils"
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDatabaseDSN() string {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		utils.GetEnv("DB_HOST", ""),
		utils.GetEnv("DB_USER", ""),
		utils.GetEnv("DB_PASS", ""),
		utils.GetEnv("DB_NAME", ""),
		utils.GetEnv("DB_PORT", ""),
		utils.GetEnv("DB_SSLMODE", ""),
		utils.GetEnv("DB_TIMEZONE", ""),
	)
	return dsn
}

func InitiateDatabaseConnection() (err error) {
	dsn := GetDatabaseDSN()
	log.Print("test ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func Migrate() error {
	if DB == nil {
		return errors.New("No Database Connection is Initiated.")
	}
	err := DB.AutoMigrate(
		&models.SignalData{},
	)
	if err != nil {
		return err
	}
	return nil
}
