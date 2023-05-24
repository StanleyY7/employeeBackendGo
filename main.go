package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id string `json:"id"`
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
 {Id: "1", FirstName: "John", LastName: "Smith", Email: "johnsmith123@gmail.com", PhoneNumber: "04123456789", Address: "1 Fake Street, Faker, NSW, 2000", ContractType: "Permanent", OnGoing:  false, EmploymentType: "Full-Time", HoursPW: 35, DatesEmployedFirst: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.UTC), DatesEmployedEnd: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.Local)},
 {Id: "2", FirstName: "Johner", LastName: "Smith", Email: "johnsmith123@gmail.com", PhoneNumber: "04123456789", Address: "1 Fake Street, Faker, NSW, 2000", ContractType: "Permanent", OnGoing:  false, EmploymentType: "Full-Time", HoursPW: 35, DatesEmployedFirst: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.UTC), DatesEmployedEnd: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.Local)},
 {Id: "3", FirstName: "Johnane", LastName: "Smith", Email: "johnsmith123@gmail.com", PhoneNumber: "04123456789", Address: "1 Fake Street, Faker, NSW, 2000", ContractType: "Permanent", OnGoing:  true, EmploymentType: "Full-Time", HoursPW: 35, DatesEmployedFirst: time.Date(2023, time.May, 24, 12, 30, 0, 0, time.UTC), DatesEmployedEnd: time.Now()},
}

// CRUD

// GET

func getEmployees(c *gin.Context){
	c.IndentedJSON(http.StatusOK, employees)
}

// get by id

func getEmployeeById(id string)(*Employee, error){
for index, item := range employees {
if (item.Id == id) {
	return &employees[index], nil
}
}
return nil, errors.New("Employee not found")
}

func EmployeeById(c *gin.Context){
	id := c.Param("id")
	Employee, err := getEmployeeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found :("})
		return
	}

	c.IndentedJSON(http.StatusOK, Employee)
	
}

// POST

func postEmployees(c *gin.Context){
	var newEmployee Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to make new employee, please try again later"})
		return
	}

	employees = append(employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

// PATCH

func patchEmployee(c *gin.Context){

	id := c.Param("id")

	employee, _ := getEmployeeById(id)

	if employee == nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Employee not found"})
		return
	}

	var updateEmployee Employee

	if err := c.ShouldBindJSON(&updateEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Bad Request"})
		return
	}

	if updateEmployee.FirstName != ""{
		employee.FirstName = updateEmployee.FirstName 
	}
	if updateEmployee.LastName != ""{
		employee.LastName = updateEmployee.LastName 
	}
	if updateEmployee.PhoneNumber != ""{
	employee.PhoneNumber = updateEmployee.PhoneNumber
	}
	if updateEmployee.Email != ""{
		employee.Email = updateEmployee.Email
	}
	if updateEmployee.Address != ""{
		employee.Address = updateEmployee.Address
	}
	if updateEmployee.ContractType != ""{
		employee.ContractType = updateEmployee.ContractType
	}
	if updateEmployee.EmploymentType != ""{
		employee.EmploymentType = updateEmployee.EmploymentType
	}
	if updateEmployee.HoursPW != 0 {
		employee.HoursPW = updateEmployee.HoursPW
	}
	if updateEmployee.OnGoing != false || true {
		employee.OnGoing = updateEmployee.OnGoing
	}
	if updateEmployee.DatesEmployedFirst !=  time.Date(1980, time.January, 1, 12, 30, 0, 0, time.Local){
		employee.DatesEmployedFirst = updateEmployee.DatesEmployedFirst
	}
	if updateEmployee.DatesEmployedEnd !=  time.Date(1980, time.January, 1, 12, 30, 0, 0, time.Local){
		employee.DatesEmployedEnd = updateEmployee.DatesEmployedEnd
	}

	c.IndentedJSON(http.StatusOK, employee)
}

// DELETE BY ID

func deleteEmployee(c *gin.Context){
	
	id := c.Param("id")

	index := -1

	// Find employee by index, checking it exists or not

	for i, element := range employees {
		if element.Id == id {
			index = i
			break
		}
	}
	
	// returns error if it doesn't exist

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	// returns new slice without the employee
	// employees[:index] is new slice of all employees before index of one we are deleting
	// employees[index+1:] is new slice of all employees after the index we are deleting
	// employees = is reassigning the variable with the 2 slices in one. 

	employees = append(employees[:index], employees[index+1:]... )

	c.IndentedJSON(http.StatusOK, gin.H{"message":"Employee deleted"})
}


func main(){
	router := gin.Default()
	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", EmployeeById)
	router.POST("/employees", postEmployees)
	router.DELETE("/employees/:id", deleteEmployee)
	router.PATCH("/employees/:id", patchEmployee)
	router.Run("localhost:8080")
}