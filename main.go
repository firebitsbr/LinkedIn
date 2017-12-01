package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"
	log.Print("Running server on "+ PORT)
	http.HandleFunc("/",CompleteTaskFunc)
	http.HandleFunc("/tasks/",GetTaskFunc)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

//GetTaskFunc is used to delete a task,
func GetTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Path[len("/tasks/"):]
		w.Write([]byte("Get the task " + id))
	}
}