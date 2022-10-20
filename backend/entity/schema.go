package entity

import (
	"time"

	"gorm.io/gorm"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
type Employee struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string `json:"-"`
	USERS    []User `gorm:"foreignKey:EmployeeID"`
}

type Role struct {
	gorm.Model
	NAME        string
	BORROW_DAY  int
	BOOKROOM_HR int
	BOOKCOM_HR  int
	USERS       []User `gorm:"foreignKey:RoleID"`
}

type Province struct {
	gorm.Model
	NAME  string
	USERS []User `gorm:"foreignKey:ProvinceID"`
}

type MemberClass struct {
	gorm.Model
	NAME     string
	DISCOUNT int
	USERS    []User `gorm:"foreignKey:MemberClassID"`
}

type User struct {
	gorm.Model
	PIN       string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	CIV       string `gorm:"uniqueIndex"`
	PHONE     string
	Email     string `gorm:"uniqueIndex"`
	Password  string `json:"-"`
	ADDRESS   string
	//FK
	EmployeeID    *uint
	RoleID        *uint
	ProvinceID    *uint
	MemberClassID *uint
	//JOIN
	Province    Province    `gorm:"references:id"`
	Role        Role        `gorm:"references:id"`
	MemberClass MemberClass `gorm:"references:id"`
	Employee    Employee    `gorm:"references:id"`
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

type TimeRoom struct {
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

	TimeRoomID *uint    //FK
	TimeRoom   TimeRoom `gorm:"references:id"`
}
