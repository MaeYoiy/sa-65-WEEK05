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
		//Book
		&Book{},
		&BookType{},
		&Shelf{},
		//User
		&Province{},
		&MemberClass{},
		&User{},
		&Role{},
		&Employee{},
		//Borrow
		&Borrow{},
		//reseachroom
		&RoomType{},
		//&Equipment{},
		&ResearchRoom{},
		&TimeRoom{},
		&AddOn{},
		&ResearchRoomReservationRecord{},
		//Bill
		&Bill{},
		//Com_reser
		&Computer_os{},
		&Computer_reservation{},
		&Computer{},
		&Time_com{},
		//Problem
		&Place_Class{},
		&Relation{},
		&Toilet{},
		&ReadingZone{},
		&ProblemReport{},
	)

	db = database

	password1, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("3333333333333"), 14)

	//add example data

	//Role
	student := Role{
		Name:       "Student",
		BorrowDay:  3,
		BookRoomHR: 3,
		BookComHR:  4,
	}

	db.Model(&Role{}).Create(&student)

	teacher := Role{
		Name:       "Teacher",
		BorrowDay:  7,
		BookRoomHR: 12,
		BookComHR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	employee := Role{
		Name:       "Employee",
		BorrowDay:  5,
		BookRoomHR: 6,
		BookComHR:  6,
	}
	db.Model(&Role{}).Create(&employee)

	//province
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		Name:     "classic",
		Discount: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		Name:     "silver",
		Discount: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		Name:     "gold",
		Discount: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		Name:     "platinum",
		Discount: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	db.Model(&User{}).Create(&User{
		Pin:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		Civ:       "1111111111111",
		Phone:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  string(password1),
		Address:   "ถนน a อำเภอ v",
		//FK
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "E123456",
		FirstName: "kawin",
		LastName:  "l.",
		Civ:       "1234567890123",
		Phone:     "0899999999",
		Email:     "kawin@mail.com",
		Password:  string(password2),
		Address:   "หอ b อำเภอ r",
		//FK
		Role:        employee,
		Province:    chon,
		MemberClass: silver,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "T8888",
		FirstName: "sirinya",
		LastName:  "kotpanya",
		Civ:       "3333333333333",
		Phone:     "0823456789",
		Email:     "sirinya@mail.com",
		Password:  string(password3),
		Address:   "บ้าน c อำเภอ q",
		//FK
		Role:        teacher,
		Province:    bangkok,
		MemberClass: plat,
	})

	//========================================Report============================================================================
	pcRdZone := Place_Class{
		Name: "Reading Zone",
	}
	db.Model(&Place_Class{}).Create(&pcRdZone)

	pcTlt := Place_Class{
		Name: "Toilet",
	}
	db.Model(&Place_Class{}).Create(&pcTlt)

	pcReschRoom := Place_Class{
		Name: "Research Room",
	}
	db.Model(&Place_Class{}).Create(&pcReschRoom)

	pcCom := Place_Class{
		Name: "Computer",
	}
	db.Model(&Place_Class{}).Create(&pcCom)

	//============================================ResearchRoomReservationRecord================================================================================
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

	//Research room
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "SR01",
		RoomType:    single_room,
		Place_Class: pcReschRoom,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "GR02",
		RoomType:    group_room,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "TR03",
		RoomType:    tutor_room,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "SR04",
		RoomType:    single_room,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "GR05",
		RoomType:    group_room,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "TR06",
		RoomType:    tutor_room,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "SR07",
		RoomType:    single_room,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "GR08",
		RoomType:    group_room,
		Place_Class: pcReschRoom,
	})
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		Name:        "TR09",
		RoomType:    tutor_room,
		Place_Class: pcReschRoom,
	})

	//ดึง Data ของ User มาเก็บไว้ในตัวแปรก่อน
	var preecha User
	db.Raw("SELECT * FROM users WHERE email = ?", "preechapat@mail.com").Scan(&preecha)

	var kawin User
	db.Raw("SELECT * FROM users WHERE email = ?", "kawin@mail.com").Scan(&kawin)

	//ดึง Data ของ researchroom มาเก็บไว้ในตัวแปรก่อน
	var Room1 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE id = ?", 1).Scan(&Room1)
	var Room2 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE id = ?", 2).Scan(&Room2)

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

}
