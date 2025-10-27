package config

import (
	"log"
	"uber-data-analytics/models"

	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)


var DB *gorm.DB

var Cfg = GetConfig()

func InitDB(){
	log.Println("DB DSN:", os.Getenv("DB_HOST")) // Add this line
    dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", Cfg.DbHost, Cfg.DbUser, Cfg.DbPassword, Cfg.DbName, Cfg.DbPort)
    db, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err!= nil{
        log.Fatal("Failed to connect DB", err)
    }
    DB=db
    DB.AutoMigrate(&models.Ride{})
}
