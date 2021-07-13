package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CustomMiddlewareOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CustomMiddlewareOne start")
		c.Next()
		fmt.Println("CustomMiddlewareOne end")
	}
}

func CustomMiddlewareTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CustomMiddlewareTwo start")
		c.Next()
		fmt.Println("CustomMiddlewareTwo end")
	}
}

func main() {
	r := gin.Default()
	r.Use(CustomMiddlewareOne())
	r.Use(CustomMiddlewareTwo())

	r.GET("/test", func(c *gin.Context) {
		fmt.Println("test run")
		c.String(200, "this is a test")
	})

	r.Use(CustomMiddlewareTwo())

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
