package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    dsn := "root:@tcp(127.0.0.1:3306)/tutor_booking_system?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB = db // Assign the connection to the global variable
    return db
}
