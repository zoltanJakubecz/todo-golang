package handler

import (
	"go-todo-app/platform/todo"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type todoPostRequest struct {
	Text string `json:"text"`
	Priority int `json:"priority"`
}

func TodoPost(todos *todo.Repo) gin.HandlerFunc {
	return func(c *gin.Context){
		requestBody := todoPostRequest{}
		c.Bind(&requestBody)
		newTodo := todo.Todo{
			Id: uuid.NewV4(),
			Text: requestBody.Text,
			Priority: requestBody.Priority,
			Done: false,
		}
		todos.Add(newTodo)

		c.Status(http.StatusNoContent)

	}
}
