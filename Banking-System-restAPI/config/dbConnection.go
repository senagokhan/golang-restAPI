package config

import (
	Models "Project5-API1/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}
	DB = db
	fmt.Println("Connected to database successfully")

	AutoMigrate(db)

}

func AutoMigrate(connection *gorm.DB) {
	err := connection.Debug().AutoMigrate(
		&Models.Accounts{},
		&Models.Investment{},
		&Models.Card{},
		&Models.Customer{},
		&Models.Representative{},
		&Models.Branches{},
		&Models.Transactions{},
	)
	if err != nil {
		return
	}
}
