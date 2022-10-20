package main

import (
	"github.com/MaeYoiy/sa-65-WEEK05/controller"
	"github.com/MaeYoiy/sa-65-WEEK05/entity"

	"github.com/MaeYoiy/sa-65-WEEK05/middlewares"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// User Routes
			router.GET("/users", controller.ListUsers)
			router.GET("/user/:id", controller.GetUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			// Research_Room Routes
			router.GET("/researchrooms", controller.ListResearchRooms)
			router.GET("/researchroom/:id", controller.GetResearchRoom)
			router.POST("/researchrooms", controller.CreateResearchRoom)
			router.PATCH("/researchrooms", controller.UpdateResearchRoom)
			router.DELETE("/researchrooms/:id", controller.DeleteResearchRoom)

			// Equipment Routes
			router.GET("/equipments", controller.ListEquipments)
			router.GET("/equipment/:id", controller.GetEquipment)
			router.POST("/equipments", controller.CreateEquipment)
			router.PATCH("/equipments", controller.UpdateEquipment)
			router.DELETE("/equipments/:id", controller.DeleteEquipment)

			// Room_Type Routes
			router.GET("/roomtypes", controller.ListRoomTypes)
			router.GET("/roomtype/:id", controller.GetRoomType)
			router.POST("/roomtypes", controller.CreateRoomType)
			router.PATCH("/roomtypes", controller.UpdateRoomType)
			router.DELETE("/roomtypes/:id", controller.DeleteRoomType)

			// AddOn Routes
			router.GET("/addons", controller.ListAddOns)
			router.GET("/addon/:id", controller.GetAddOn)
			router.POST("/addons", controller.CreateAddOn)
			router.PATCH("/addons", controller.UpdateAddOn)
			router.DELETE("/addons/:id", controller.DeleteAddOn)

			// Timeroom Routes
			router.GET("/timerooms", controller.ListTimes)
			router.GET("/timeroom/:id", controller.GetTime)
			router.POST("/timerooms", controller.CreateTime)
			router.PATCH("/timerooms", controller.UpdateTime)
			router.DELETE("/timerooms/:id", controller.DeleteTime)

			// Research_Room_Reservation_Record Routes
			router.GET("/researchroomreservationrecords", controller.ListResearchRoomReservationRecords)
			router.GET("/researchroomreservationrecord/:id", controller.GetResearchRoomReservationRecord)
			router.POST("researchroomreservationrecords", controller.CreateResearchRoomReservationRecord)
			router.PATCH("/researchroomreservationrecords", controller.UpdateResearchRoomReservationRecord)
			router.DELETE("/researchroomreservationrecords/:id", controller.DeleteResearchRoomReservationRecord)

		}
	}

	// Signup User Route
	r.POST("/signup", controller.CreateUser)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("0.0.0.0:8080")
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
