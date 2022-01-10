package service

import (
	"encoding/json"
	"fmt"
	"go-todo-app/platform/todo"
	"os"

	//	"encoding/json"
	"io/ioutil"
)

func GetDataFromFile() *todo.Repo {
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		fmt.Println(err)
		return todo.New()
	}

	todos := todo.New()

	fmt.Println("Data file opened successful")
	json.Unmarshal(data, todos)
	return todos
}

func WriteDataToFile(todos *todo.Repo) {
	file, _ := json.MarshalIndent(todos, "", " ")
	os.WriteFile("./data/data.json", file, 0644)
}
