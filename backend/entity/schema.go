package entity

import (
	"time"

	"gorm.io/gorm"
)

// ====================User================================
type Employee struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Book     []Book   `gorm:"foreignKey:Employee_ID"`
	Borrow   []Borrow `gorm:"foreingnKey:Employee_ID"`
	Bills    []Bill   `gorm:"foreignKey:EmployeeID"`
}

type Role struct {
	gorm.Model
	Name       string
	BorrowDay  int
	BookRoomHR int
	BookComHR  int
	Users      []User `gorm:"foreignKey:RoleID"`
	Book       []Book `gorm:"foreignKey:Employee_ID"`
}

type Province struct {
	gorm.Model
	Name  string
	Users []User `gorm:"foreignKey:ProvinceID"`
}

type MemberClass struct {
	gorm.Model
	Name     string
	Discount int
	Users    []User `gorm:"foreignKey:MemberClassID"`
}

type User struct {
	gorm.Model
	Pin       string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Civ       string `gorm:"uniqueIndex"`
	Phone     string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Address   string
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

	Borrow []Borrow `gorm:"foreingnKey:User_ID"`
	Bills  []Bill   `gorm:"foreignKey:UserID"`
}

// ====================Book================================
type Book struct {
	gorm.Model
	Name string
	//ทำหน้าที่เป็น FK
	Employee_ID *uint
	Booktype_ID *uint
	Shelf_ID    *uint
	ROLE_ID     *uint
	//join ให้งายขึ้น
	Employee Employee `gorm:"references:id"`
	Booktype BookType `gorm:"references:id"`
	Shelf    Shelf    `gorm:"references:id"`
	Role     Role     `gorm:"references:id"`
	Author   string
	Page     int
	Quantity int
	Price    int
	Date     time.Time

	Borrow []Borrow `gorm:"foreingnKey:Book_ID"`
}

type BookType struct {
	gorm.Model
	Type string
	//1 book type มีได้หลาย book
	Book []Book `gorm:"foreignKey:Booktype_ID"`
}

type Shelf struct {
	gorm.Model
	Type  string
	Floor uint
	//1 shelf มีได้หลาย book
	Book []Book `gorm:"foreignKey:Shelf_ID"`
}

// ====================Borrow================================
type Borrow struct {
	gorm.Model

	//ทำหน้าที่เป็น FK
	Employee_ID *uint
	Employee    Employee `gorm:"references:id"`

	Book_ID *uint
	Book    Book `gorm:"references:id"`

	User_ID *uint
	User    User `gorm:"references:id"`

	User_Name string
	Book_Name string
	Status    string
	DateTime  time.Time
	DueDate   time.Time
	Fine      int
}

// ====================ResearchRoom================================
type RoomType struct {
	gorm.Model
	Type string
	// ResearchRoom []ResearchRoom `gorm:"foreignKey:RoomTypeID"`
}

type ResearchRoom struct {
	gorm.Model
	Name string

	RoomTypeID *uint    //FK
	RoomType   RoomType `gorm:"references:id"` //JOIN //ทำการตึง id ของ RoomType

	Place_Class_ID *uint //researchroom ดึง placeclass
	Place_Class    Place_Class
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"` //problemreport ดึง ResearchRoom
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
	User   User  `gorm:"references:id"` //JOIN

	ResearchRoomID *uint        //FK
	ResearchRoom   ResearchRoom `gorm:"references:id"`

	AddOnID *uint //FK
	AddOn   AddOn `gorm:"references:id"`

	TimeRoomID *uint    //FK
	TimeRoom   TimeRoom `gorm:"references:id"`
}

// ====================Bill================================
type Bill struct { //เป็นการ get api มาจาก code จะไปอยู่ในส่วนของ front end
	gorm.Model
	//ทำหน้าที่เป็น FK
	BookID *uint
	Book   Book `gorm:"references:id"`

	//ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	//ทำหน้าที่เป็น FK
	//MemberClass_ID *uint

	//ทำหน้าที่เป็น FK
	UserID *uint
	User   User `gorm:"references:id"`

	//join ให้งายขึ้น

	Book_Name        string
	MemberClass_Name string
	Book_Price       uint //uint ไม่มีเครื่องหมายติดลบ
	Discount         uint
	Total            uint
	BillTime         time.Time
}

// ====================ComputerReservation================================
type Computer_reservation struct {
	gorm.Model
	Date time.Time

	Computer_id *uint    //FK
	Computer    Computer `gorm:"references:id"` //JOIN

	// COMPUTER_OS_ID *uint       //FK
	// COMPUTER_OS    COMPUTER_OS `gorm:"references:id"` //JOIN

	Time_com_id *uint    //FK
	Time_com    Time_com `gorm:"references:id"` //JOIN

	UserID *uint //FK
	User   User  `gorm:"references:id"` //JOIN
}

type Computer struct {
	gorm.Model
	Comp_name string
	Comp_room string

	Computer_os_id *uint       //FK
	Computer_os    Computer_os `gorm:"references:id"` //JOIN
	Place_Class_ID *uint
	Place_Class    Place_Class
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`

	// COMPUTER_RESERVATION []COMPUTER_RESERVATION `gorm:"foreignKey:COMPUTER_ID"`
}

type Computer_os struct {
	gorm.Model
	Name string
	// COMPUTER []COMPUTER `gorm:"foreignKey:COMPUTER_OS_ID"`
}

type Time_com struct {
	gorm.Model
	Time_com_period      string
	Computer_reservation []Computer_reservation `gorm:"foreignKey:Time_com_id"`
}

// ====================Problem================================
type Problem struct {
	gorm.Model
	Name string
	// 1 Problem อยู่ได้ในหลาย Relation
	Relations []Relation `gorm:"foreignKey:Problem_ID"`
	// 1 Problem อยู่ได้ในหลาย Report
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`
}

type Place_Class struct {
	gorm.Model
	Name string

	Relations     []Relation     `gorm:"foreignKey:Place_Class_ID"`
	Toilets       []Toilet       `gorm:"foreignKey:Place_Class_ID"`
	ReadingZones  []ReadingZone  `gorm:"foreignKey:Place_Class_ID"`
	ResearchRooms []ResearchRoom `gorm:"foreignKey:Place_Class_ID"`
	Computers     []Computer     `gorm:"foreignKey:Place_Class_ID"`
}

type Relation struct {
	gorm.Model
	Place_Class_ID *uint
	Problem_ID     *uint
	//JOIN
	Place_Class Place_Class
	Problem     Problem
}

type Toilet struct {
	gorm.Model
	Name string
	// Place_Problem_ID ทำหน้าที่เป็น FK
	Place_Class_ID *uint
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`
}

type ReadingZone struct {
	gorm.Model
	Name string
	// Place_Problem_ID ทำหน้าที่เป็น FK
	Place_Class_ID *uint
	ProblemReports []ProblemReport `gorm:"foreignKey:Problem_ID"`
}

type ProblemReport struct {
	gorm.Model

	USER_ID *uint
	User    User

	Problem_ID *uint
	Problem    Problem

	RdZone_id *uint
	RdZone    ReadingZone

	Tlt_id *uint
	Tlt    Toilet

	ReschRoom_id *uint
	ReschRoom    ResearchRoom

	Com_id *uint
	Com    Computer

	Place_Class_ID *uint
	Place_Class    Place_Class

	Comment string
}
