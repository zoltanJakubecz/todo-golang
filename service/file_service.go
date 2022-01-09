package service

import (
	"encoding/json"
	"fmt"
	"go-todo-app/platform/todo"

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

	fmt.Println("Data fie opened")
	json.Unmarshal(data , todos)
	fmt.Println(todos)
	return todos
}

func WriteDataToFile(todos *todo.Repo){
	file, _ := json.MarshalIndent(todos,"" ," ")
	_ = ioutil.WriteFile("data/data.json", file, 0644)
}
