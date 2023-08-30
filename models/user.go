package models

import (
	"encoding/xml"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `xml:"name"`
	Lastname string `xml:"lastname"`
	Address  string `xml:"address"`
	Contact  string `xml:"contact"`
	Email    string `xml:"email"`
	Password string `xml:"password"`
}

type UsersInfo struct {
	gorm.Model
	CorporateID      string `xml:"corporate"`
	BranchID         string `xml:"branchid"`
	TransactionKey   string `xml:"transactionkey"`
	RequestRefNo     string `xml:"requestno"`
	TransactionType  string `xml:"transactiontype"`
	RequestTimeStamp string `xml:"requesttimestamp"`
	TerminalID       string `xml:"terminalid"`
	Address          string `xml:"address"`
}

type Deped struct {
	XMLName       xml.Name `xml:"DepedHead"`
	School        string   `xml:"DepedSchool"`
	PrincipalHead PrincipalHead
}

type PrincipalHead struct {
	XMLName     xml.Name `xml:"soapenv:Header"`
	Principal   string   `xml:"principal"`
	TeacherBody TeacherBody
}

type TeacherBody struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Teacher string   `xml:"teacher"`
	Student Student
}

type Student struct {
	gorm.Model
	XMLName      xml.Name `xml:"Deped:StudentInformation"`
	StudentName  string   `xml:"studentname"`
	StudentID    string   `xml:"studentid"`
	Section      string   `xml:"section"`
	MajorSubject string   `xml:"majorsubject"`
}

// type Employee struct {
// 	ID     int    `xml:"id"`
// 	Name   string `xml:"name"`
// 	Salary int    `xml:"salary"`
// }

// type User1 struct {
// 	ID       int    `xml:"primaryKey"`
// 	Name     string `xml:"unique"`
// 	Password string
// 	// Add other fields as needed
// }

type LoginResponse struct {
	XMLName xml.Name `xml:"response"`
	Message string   `xml:"message"`
}
