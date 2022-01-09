package main

import (
	"go-todo-app/httpd/handler"
	"go-todo-app/service"

	"github.com/gin-gonic/gin"
)

func main() {

	todos := service.GetDataFromFile()

	router := gin.Default()

	router.GET("/todos", handler.TodoGet(todos))
	router.POST("/todos", handler.TodoPost(todos))

	router.Run()

}
