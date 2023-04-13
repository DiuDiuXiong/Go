package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

/**
Gin is a web framework written in Go that is used to build web applications and APIs.
It provides a set of tools and libraries to help developers quickly build high-performance web applications in Go.
Some of the key features of Gin include routing, middleware, JSON parsing, and error handling. Gin is designed to be
lightweight and efficient, and it is often used in production environments because of its speed and reliability.

go get -u github.com/gin-gonic/gin
*/

func main() {
	r := gin.Default() // return a gin engin

	// r.Use() to add middle-ware
	// No matter which URL it is requesting, will go through these first

	// log middle-ware
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(c *gin.Context) {
		// log latency, response code, path
		t := time.Now()
		c.Next() // let it continue to run, this function means run the actual query

		// we put the logging at back since we want to get the response code, which only available after c finish
		// path
		logger.Info("Incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", time.Now().Sub(t)),
		)
		// response code
	}, func(c *gin.Context) { // another middleware which add a request id
		c.Set("requestId", rand.Int())
		c.Next()
	})

	/*
		the Next() function call in the first middleware function allows the control to pass to the next middleware function
		in the chain. In this case, the next middleware function is the one that adds the request ID to the context.
		So the request ID middleware function is still executed, and the request ID is set in the context before the handler function is executed.

		Yes, the middlewares added using r.Use() are executed in the order they are added. So the first middleware added
		will be executed first, then the second, and so on. Each middleware can modify the request or response before
		passing it to the next middleware in the chain. If a middleware does not call c.Next(), the chain will be
		terminated and no further middlewares will be executed.
	*/

	// Adding resolve functions for request of different URLs
	// when request /ping, return json with pong
	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}

		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "/hello")
	})

	r.Run()
}
