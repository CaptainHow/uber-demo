package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)



type Trip struct {
	Tripid, Source, Destination string 
	Cost float64 
	PaymentMethod string
	Status int
}


type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rating int `json:"rating"`
	Trips []Trip `json:"trips"`
}

func (trip Trip) String() string {
	return fmt.Sprintf("TripId: %s", trip.Tripid)
}

var customersArray = make([]Customer, 0)

func (cust *Customer)Save() string {
	cust.Id = uuid.NewString()
	customersArray = append(customersArray, *cust)
	fmt.Println(customersArray)
	return cust.Id
}

func createCustomer(c *gin.Context) {
	var customer Customer
	if err := c.BindJSON(&customer); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Parameters!",
		})
		return
	}
	var custID = customer.Save()
	c.JSON(http.StatusOK, gin.H{
		"Message": custID,
	})
}

func main() {
	// fmt.Println(trip)
	r := gin.Default()
	r.GET("/ping",func(c *gin.Context) {
		trip := Trip{Tripid: "1", Source: "2", Destination: "3", Cost: 123.0, PaymentMethod: "abc", Status: 0}
		c.JSON(200, gin.H{
			"message": "pong",
			"trip":trip,
		})
	})
	r.POST("/v1/customer",createCustomer)
	// r.GET("/v1/customer",createUser)
	// r.POST("/v1/trip",createUser)
	r.Run()
}/*  */