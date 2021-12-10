package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magicxiaobao/go-todo/model"
	"net/http"
	"strconv"
)

func ListTodos(ctx *gin.Context) {
	var todos []model.Todo
	err := model.GetAllTodos(&todos)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, todos)
	}
}

func CreateTodo(ctx *gin.Context) {

	var todo model.Todo
	ctx.BindJSON(&todo)
	marshal, jsonError := json.Marshal(todo)
	if jsonError != nil {
		fmt.Println(jsonError)
	}
	fmt.Println("CreateTodo,param" + string(marshal))
	err := model.CreateTodo(&todo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func GetTodoById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var todo model.Todo
	err := model.GetTodoById(&todo, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func PutTodoById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var todo model.Todo
	err := model.GetTodoById(&todo, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	ctx.BindJSON(&todo)
	err = model.UpdateTodoById(&todo, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, todo)
	}

}

func DeleteTodoById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var todo model.Todo
	err := model.GetTodoById(&todo, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	err = model.DeleteTodoById(&todo, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id:" + strconv.Itoa(int(todo.ID)): "deleted"})
	}
}


