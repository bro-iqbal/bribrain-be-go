package database

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	sqlDB, err := sql.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("[*] Open DB connection failed", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Println("[*] Open DB connection failed", err)
	}

	log.Println("[*] Open DB connection successful")

	return db
}
