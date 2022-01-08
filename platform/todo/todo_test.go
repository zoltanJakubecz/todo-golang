package todo

import "testing"

func TestAdd(t *testing.T){
	todos := New()
	todos.Add(Item{})
	if len(todos.Items) == 0 {
		t.Errorf("Item was not added")
	}
}
