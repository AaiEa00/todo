package main

import (
	"net/http"
	"github.com/AaiEa00/todo/handlers"
)

func main() {
	// ハンドラを登録する
	http.HandleFunc("/todos", handlers.GetTodos)
	http.HandleFunc("/todos/create", handlers.CreateTodo)
	// Webサーバを起動する
	http.ListenAndServe(":8080", nil)
}
