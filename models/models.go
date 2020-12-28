package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TodoItemModel struct {
	ID          uint `gorm:"primary_key"`
	Description string
	Completed   bool
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/todolist?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		db.LogMode(true)
		db.DropTableIfExists(&TodoItemModel{})
		db.AutoMigrate(&TodoItemModel{})
		return db, err
	}
	return nil, err
}

func (item *TodoItemModel) Insert() error {
	return DB.Create(item).Error
}

func (item *TodoItemModel) UpdateDescription() error {
	return DB.Model(item).Update(map[string]interface{}{
		"description": item.Description,
	}).Error
}

func (item *TodoItemModel) UpdateCompleted() error {
	return DB.Model(item).Update(map[string]interface{}{
		"completed": item.Completed,
	}).Error
}

func (item *TodoItemModel) Delete() error {
	return DB.Delete(item).Error
}

func GetLastItem() (*TodoItemModel, error) {
	var item TodoItemModel
	err := DB.Last(&item).Error
	return &item, err
}

func GetItemByID(Id int) (*TodoItemModel, error) {
	var item TodoItemModel
	err := DB.First(&item, Id).Error
	return &item, err
}

func GetTodoItem(completed bool) (*[]TodoItemModel, error) {
	var todos []TodoItemModel
	err := DB.Where("completed = ?", completed).Find(&todos).Error
	return &todos, err
}
