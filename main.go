package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {
	runGinRouter()
}

func initializeDB() *pgx.Conn {
	connUrl := "postgres://xpjczzaqerighx:e0291b716ea919b11e07c2de123a3164306e2cb0c721801d9b037404c620d39a@ec2-35-168-80-116.compute-1.amazonaws.com:5432/dcgmgicspgnus3"
	// connUrl := "postgres://cty:password@localhost:5432/goworlddb"
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func runGinRouter() {
	user := User{ID: 0, FirstName: "Cassiano", LastName: "Yasumitsu", Email: "cassiano@email.com"}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static/", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "GoWorld - Main Page", "payload": user},
		)
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "GoWorld - Main Page", "payload": user},
		)
	})
	router.GET("/hotels", func(c *gin.Context) {
		conn := initializeDB()
		defer conn.Close(context.Background())

		var allHotels []Hotel
		if rows, err := conn.Query(context.Background(), "SELECT id, country, destination, hotel FROM hotels"); err != nil {
			fmt.Println("Unabel to get list of hotels due to: ", err)
		} else {
			// deferring query closing
			defer rows.Close()

			// Using tmp variable for reading
			var hotel Hotel

			// Next prepares the next row for reading.
			for rows.Next() {
				// Scan reads the values from the current row into tmp
				rows.Scan(&hotel.ID, &hotel.Country, &hotel.Destination, &hotel.Name)
				allHotels = append(allHotels, hotel)
			}
		}

		c.HTML(
			http.StatusOK,
			"hotels.html",
			gin.H{"title": "GoWorld - Hotels", "payload": user, "hotels": allHotels, "menuHotels": true},
		)
	})
	router.GET("/contact", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"contact.html",
			gin.H{"title": "GoWorld - Main Page"},
		)
	})

	// api routes
	v1 := router.Group("/v1/hotels")
	{
		v1.POST("/", createHotel)
		v1.GET("/", getAllHotels)
		v1.GET("/:id", getHotelByID)
	}

	router.Run()
}
