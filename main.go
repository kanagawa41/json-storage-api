package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := os.MkdirAll(baseURL, 0777); err != nil {
		panic(err)
	}

	// Get stock
	router.GET("/api/v1/stocks/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")

		_, err := selectStock(uuid)

		if err != nil {
			serverErrorHandler(c, err)
			return
		}

		c.File(getFilePath(uuid))
	})

	// Create stock
	router.POST("/api/v1/stocks", func(c *gin.Context) {
		s, err := createStock()

		if err != nil {
			serverErrorHandler(c, err)
			return
		}

		c.JSON(200, gin.H{
			"error":  nil,
			"result": s,
		})
	})

	// Update or delete stock
	router.POST("/api/v1/stocks/:uuid/*action", func(c *gin.Context) {
		action := c.Param("action")

		uuid := c.Param("uuid")
		json := c.PostForm("json")
		fmt.Println(json)
		switch action {
		case "/put":
			updateStockHandler(c, uuid, json)
		case "/delete":
			deleteStockHandler(c, uuid)
		default:
			c.JSON(404, gin.H{
				"error":  "404 page not found",
				"result": nil,
			})
		}
	})

	router.Run(":8090")
}

func updateStockHandler(c *gin.Context, uuid string, json string) {
	_, err := updateStock(uuid, json)
	if err != nil {
		serverErrorHandler(c, err)
		return
	}

	c.JSON(200, gin.H{
		"error":  nil,
		"result": len(json),
	})
}
func deleteStockHandler(c *gin.Context, uuid string) {
	_, err := deleteStock(uuid)
	if err != nil {
		serverErrorHandler(c, err)
		return
	}

	c.JSON(200, gin.H{
		"error":  nil,
		"result": true,
	})
}

func serverErrorHandler(c *gin.Context, err error) {
	fmt.Println(err)

	c.JSON(500, gin.H{
		"error":  "Server Error",
		"result": nil,
	})
}
