package main

import (
	"github.com/gin-gonic/gin"
)

type Item struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool `jsojn:"completed"`
}

var Items = []Item{}

func main() {
	router := gin.Default()

	// localhost:8080/items
	router.GET("/items", func(c *gin.Context) {
		c.IndentedJSON(200, Items)
	})

	router.POST("/items", func(c *gin.Context) {
		var newItem Item
		if err := c.BindJSON(&newItem); err != nil {
			return
		}
		Items = append(Items, newItem)
		c.IndentedJSON(200, newItem)
	})

	router.Run(":8080")
}