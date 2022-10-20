package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
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
		&Employee{},
		&Role{},
		&Province{},
		&MemberClass{},
		&User{},
		&RoomType{},
		&Equipment{},
		&ResearchRoom{},
		&TimeRoom{},
		&AddOn{},
		&ResearchRoomReservationRecord{},
	)

	db = database

	//===========================================================================================================================
	//User
	password1, err := bcrypt.GenerateFromPassword([]byte("zaq1@wsX"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("zxvseta"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)

	//add example data
	//emp

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Sirinya",
		Email:    "sirinya@mail.com",
		Password: string(password1),
	})

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Attawit",
		Email:    "attawit@mail.com",
		Password: string(password2),
	})

	var sirin Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "sirinya@mail.com").Scan(&sirin)

	//Role
	student := Role{
		NAME:        "Student",
		BORROW_DAY:  3,
		BOOKROOM_HR: 3,
		BOOKCOM_HR:  4,
	}
	db.Model(&Role{}).Create(&student)

	teacher := Role{
		NAME:        "Teacher",
		BORROW_DAY:  7,
		BOOKROOM_HR: 12,
		BOOKCOM_HR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	//province
	korat := Province{
		NAME: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		NAME: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		NAME: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		NAME:     "classic",
		DISCOUNT: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		NAME:     "silver",
		DISCOUNT: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		NAME:     "gold",
		DISCOUNT: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		NAME:     "platinum",
		DISCOUNT: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	db.Model(&User{}).Create(&User{
		PIN:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		CIV:       "1111111111111",
		PHONE:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  string(password3),
		ADDRESS:   "ถนน a อำเภอ v",
		//FK
		Employee:    sirin,
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})
	//====================================================================================================================
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
	timeMorning := TimeRoom{
		Period: "08:00 - 12:00",
	}
	db.Model(&TimeRoom{}).Create(&timeMorning)

	timeAfternoon := TimeRoom{
		Period: "13:00 - 17:00",
	}
	db.Model(&TimeRoom{}).Create(&timeAfternoon)

	//RRRR 1
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room1,
		User:         preecha,
		AddOn:        powerPlugAndAdapter,
		BookDate:     time.Now(),
		TimeRoom:     timeMorning,
	})

	//RRRR 2
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room2,
		User:         kawin,
		AddOn:        powerPlugAndAdapterAndPillow,
		BookDate:     time.Now(),
		TimeRoom:     timeAfternoon,
	})

	// //
	// // ===Query
	// //

	// var target User
	// db.Model(&User{}).Find(&target, db.Where("pin = ?", "B6111111"))

}
