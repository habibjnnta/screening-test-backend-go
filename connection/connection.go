package connection

import (
	"log"
	migrtations "screening-test/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/screening-test-backend-habib?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&migrtations.Player{}, &migrtations.Team{})

	return db
}
