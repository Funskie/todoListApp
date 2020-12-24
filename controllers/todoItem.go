package controllers

import (
	"github.com/Funskie/todoListApp/helpers"
	"github.com/Funskie/todoListApp/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateTodoItem(c *gin.Context) {
	description := c.PostForm("description")
	// completed := c.PostForm("completed")
	log.WithFields(log.Fields{"description": description}).Info("Add new TodoItem. Saving to database.")
	todo := &models.TodoItemModel{
		Description: description,
		Completed:   false,
	}
	err := todo.Insert()
	if err != nil {
		helpers.WrapResponse(c, nil, err)
		return
	}
	res, err := todo.GetLastItem()
	helpers.WrapResponse(c, res, err)
}
