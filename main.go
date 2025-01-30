package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Restaurant struct
type Restaurant struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Menu []string `json:"menu"`
}

var restaurants = []Restaurant{
	{ID: 1, Name: "Sadguru Special", Menu: []string{"Paneer Tikka", "Dal Makhani", "Butter Naan"}},
	{ID: 2, Name: "Tasty Bites", Menu: []string{"Veg Biryani", "Chole Bhature", "Gulab Jamun"}},
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Sadguru Catering!"})
	})

	r.GET("/restaurants", func(c *gin.Context) {
		c.JSON(http.StatusOK, restaurants)
	})

	r.GET("/restaurant/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, r := range restaurants {
			if fmt.Sprintf("%d", r.ID) == id {
				c.JSON(http.StatusOK, r)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
	})

	r.Run(":8080") // Start server
}
