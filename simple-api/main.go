package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddRequest struct {
	A float64 `json:"a" binding:"required"`
	B float64 `json:"b" binding:"required"`
}

type AddResponse struct {
	Sum float64 `json:"sum"`
}

func main() {
	// gin.Default() includes:
	// - Logger middleware (logs requests)
	// - Recovery middleware (prevents crashes on panics)
	r := gin.Default()

	// GET /hello
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Developer",
		})
	})

	// POST /add
	r.POST("/add", func(c *gin.Context) {
		var req AddRequest

		// Bind JSON body into req, and validate required fields
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid input. Expected JSON: {\"a\": number, \"b\": number}",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, AddResponse{
			Sum: req.A + req.B,
		})
	})

	// Start server on localhost:8080
	r.Run(":8080")
}
