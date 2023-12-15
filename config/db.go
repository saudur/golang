package config

import (
	"sekolah/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=postgres password=saudur123 dbname=sekolah port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db
	// var err error
	// // host := os.Getenv("DB_HOST")
	// // username := os.Getenv("DB_USER")
	// // password := os.Getenv("DB_PASS")
	// // databaseName := os.Getenv("DB_NAME")
	// // port := os.Getenv("DB_PORT")
	// host := "localhost"
	// username := "postgres"
	// password := "saudur123"
	// databaseName := "sekolah"
	// port := "saudur123"

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, databaseName, port)
	// Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.User{})

}
