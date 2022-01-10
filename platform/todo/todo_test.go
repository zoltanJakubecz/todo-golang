package todo

import "testing"

func TestAdd(t *testing.T) {
	todos := New()
	todos.Add(Todo{})
	if len(todos.Todos) == 0 {
		t.Errorf("Item was not added")
	}
}
