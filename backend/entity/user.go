package entity

import (
	"time"

	"gorm.io/gorm"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
type EMPLOYEE struct {
	gorm.Model
	NAME     string
	PASSWORD string
	USERS    []User `gorm:"foreignKey:EMP_ID"`
}

type ROLE struct {
	gorm.Model
	NAME        string
	BORROW_DAY  int
	BOOKROOM_HR int
	BOOKCOM_HR  int
	USERS       []User `gorm:"foreignKey:ROLE_ID"`
}

type PROVINCE struct {
	gorm.Model
	NAME  string
	USERS []User `gorm:"foreignKey:PROVINCE_ID"`
}

type MemberClass struct {
	gorm.Model
	NAME     string
	DISCOUNT string
	USERS    []User `gorm:"foreignKey:MemberClass_ID"`
}

type User struct {
	gorm.Model
	PIN       string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	CIV       string `gorm:"uniqueIndex"`
	PHONE     string
	EMAIL     string `gorm:"uniqueIndex"`
	PASSWORD  string
	ADDRESS   string
	//FK
	EMP_ID         *uint
	ROLE_ID        *uint
	PROVINCE_ID    *uint
	MemberClass_ID *uint
	//JOIN
	PROVINCE    PROVINCE
	ROLE        ROLE
	MemberClass MemberClass
	EMP         EMPLOYEE
	RRRR        []ResearchRoomReservationRecord `gorm:"foreignKey:UserID"`
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//

type RoomType struct {
	gorm.Model
	Type string
	// ResearchRoom []ResearchRoom `gorm:"foreignKey:RoomTypeID"`
}

type Equipment struct {
	gorm.Model
	Name string
	// ResearchRoom []ResearchRoom `gorm:"foreignKey:EquipmentID"`
}

type ResearchRoom struct {
	gorm.Model
	Name string

	RoomTypeID *uint    //FK
	RoomType   RoomType `gorm:"references:id"` //JOIN //ทำการตึง id ของ RoomType

	EquipmentID *uint     //FK
	Equipment   Equipment `gorm:"references:id"` //JOIN

	// Place_ProblemID *uint         //FK
	// Place_Problem   Place_Problem //JOIN
	// RRRR []ResearchRoomReservationRecord `gorm:"foreignKey:ResearchRoomID"`
}

type Time struct {
	gorm.Model
	Period string
	// RRRR   []ResearchRoomReservationRecord `gorm:"foreignKey:TimeID"`
}

type AddOn struct {
	gorm.Model
	Name string
	// RRRR []ResearchRoomReservationRecord `gorm:"foreignKey:AddOnID"`
}

type ResearchRoomReservationRecord struct {
	gorm.Model
	BookDate time.Time

	UserID *uint //FK
	User   User  //JOIN

	ResearchRoomID *uint        //FK
	ResearchRoom   ResearchRoom `gorm:"references:id"`

	AddOnID *uint //FK
	AddOn   AddOn `gorm:"references:id"`

	TimeID *uint //FK
	Time   Time  `gorm:"references:id"`
}
