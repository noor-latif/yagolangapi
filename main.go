package main

import (
	"errors"
	"net/http"
	"strconv"

	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"systementor.se/yagolangapi/data"
)

// Define a struct for the page view
type PageView struct {
	Title  string
	Rubrik string
}

// Declare a global variable for the random number generator
var theRandom *rand.Rand

// Define a function to start the application
func start(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "test", Rubrik: "Hej Golang"})
}

// Define a function to return employees in JSON format
func employeesJson(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.JSON(http.StatusOK, employees)
}

// Define a function to add a single employee
func addEmployee(c *gin.Context) {
	data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})
}

// Define a function to add multiple employees
func addManyEmployees(c *gin.Context) {
	//Here we create 10 Employees
	for i := 0; i < 10; i++ {
		data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})
	}
}

// Define a function to return employees in indented JSON format
func apiEmployee(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.IndentedJSON(http.StatusOK, employees)
}

// Define a function to return an employee by ID in indented JSON format
func apiEmployeeById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

// Define a function to update an employee by ID in indented JSON format
func apiEmployeeUpdateById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		if err := c.BindJSON(&employee); err != nil {
			return
		}
		employee.Id, _ = strconv.Atoi(id)
		data.DB.Save(&employee)
		c.IndentedJSON(http.StatusOK, employee)
	}
}

// Define a function to delete an employee by ID in indented JSON format
func apiEmployeeDeleteById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		data.DB.Delete(&employee)
		c.IndentedJSON(http.StatusNoContent, employee)
	}
}

// Define a function to add an employee in indented JSON format
func apiEmployeeAdd(c *gin.Context) {
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id = 0
	err := data.DB.Create(&employee).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, employee)
	}
}

// Define a function to return a list of all employees' name and city
func apiNoor(c *gin.Context) {
	var employees []data.Employee
	data.DB.Select("Namn, City").Find(&employees)
	c.IndentedJSON(http.StatusOK, employees)
}

func apiTest(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

// Declare a global variable for the configuration
//var config Config

// Define the main function
func main() {
	// Initialize the random number generator
	theRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Read the configuration file
	//	readConfig(&config)

	// Initialize the database
	/* data.InitDatabase(config.Database.File,
	config.Database.Server,
	config.Database.Database,
	config.Database.Username,
	config.Database.Password,
	config.Database.Port) */

	// Initialize the router
	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/api/employee", apiEmployee)
	router.GET("/api/employee/:id", apiEmployeeById)
	router.PUT("/api/employee/:id", apiEmployeeUpdateById)
	router.DELETE("/api/employee/:id", apiEmployeeDeleteById)
	router.POST("/api/employee", apiEmployeeAdd)
	router.GET("/api/employees", employeesJson)
	router.GET("/api/addemployee", addEmployee)
	router.GET("/api/addmanyemployees", addManyEmployees)
	router.GET("/api/noor", apiNoor)
	router.GET("/api/test", apiTest)
	router.Run(":8080")
}
