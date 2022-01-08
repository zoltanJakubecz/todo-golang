package main

import (
	"go-todo-app/httpd/handler"
	"go-todo-app/platform/todo"

	"github.com/gin-gonic/gin"
)

func main() {

	todos := todo.New()

	router := gin.Default()

	router.GET("/todos", handler.TodoGet(todos))
	router.POST("/todos", handler.TodoPost(todos))

	router.Run()

}
