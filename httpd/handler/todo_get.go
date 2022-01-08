package handler

import (
	"go-todo-app/platform/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoGet(todos *todo.Repo) gin.HandlerFunc {
	return func(c *gin.Context){
		result := todos.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
