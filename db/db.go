// db/db.go
package db

import (
	"binod210/task_remainder/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}

func Migrate() {
	DB.AutoMigrate(&models.Task{}, &models.Reminder{})
}
