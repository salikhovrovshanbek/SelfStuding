package DB

import (
	"api_gin_gorm/structures"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(&structures.Book{})
	if err != nil {
		log.Fatal(err)
		return
	}

	DB = database
}
