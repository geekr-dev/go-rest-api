package main

import (
	"net/http"

	"github.com/geekr-dev/go-rest-api/config"
	"github.com/geekr-dev/go-rest-api/model"
	"github.com/geekr-dev/go-rest-api/pkg/log"
	"github.com/geekr-dev/go-rest-api/router"
	"github.com/geekr-dev/go-rest-api/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
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

	// init logger
	log.Init(config.Data.Log)
	defer log.Close()

	// init db
	model.DB.Init(config.Data.Db)

	// create gin engine
	gin.SetMode(config.Data.Mode)
	g := gin.New()

	// middlewares
	// middlewares := []gin.HandlerFunc{}

	// routes
	router.Load(
		g,
		// middlewares...,
		middleware.Logging(),
		middleware.RequestId(),
	)

	// start server: tls 证书不为空则启动 https
	if config.Data.Tls.Cert != "" && config.Data.Tls.Key != "" {
		log.Info("Start to listening incoming requests on http address: %s", config.Data.Tls.Addr)
		log.Info(http.ListenAndServeTLS(config.Data.Tls.Addr, config.Data.Tls.Cert, config.Data.Tls.Key, g).Error())
	} else {
		log.Info("Start to listening incoming requests on http address: %s", config.Data.Addr)
		log.Info(http.ListenAndServe(config.Data.Addr, g).Error())
	}
}
