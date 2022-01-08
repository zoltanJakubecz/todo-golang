package todo

import uuid "github.com/satori/go.uuid"


type Item struct {
	Id 			uuid.UUID 	`json:"id"`
	Text 		string 	`json:"text"`
	Priority 	int 	`json:"priority"`
	Done 		bool 	`json:"done"`
}


type Repo struct{
	Items []Item
}


func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}


func (r *Repo) Add(item Item){
	r.Items = append(r.Items, item)
}


func (r *Repo) GetAll() []Item {
	return r.Items
}
