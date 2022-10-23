package controller

import (
	"net/http"

	"github.com/MaeYoiy/sa-65-WEEK05/entity"
	"github.com/gin-gonic/gin"
)

// POST /researchrooms
func CreateResearchRoom(c *gin.Context) {
	var researchroom entity.ResearchRoom
	var roomtype entity.RoomType
	var equipment entity.Equipment

	//รับค่ามาจาก body ก่อน
	if err := c.ShouldBindJSON(&researchroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา roomtype ด้วย id
	if tx := entity.DB().Where("id = ?", researchroom.RoomTypeID).First(&roomtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room type not found"})
		return
	}

	// 10: ค้นหา equipmrnt ด้วย id
	if tx := entity.DB().Where("id = ?", researchroom.EquipmentID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	// 12: สร้าง ResearchRoom
	RR := entity.ResearchRoom{
		RoomType:  roomtype,  // โยงความสัมพันธ์กับ Entity RoomType
		Equipment: equipment, // โยงความสัมพันธ์กับ Entity Equipment
	}

	// 13: บันทึก
	if err := entity.DB().Create(&RR).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RR})

}

// GET /researchroom/:id
func GetResearchRoom(c *gin.Context) {
	var researchroom entity.ResearchRoom
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&researchroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "researchroom not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchroom})
}

// GET /researchrooms
func ListResearchRooms(c *gin.Context) {
	var researchrooms []entity.ResearchRoom
	if err := entity.DB().Preload("RoomType").Preload("Equipment").Raw("SELECT * FROM research_rooms").Find(&researchrooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchrooms})
}

// DELETE /researchrooms/:id
func DeleteResearchRoom(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM research_rooms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ResearchRoom not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /researchrooms
func UpdateResearchRoom(c *gin.Context) {
	var researchroom entity.ResearchRoom
	if err := c.ShouldBindJSON(&researchroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", researchroom.ID).First(&researchroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ResearchRoom not found"})
		return
	}
	if err := entity.DB().Save(&researchroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchroom})
}
