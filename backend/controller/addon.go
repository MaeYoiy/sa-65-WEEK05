package controller

import (
	"net/http"

	"github.com/MaeYoiy/sa-65-WEEK05/entity"
	"github.com/gin-gonic/gin"
)

// POST /addons
func CreateAddOn(c *gin.Context) {
	var addon entity.AddOn

	if err := c.ShouldBindJSON(&addon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&addon).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": addon})
}

// GET /addon/:id
func GetAddOn(c *gin.Context) {
	var addon entity.AddOn
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM add_ons WHERE id = ?", id).Scan(&addon).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": addon})
}

// GET /addons
func ListAddOns(c *gin.Context) {
	var addons []entity.AddOn
	if err := entity.DB().Raw("SELECT * FROM add_ons").Scan(&addons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": addons})
}

// DELETE /addons/:id
func DeleteAddOn(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM add_ons WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "add_on not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /addons
func UpdateAddOn(c *gin.Context) {
	var addon entity.AddOn
	if err := c.ShouldBindJSON(&addon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", addon.ID).First(&addon); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "add_on not found"})
		return
	}
	if err := entity.DB().Save(&addon).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": addon})
}
