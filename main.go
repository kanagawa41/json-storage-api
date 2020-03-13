package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Get stock
	router.GET("/api/v1/stocks/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")

		db, err := connection()
		if err != nil {
			serverErrorHandler(c, err)
			return
		}
		defer db.Close()

		s, err := selectStock(db, uuid)

		if err == sql.ErrNoRows {
			c.JSON(400, gin.H{
				"error":  "Not found",
				"result": nil,
			})
			return
		}

		c.JSON(200, gin.H{
			"error":  nil,
			"result": s,
		})
	})

	// Create stock
	router.POST("/api/v1/stocks", func(c *gin.Context) {
		db, err := connection()
		if err != nil {
			serverErrorHandler(c, err)
			return
		}
		defer db.Close()

		las := lastID(db)
		s, err := createStock(db, las+1)

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
	db, err := connection()
	if err != nil {
		serverErrorHandler(c, err)
		return
	}
	defer db.Close()

	_, err = updateStock(db, uuid, json)
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
	db, err := connection()
	if err != nil {
		serverErrorHandler(c, err)
		return
	}
	defer db.Close()

	_, err = deleteStock(db, uuid)
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
