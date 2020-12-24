package main

import (
	"github.com/Funskie/todoListApp/controllers"
	"github.com/Funskie/todoListApp/helpers"
	"github.com/Funskie/todoListApp/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Error("err open databases", err)
		return
	}
	defer db.Close()

	log.Info("Starting Todolist API server")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) { helpers.WrapResponse(c, "Todolist API", nil) })
	router.GET("/healthz", controllers.Healthz)

	router.POST("/todo", controllers.CreateTodoItem)

	router.Run(":80")
}
