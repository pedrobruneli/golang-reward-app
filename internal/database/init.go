package database

import (
	"os"
	"rewards/internal/usecases/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&users.User{})
	return db, nil
}
