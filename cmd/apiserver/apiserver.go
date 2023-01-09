package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/yxtiblya/internal/apiserver/auth"
	"github.com/yxtiblya/internal/apiserver/routehandler"
	"github.com/yxtiblya/internal/cfg"
	"github.com/yxtiblya/internal/store"
)

func main() {
	// parsing toml file to config
	config := cfg.NewConfig()
	if _, err := toml.DecodeFile("configs/config.toml", config); err != nil {
		panic("failed to decode toml file")
	}
	cfg.ChangeConfig(config)

	// connection to db
	_, err := store.NewDB(config)
	if err != nil {
		panic("failed to connect database")
	}

	// logger writer in file and console
	f, _ := os.Create("logger/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// initialize custom logger middleware
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	// URLS
	//
	// returning jwt
	r.GET("/api/getJWT", routehandler.GetJWT)

	// urls with JWT auth
	authorized := r.Group("/")
	authorized.Use(auth.ValidateJWT)
	{
		authorized.POST("/api/contact/post", routehandler.Contact_post)
		authorized.GET("/api/contact/get", routehandler.Contact_get)
		authorized.PUT("/api/contact/put", routehandler.Contact_put)
		authorized.DELETE("/api/contact/delete", routehandler.Contact_delete)

		authorized.POST("/api/mailing/post", routehandler.Mailing_post)
		authorized.GET("/api/mailing/get", routehandler.Mailing_get)
		authorized.PUT("/api/mailing/put", routehandler.Mailing_put)
		authorized.DELETE("/api/mailing/delete", routehandler.Mailing_delete)

		authorized.POST("/api/message/post", routehandler.Message_post)
		authorized.GET("/api/message/get", routehandler.Message_get)
		authorized.PUT("/api/message/put", routehandler.Message_put)
		authorized.DELETE("/api/message/delete", routehandler.Message_delete)

		authorized.GET("/api/stat/get/:id", routehandler.Stat_get)
	}

	r.Run(":" + config.BindAddr)
}
