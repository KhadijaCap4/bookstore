package database

import (
	"example/bookstore/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "root:@tcp(localhost:3306)/bookstore"  
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Connexion to the database failed")
    }
    DB = db

    db.AutoMigrate(&models.Book{})
}
