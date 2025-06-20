package handlers

import (
	"github.com/arthurpassere/habit-tracker/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, c *gin.Context) {
	var user *models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if result := db.Create(&user); result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(201, user)
}

func CreateHabit(db *gorm.DB, c *gin.Context) {
	var habit *models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if result := db.Create(&habit); result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(201, habit)
}
