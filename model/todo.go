package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/magicxiaobao/go-todo/config"
)

type Todo struct {
	ID uint 	`json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

func (todo *Todo) TableName() string {
	return "todo"
}

func GetAllTodos(todo *[]Todo) (err error) {
	if err := config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {
	fmt.Println("create todo1")
	if err := config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodoById(todo *Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodoById(todo *Todo, id string) (err error) {
	config.DB.Save(todo)
	return nil
}

func DeleteTodoById(todo *Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
