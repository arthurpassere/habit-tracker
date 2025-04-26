package database

import (
	"log"

	"github.com/arthurpassere/habit-tracker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("habit_tracker.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Habit{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
	