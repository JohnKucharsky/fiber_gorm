package db

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/model"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func New() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Println("Can't load env")
	}

	dbAddress := "postgres://postgres:pass@db:5432/data?sslmode=disable"
	dbAddressEnv := os.Getenv("DB_URL")
	var dbAddressString = dbAddressEnv
	if dbAddressEnv == "" {
		dbAddressString = dbAddress
	}

	db, err := gorm.Open(
		postgres.Open(dbAddressString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		log.Fatal("Can't connect to db", err.Error())
	}

	return db
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}
