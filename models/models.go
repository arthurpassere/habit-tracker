package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Habits []Habit `gorm:"foreignKey:UserID"`
}

type Habit struct {
	gorm.Model
	Title  string      `json:"title"`
	UserID uint        `json:"user_id"`
	Logs   []HabitLogs `gorm:"foreignKey:HabitID"`
}

type HabitLogs struct {
	gorm.Model
	HabitID uint   `json:"habit_id"`
	Date    string `json:"date"`
	Done    bool   `json:"done"`
}
