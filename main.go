package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve HTML landing page
	router.LoadHTMLGlob("/Users/aviranmashiach/Documents/go/templates/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Calculate endpoint
	router.POST("/calculate", func(c *gin.Context) {
		num1Str := c.PostForm("num1")
		num2Str := c.PostForm("num2")
		operation := c.PostForm("operation")

		num1, _ := strconv.ParseFloat(num1Str, 64)
		num2, _ := strconv.ParseFloat(num2Str, 64)

		var result float64
		switch operation {
		case "add":
			result = num1 + num2
		case "subtract":
			result = num1 - num2
		case "multiply":
			result = num1 * num2
		case "divide":
			if num2 != 0 {
				result = num1 / num2
			} else {
				c.JSON(400, gin.H{"error": "division by zero"})
				return
			}
		}

		c.JSON(200, gin.H{"result": result})
	})

	router.Run(":8080")
}
