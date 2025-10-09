package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var Todos = make(map[int]Todo)
var NextID = 1
var Mu sync.Mutex

func HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(Todos); err != nil {
		log.Printf("JSONコードエラー: %v", err)
		http.Error(w, "サーバエラー", http.StatusInternalServerError)
	}
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "不正なリクエストボディです", http.StatusBadRequest)
	}

	Mu.Lock()
	newTodo.ID = 1
	Todos[NextID] = newTodo
	NextID++
	Mu.Unlock()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		HandleGet(w, r)
	case http.MethodPost:
		HandlePost(w, r)
	default:
		http.Error(w, "サポートされていないです", http.StatusMethodNotAllowed)
	}
}
