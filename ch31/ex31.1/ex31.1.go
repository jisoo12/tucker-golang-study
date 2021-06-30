package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type Todo struct {
	ID				int			`json:"id,omitempty"`
	Name			string	`json:"name"`
	Completed	bool		`json:"completed,omitempty"`
}

var todoMap map[int]Todo
var lastID int = 0

func MakeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandlerFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
	return mux
}

type Todos []Todo

func (t Todos) Len() int {
	return len(t)
}

func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	rd.JSON(w, http.StatusOK, list)
}

func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++
	todoMap[lastID] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}