package main

import (
	"log"
	"net/http"

	"github.com/geekr-dev/go-rest-api/config"
	"github.com/geekr-dev/go-rest-api/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {
	// init config
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// create gin engin
	gin.SetMode(viper.GetString("mode"))
	g := gin.New()

	// middlewares
	middlewares := []gin.HandlerFunc{}

	// routes
	router.Load(
		g,
		middlewares...,
	)

	// start server
	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
