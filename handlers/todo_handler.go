package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/AaiEa00/todo/models"
)

var todos []models.TodoItem

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// ToDoアイテムの一覧を取得する
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// ToDoアイテムを作成する
	var todo models.TodoItem
	_ = json.NewDecoder(r.Body).Decode(&todo)
	// ToDoアイテムを保存する
	todos = append(todos, todo)
	fmt.Fprintf(w, "ToDo created: %s", todo.Title)
}