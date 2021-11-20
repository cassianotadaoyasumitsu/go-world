package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createHotel(c *gin.Context) {
	conn := initializeDB()
	defer conn.Close(context.Background())

	var hotel Hotel

	if err := c.ShouldBindJSON(&hotel); err == nil {
		row := conn.QueryRow(context.Background(), "INSERT INTO hotels(country, destination, hotel VALUES($1, $2, $3) return id", hotel.Country, hotel.Destination, hotel.Name)
		var id int
		err = row.Scan(&id)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		c.JSON(http.StatusAccepted, gin.H{"status": "created"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to add new hotel due to ": err.Error()})
	}
}

func getAllHotels(c *gin.Context) {
	conn := initializeDB()
	defer conn.Close(context.Background())

	var result []Hotel
	if hotels, err := conn.Query(context.Background(), "SELECT id, country, destination, hotel FROM hotels"); err != nil {
		fmt.Println("Unable to get list of hotels due to: ", err)
	} else {
		defer hotels.Close()

		var hotel Hotel

		for hotels.Next() {
			hotels.Scan(&hotel.ID, &hotel.Country, &hotel.Destination, &hotel.Name)
			result = append(result, hotel)
		}
		if result != nil {
			c.JSON(http.StatusOK, gin.H{"result": result})
		} else {
			c.JSON(http.StatusNotFound, nil)
		}
		if hotels.Err() != nil {
			fmt.Println("Error while reading hotels table: ", err)
		}
	}
}

func getHotelByID(c *gin.Context) {
	conn := initializeDB()
	defer conn.Close(context.Background())

	hotelID := c.Param("id")

	if hotels, err := conn.Query(context.Background(), "SELECT id, country, destination, hotel FROM hotels h WHERE h.ID = $1", hotelID); err != nil {
		fmt.Println("Unable to get a hotel by ID due to: ", err)
	} else {
		defer hotels.Close()

		var hotel Hotel

		if hotels.Next() {
			hotels.Scan(&hotel.ID, &hotel.Country, &hotel.Destination, &hotel.Name)
			c.JSON(http.StatusOK, gin.H{"result": hotel})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"result": nil})
		}
		if hotels.Err() != nil {
			fmt.Println("Error while reading hotels table: ", err)
		}
	}
}
