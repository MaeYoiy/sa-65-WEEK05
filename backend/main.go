package main

import (
	"github.com/MaeYoiy/sa-65-WEEK05/controller"
	"github.com/MaeYoiy/sa-65-WEEK05/entity"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()

	r.Use(CORSMiddleware())

	// User Routes
	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	// Research_Room Routes
	r.GET("/researchrooms", controller.ListResearchRooms)
	r.GET("/researchroom/:id", controller.GetResearchRoom)
	r.POST("/researchrooms", controller.CreateResearchRoom)
	r.PATCH("/researchrooms", controller.UpdateResearchRoom)
	r.DELETE("/researchrooms/:id", controller.DeleteResearchRoom)

	// Equipment Routes
	r.GET("/equipments", controller.ListEquipments)
	r.GET("/equipment/:id", controller.GetEquipment)
	r.POST("/equipments", controller.CreateEquipment)
	r.PATCH("/equipments", controller.UpdateEquipment)
	r.DELETE("/equipments/:id", controller.DeleteEquipment)

	// Room_Type Routes
	r.GET("/roomtypes", controller.ListRoomTypes)
	r.GET("/roomtype/:id", controller.GetRoomType)
	r.POST("/roomtypes", controller.CreateRoomType)
	r.PATCH("/roomtypes", controller.UpdateRoomType)
	r.DELETE("/roomtypes/:id", controller.DeleteRoomType)

	// AddOn Routes
	r.GET("/addons", controller.ListAddOns)
	r.GET("/addon/:id", controller.GetAddOn)
	r.POST("/addons", controller.CreateAddOn)
	r.PATCH("/addons", controller.UpdateAddOn)
	r.DELETE("/addons/:id", controller.DeleteAddOn)

	// Time Routes
	r.GET("/times", controller.ListTimes)
	r.GET("/time/:id", controller.GetTime)
	r.POST("/times", controller.CreateTime)
	r.PATCH("/times", controller.UpdateTime)
	r.DELETE("/times/:id", controller.DeleteTime)

	// Research_Room_Reservation_Record Routes
	r.GET("/researchroomreservationrecords", controller.ListResearchRoomReservationRecords)
	r.GET("/researchroomreservationrecord/:id", controller.GetResearchRoomReservationRecord)
	r.POST("researchroomreservationrecords", controller.CreateResearchRoomReservationRecord)
	r.PATCH("/researchroomreservationrecords", controller.UpdateResearchRoomReservationRecord)
	r.DELETE("/researchroomreservationrecords/:id", controller.DeleteResearchRoomReservationRecord)

	// Run the server
	r.Run()
}
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
