package main

import (
	"context"
	controller "sekawan-exercise/app/main/controller"
	"sekawan-exercise/app/main/server"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/gin-contrib/cors"
)

var db = make(map[string]string)

func main() {

	server.InitConfig()
	server.InitLogrusFormat()

	// running open telemetry
	cleanup := server.InitTracer()
	defer cleanup(context.Background())

	gin.SetMode(server.GIN_MODE)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(server.RequestResponseLogger())
	router.Use(otelgin.Middleware(server.APP_NAME))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	router.Use(cors.New(corsConfig))

	_ = server.InitPostgreSQL()
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.Static("/scss", "./static/scss")
	router.Static("/vendor", "./static/vendor")
	router.Static("/js", "./static/js")
	router.StaticFile("/favicon.ico", "./img/favicon.ico")

	router.LoadHTMLGlob("templates/**/*")
	router.GET("/", controller.Index)
	router.GET("/contact", controller.Contact)
	router.GET("/about", controller.About)

	router.Run(":" + server.HTTP_SERVER_PORT)
}
