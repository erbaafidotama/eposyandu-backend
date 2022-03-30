package config

import (
	"eposyandu/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB() *gorm.DB {
	godotenv.Load()
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	fmt.Println(dbName)
	dsn := "user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	fmt.Printf("\n\n%#v\n\n", db)
	if err != nil {
		panic("Connecting database failed:" + err.Error())
	}
	db.AutoMigrate(&models.Account{})
	// db.AutoMigrate(&models.Warga{})
	// db.AutoMigrate(&models.Lookup{})
	// db.AutoMigrate(&models.LookupDetail{})
	db.AutoMigrate(&models.AllAnak{})
	return db
}
