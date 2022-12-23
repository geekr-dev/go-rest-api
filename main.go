package main

import (
	"log"
	"net/http"

	"github.com/geekr-dev/go-rest-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// create gin engine
	g := gin.New()

	// middlewares
	middlewares := []gin.HandlerFunc{}

	// routes
	router.Load(
		g,
		middlewares...,
	)

	// start server
	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}
