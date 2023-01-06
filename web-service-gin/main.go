package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func loadDatabase() {
	Connect()
}

func loadEnv() {
	err := godotenv.Load("cust.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func serveApplication() {
	router := gin.Default()
	router.GET("/customers/:id", getCustomer)
	router.POST("/customers/:id", insertCustomer)
	router.PUT("/customers/:id", updateCustomer)
	router.DELETE("/customers/:id", deleteCustomer)
	router.Run("localhost:8080")
	fmt.Println("Server running on port 8000")
}

type customer struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Accbal float64 `json:"accbal"`
}

func getCustomerById(c *gin.Context) (customer, error) {
	id := c.Param("id")
	fmt.Println("Here is the ID", id)
	var customer customer

	err := Database.Where("ID=?", id).Find(&customer).Error
	fmt.Println("Error is ", err)
	return customer, err
}

// getcustomer responds with the list of all records as JSON.
func getCustomer(c *gin.Context) {
	var getcustomer customer
	var errors error
	getcustomer, errors = getCustomerById(c)
	fmt.Println("get customer response is ", getcustomer)
	if errors != nil {
		c.IndentedJSON(http.StatusNotFound, "")
	}
	c.IndentedJSON(http.StatusOK, getcustomer)
}
func insertCustomer(c *gin.Context) {
	/* reading the uri parameter to get the ID */

	id := c.Param("id")
	fmt.Println("Here is the ID", id)

	/* Reading the request body to get new name and new balance */

	var reqBody customer
	if err := c.BindJSON(&reqBody); err != nil {
		return
	}

	fmt.Println("Here is the New Name and Account balance", reqBody.Name, reqBody.Accbal)
	var updatedCustomer customer

	sqlStatement := `insert customers SET Name = $2, Accbal= $3 WHERE id = $1 RETURNING id, Name, Accbal ;`
	result := Database.Exec(sqlStatement, id, reqBody.Name, reqBody.Accbal)

	if result != nil {
		err := Database.Where("ID=?", id).Find(&updatedCustomer).Error
		fmt.Println("Error is ", err)
		c.IndentedJSON(http.StatusOK, updatedCustomer)
	}

}

func updateCustomer(c *gin.Context) {
	/* reading the uri parameter to get the ID */

	id := c.Param("id")
	fmt.Println("Here is the ID", id)

	/* Reading the request body to get new name and new balance */

	var reqBody customer
	if err := c.BindJSON(&reqBody); err != nil {
		return
	}

	fmt.Println("Here is the New Name and Account balance", reqBody.Name, reqBody.Accbal)
	var updatedCustomer customer

	sqlStatement := `UPDATE  customers SET Name = $2, Accbal= $3 WHERE id = $1 RETURNING id, Name, Accbal ;`
	result := Database.Exec(sqlStatement, id, reqBody.Name, reqBody.Accbal)

	if result != nil {
		err := Database.Where("ID=?", id).Find(&updatedCustomer).Error
		fmt.Println("Error is ", err)
		c.IndentedJSON(http.StatusOK, updatedCustomer)
	}

}

func deleteCustomer(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("Here is the ID", id)

	sqlStatement := `DELETE FROM customers WHERE id = $1 ;`
	err := Database.Exec(sqlStatement, id)
	if err != nil {
		c.IndentedJSON(http.StatusOK, "")
	}

}
