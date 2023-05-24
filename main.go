package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string	`json:"lastName"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address string `json:"address"`
	ContractType string `json:"contractType"`
	OnGoing bool `json:"onGoing"`
	DatesEmployedFirst time.Time `json:"datesEmployedFirst"`
	DatesEmployedEnd time.Time `json:"datesEmployedEnd"`
	EmploymentType string `json:"employmentType"`
	HoursPW int `json:"hoursPW"`
}

var employees = []Employee{
 {Id: 1, FirstName: "John", LastName: "Smith", Email: "johnsmith123@gmail.com", PhoneNumber: "04123456789", Address: "1 Fake Street, Faker, NSW, 2000", ContractType: "Permanent", OnGoing:  false, EmploymentType: "Full-Time", HoursPW: 35, DatesEmployedFirst: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.UTC), DatesEmployedEnd: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.Local)},
 {Id: 2, FirstName: "Johner", LastName: "Smith", Email: "johnsmith123@gmail.com", PhoneNumber: "04123456789", Address: "1 Fake Street, Faker, NSW, 2000", ContractType: "Permanent", OnGoing:  false, EmploymentType: "Full-Time", HoursPW: 35, DatesEmployedFirst: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.UTC), DatesEmployedEnd: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.Local)},
 {Id: 1, FirstName: "Johnane", LastName: "Smith", Email: "johnsmith123@gmail.com", PhoneNumber: "04123456789", Address: "1 Fake Street, Faker, NSW, 2000", ContractType: "Permanent", OnGoing:  true, EmploymentType: "Full-Time", HoursPW: 35, DatesEmployedFirst: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.UTC), DatesEmployedEnd: time.Now()},
}

// CRUD

// GET

func getEmployees(c *gin.Context){
	c.IndentedJSON(http.StatusOK, employees)
}

// POST

func main(){
	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.Run("localhost:8080")
}