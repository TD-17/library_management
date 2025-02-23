package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	dsn := "myuser:password@tcp(localhost:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	fmt.Println("Connected to MySQL successfully!")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	defer sqlDB.Close()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database is not reachable:", err)
	}
	fmt.Println("MySQL connection is alive!")
	return db
}
