package controllers

import (
	"github.com/Funskie/todoListApp/helpers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Healthz(c *gin.Context) {
	log.Info("API health is OK!")
	type log struct {
		Alive bool `json:"alive"`
	}
	helpers.WrapResponse(c, log{Alive: true}, nil)
}
