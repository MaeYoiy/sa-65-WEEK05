package controller

import (
	"net/http"
	"time"

	"github.com/MaeYoiy/sa-65-WEEK05/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type researchroomlast struct {
	//ResearchRoom
	RoomTypeID  *uint
	EquipmentID *uint
	//researchroomreservaionrecord
	AddOnID    *uint
	TimeRoomID *uint
	UserID     *uint
	BookDate   time.Time
}

// POST /researchroomreservationrecords
func CreateResearchRoomReservationRecord(c *gin.Context) {

	//var researchroomreservationrecord entity.ResearchRoomReservationRecord
	var researchroom entity.ResearchRoom
	var user entity.User
	var addon entity.AddOn
	var timeroom entity.TimeRoom
	var roomtype entity.RoomType
	var equipment entity.Equipment
	var researchroomlast researchroomlast

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร researchRoomReservationRecord
	//รับค่ามาจาก body ก่อน
	if err := c.ShouldBindJSON(&researchroomlast); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา roomtype ด้วย id
	if tx := entity.DB().Where("id = ?", researchroomlast.EquipmentID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room type not found"})
		return
	}

	// 10: ค้นหา equipment ด้วย id
	if tx := entity.DB().Where("id = ?", researchroomlast.RoomTypeID).First(&roomtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	// 11: ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", researchroomlast.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 12: ค้นหา AddOn ด้วย id
	if tx := entity.DB().Where("id = ?", researchroomlast.AddOnID).First(&addon); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "add-on not found"})
		return
	}

	// 13: ค้นหา Time ด้วย id
	if tx := entity.DB().Where("id = ?", researchroomlast.TimeRoomID).First(&timeroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	//สร้างตาราง ResearchRoom
	researchroom = entity.ResearchRoom{
		RoomType:  roomtype,
		Equipment: equipment,
	}

	//ทำการตรวจสอบความถูกต้องของ struct ก่อนนำไปสร้าง record
	if _, err := govalidator.ValidateStruct(researchroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//บันทึกค่าลงในตาราง ResearchRoom
	if err := entity.DB().Create(&researchroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 14: สร้าง researchRoomReservationRecord
	RRRR := entity.ResearchRoomReservationRecord{
		ResearchRoom: researchroom,              // โยงความสัมพันธ์กับ Entity ResearchRoom
		User:         user,                      // โยงความสัมพันธ์กับ Entity User
		AddOn:        addon,                     // โยงความสัมพันธ์กับ Entity AddOn
		TimeRoom:     timeroom,                  // โยงความสัมพันธ์กับ Entity TimeRoom
		BookDate:     researchroomlast.BookDate, // ตั้งค่าฟิลด์ researchroomreservationrecord.BookDat
	}

	//ทำการตรวจสอบความถูกต้องของ struct ก่อนนำไปสร้าง record
	if _, err := govalidator.ValidateStruct(RRRR); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 15: บันทึกตารางหลัก(ResearchRoomReservationRecord)
	if err := entity.DB().Create(&RRRR).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RRRR})
}

// GET /researchroomreservationrecord/:id
func GetResearchRoomReservationRecord(c *gin.Context) {
	var researchroomreservationrecord entity.ResearchRoomReservationRecord
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&researchroomreservationrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "researchroomreservationrecord not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchroomreservationrecord})
}

// GET /researchroomreservationrecords
func ListResearchRoomReservationRecords(c *gin.Context) {
	var researchroomreservationrecords []entity.ResearchRoomReservationRecord
	if err := entity.DB().Preload("ResearchRoom").Preload("ResearchRoom.RoomType").Preload("ResearchRoom.Equipment").Preload("User").Preload("AddOn").Preload("TimeRoom").Raw("SELECT * FROM research_room_reservation_records").Find(&researchroomreservationrecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchroomreservationrecords})
}

// DELETE /researchroomreservationrecords/:id
func DeleteResearchRoomReservationRecord(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM research_room_reservation_records WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "researchroomreservationrecord not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /researchroomreservationrecords
func UpdateResearchRoomReservationRecord(c *gin.Context) {
	var researchroomreservationrecord entity.ResearchRoomReservationRecord
	if err := c.ShouldBindJSON(&researchroomreservationrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", researchroomreservationrecord.ID).First(&researchroomreservationrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "researchroomreservationrecord not found"})
		return
	}
	if err := entity.DB().Save(&researchroomreservationrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": researchroomreservationrecord})
}
