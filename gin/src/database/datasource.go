package database

import (
	models "myapp/src/model"

	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(container *dig.Container) *gorm.DB {
	dsn := "host=localhost user=postgres password=349276185 dbname=go_gin port=5432 sslmode=disable TimeZone=Asia/Yangon"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Auto migrate the User model
    db.AutoMigrate(&models.User{})

	container.Provide(func () *gorm.DB{
		return db;
	});
    return db
}