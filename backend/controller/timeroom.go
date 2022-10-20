package controller

import (
	"net/http"

	"github.com/MaeYoiy/sa-65-WEEK05/entity"
	"github.com/gin-gonic/gin"
)

// POST /timerooms
func CreateTime(c *gin.Context) {
	var timeroom entity.TimeRoom
	if err := c.ShouldBindJSON(&timeroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&timeroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": timeroom})
}

// GET /timeroom/:id
func GetTime(c *gin.Context) {
	var timeroom entity.TimeRoom
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM time_rooms WHERE id = ?", id).Scan(&timeroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": timeroom})
}

// GET /timerooms
func ListTimes(c *gin.Context) {
	var timerooms []entity.TimeRoom
	if err := entity.DB().Raw("SELECT * FROM time_rooms").Scan(&timerooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": timerooms})
}

// DELETE /timerooms/:id
func DeleteTime(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM time_rooms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /timerooms
func UpdateTime(c *gin.Context) {
	var timeroom entity.TimeRoom
	if err := c.ShouldBindJSON(&timeroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", timeroom.ID).First(&timeroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}
	if err := entity.DB().Save(&timeroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": timeroom})

}
