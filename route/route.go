package route

import (
	"github.com/gin-gonic/gin"
	"github.com/magicxiaobao/go-todo/controller"
)

func SetupRoute() *gin.Engine {
	engine := gin.Default()
	groupV1 := engine.Group("v1")
	groupV1.GET("todos", controller.ListTodos)
	groupV1.POST("todo", controller.CreateTodo)
	groupV1.GET("todo/:id", controller.GetTodoById)
	groupV1.PUT("todo/:id", controller.PutTodoById)
	groupV1.DELETE("todo/:id", controller.DeleteTodoById)
	return engine
}
