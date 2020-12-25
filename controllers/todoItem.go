package controllers

import (
	"github.com/Funskie/todoListApp/helpers"
	"github.com/Funskie/todoListApp/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func CreateTodoItem(c *gin.Context) {
	description := c.PostForm("description")
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
	res, err := models.GetLastItem()
	helpers.WrapResponse(c, res, err)
}

func UpdateTodoItem(c *gin.Context) {
	type updateLog struct {
		Update bool `json:"update"`
	}
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	todo, err := models.GetItemByID(id)
	if err != nil {
		log.Warn("TodoItem with ID:", id, "not found in database")
		helpers.WrapResponse(c, nil, err)
		return
	}

	oldCompleted := todo.Completed
	completed, err := strconv.ParseBool(c.PostForm("completed"))
	if err == nil {
		if completed != oldCompleted {
			log.WithFields(log.Fields{"ID": id, "Completed": completed}).Info("Updating TodoItem")
			todo.Completed = completed
			err = todo.UpdateCompleted()
			if err != nil {
				helpers.WrapResponse(c, nil, err)
				return
			}
			helpers.WrapResponse(c, updateLog{Update: true}, err)
			return
		} else {
			log.WithFields(log.Fields{"ID": id, "Completed": completed}).Info("TodoItem completed doesn't changed")
			helpers.WrapResponse(c, updateLog{Update: false}, err)
		}
	} else {
		helpers.WrapResponse(c, nil, err)
		return
	}
}

func GetCompletedItems(c *gin.Context) {
	log.Info("Get completed TodoItems")
	completedItems, err := models.GetTodoItem(true)
	if err != nil {
		helpers.WrapResponse(c, nil, err)
		return
	}
	helpers.WrapResponse(c, completedItems, err)
}

func GetIncompletedItems(c *gin.Context) {
	log.Info("Get incompleted TodoItems")
	incompletedItems, err := models.GetTodoItem(false)
	if err != nil {
		helpers.WrapResponse(c, nil, err)
		return
	}
	helpers.WrapResponse(c, incompletedItems, err)
}

func DeleteItem(c *gin.Context) {
	type deleteLog struct {
		Delete bool `json:"delete"`
	}
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	todo, err := models.GetItemByID(id)
	if err != nil {
		log.Warn("TodoItem with ID:", id, "not found in database")
		helpers.WrapResponse(c, nil, err)
		return
	}

	log.WithFields(log.Fields{"ID": id}).Info("Deleting TodoItem")
	err = todo.Delete()
	if err != nil {
		helpers.WrapResponse(c, nil, err)
		return
	}
	helpers.WrapResponse(c, deleteLog{Delete: true}, err)
}
