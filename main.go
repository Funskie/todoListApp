package main

import (
	"github.com/Funskie/todoListApp/controllers"
	"github.com/Funskie/todoListApp/helpers"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting Todolist API server")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) { helpers.WrapResponse(c, "Todolist API", nil) })
	router.GET("/healthz", controllers.Healthz)

	router.Run(":80")
}
