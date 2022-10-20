package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65-WEEK05.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&EMPLOYEE{},
		&ROLE{},
		&PROVINCE{},
		&MemberClass{},
		&User{},
		&RoomType{},
		&Equipment{},
		&ResearchRoom{},
		&Time{},

		&AddOn{},
		&ResearchRoomReservationRecord{},
	)

	db = database

	//User
	db.Model(&User{}).Create(&User{
		PIN:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		CIV:       "1111111111111",
		PHONE:     "0811111111",
		EMAIL:     "preechapat@mail.com",
		PASSWORD:  "1111111111111",
		ADDRESS:   "ถนน a อำเภอ v",
		//ไม่มี Province, Role, member, EMP_ID เพราะยังไม่ใส่ข้อมูลตารางเหล่านี้
	})

	db.Model(&User{}).Create(&User{
		PIN:       "B6222222",
		FirstName: "kawin",
		LastName:  "l.pat",
		CIV:       "2222222222222",
		PHONE:     "0922222222",
		EMAIL:     "kawin@mail.com",
		PASSWORD:  "2222222222222",
		ADDRESS:   "หอ b อำเภอ r",
		//ไม่มี Province, Role, member, EMP_ID เพราะยังไม่ใส่ข้อมูลตารางเหล่านี้
	})

	//Equipment data
	monitor := Equipment{
		Name: "จอ monitor สำหรับการนำเสนอ",
	}
	db.Model(&Equipment{}).Create(&monitor)

	printer := Equipment{
		Name: "เครื่องปริ้นท์",
	}
	db.Model(&Equipment{}).Create(&printer)

	printerMoniter := Equipment{
		Name: "เครื่องปริ้นท์ + จอ monitor สำหรับการนำเสนอ",
	}
	db.Model(&Equipment{}).Create(&printerMoniter)

	//Room_type data
	single_room := RoomType{
		Type: "ห้องเดี่ยว",
	}
	db.Model(&RoomType{}).Create(&single_room)

	group_room := RoomType{
		Type: "ห้องกลุ่ม",
	}
	db.Model(&RoomType{}).Create(&group_room)

	tutor_room := RoomType{
		Type: "ห้องสำหรับติว",
	}
	db.Model(&RoomType{}).Create(&tutor_room)

	//Research_Room Data
	// RR001 := Research_Room{
	// 	Name:      "RR001",
	// 	RoomType:  group_room,
	// 	Equipment: monitor,
	// }
	// db.Model(&Research_Room{}).Create(&RR001)

	// RR002 := Research_Room{
	// 	Name:      "RR002",
	// 	RoomType:  group_room,
	// 	Equipment: printerMoniter,
	// }
	// db.Model(&Research_Room{}).Create(&RR002)

	// RR003 := Research_Room{
	// 	Name:      "RR003",
	// 	RoomType:  tutor_room,
	// 	Equipment: printer,
	// }
	// db.Model(&Research_Room{}).Create(&RR003)

	// RR004 := Research_Room{
	// 	Name:      "RR004",
	// 	RoomType:  single_room,
	// 	Equipment: monitor,
	// }
	// db.Model(&Research_Room{}).Create(&RR004)

	//Research room
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:      "RR001",
		RoomType:  group_room,
		Equipment: monitor,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:      "RR002",
		RoomType:  group_room,
		Equipment: printerMoniter,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:      "RR003",
		RoomType:  tutor_room,
		Equipment: printer,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:      "RR004",
		RoomType:  single_room,
		Equipment: monitor,
	})

	//ดึง Data ของ User มาเก็บไว้ในตัวแปรก่อน
	var preecha User
	db.Raw("SELECT * FROM users WHERE email = ?", "preechapat@mail.com").Scan(&preecha)

	var kawin User
	db.Raw("SELECT * FROM users WHERE email = ?", "kawin@mail.com").Scan(&kawin)

	//ดึง Data ของ researchroom มาเก็บไว้ในตัวแปรก่อน
	var Room1 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE name = ?", "RR001").Scan(&Room1)
	var Room2 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE name = ?", "RR002").Scan(&Room2)
	var Room3 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE name = ?", "RR003").Scan(&Room3)
	var Room4 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE name = ?", "RR004").Scan(&Room4)

	//Addon data
	powerPlug := AddOn{
		Name: "ปลั๊กพ่วง",
	}
	db.Model(&AddOn{}).Create(&powerPlug)

	Adapter := AddOn{
		Name: "สายชาร์จ",
	}
	db.Model(&AddOn{}).Create(&Adapter)

	Pillow := AddOn{
		Name: "หมอน",
	}
	db.Model(&AddOn{}).Create(&Pillow)

	powerPlugAndAdapter := AddOn{
		Name: "ปลั๊กพ่วง + สายชาร์จ",
	}
	db.Model(&AddOn{}).Create(&powerPlugAndAdapter)

	adapterAndPillow := AddOn{
		Name: "สายชาร์จ + หมอน",
	}
	db.Model(&AddOn{}).Create(&adapterAndPillow)

	powerPlugAndAdapterAndPillow := AddOn{
		Name: "ปลั๊กพ่วง + สายชาร์จ + หมอน",
	}
	db.Model(&AddOn{}).Create(&powerPlugAndAdapterAndPillow)

	//Time data
	timeMorning := Time{
		Period: "08:00 - 12:00",
	}
	db.Model(&Time{}).Create(&timeMorning)

	timeAfternoon := Time{
		Period: "13:00 - 17:00",
	}
	db.Model(&Time{}).Create(&timeAfternoon)

	//RRRR 1
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room1,
		User:         preecha,
		AddOn:        powerPlugAndAdapter,
		BookDate:     time.Now(),
		Time:         timeMorning,
	})

	//RRRR 2
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room2,
		User:         kawin,
		AddOn:        powerPlugAndAdapterAndPillow,
		BookDate:     time.Now(),
		Time:         timeAfternoon,
	})

	// //
	// // ===Query
	// //

	// var target User
	// db.Model(&User{}).Find(&target, db.Where("pin = ?", "B6111111"))

}
