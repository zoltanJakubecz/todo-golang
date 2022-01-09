package todo

import uuid "github.com/satori/go.uuid"


type Todo struct {
	Id 			uuid.UUID 	`json:"id"`
	Text 		string 	`json:"text"`
	Priority 	int 	`json:"priority"`
	Done 		bool 	`json:"done"`
}


type Repo struct{
	Todos []Todo `json:"todos"`
}


func New() *Repo {
	return &Repo{
		Todos: []Todo{},
	}
}


func (r *Repo) Add(todo Todo){
	r.Todos = append(r.Todos, todo)
}


func (r *Repo) GetAll() []Todo {
	return r.Todos
}
