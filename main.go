package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/izayacity/LinkedIn/views"
)

func main() {
	// bind the port of the http server
	PORT := ":5000"
	log.Print("Running server on "+ PORT)
	r := mux.NewRouter()
	r.HandleFunc("/v1/login", views.Login)
	r.HandleFunc("/v1/register", views.Register)
	r.HandleFunc("/v1/me", views.RequiresLogin(views.ShowUserProfile))
	r.HandleFunc("/v1/users/{uid}", views.ShowUserProfileById)

	// dealing with CORS issue with gorilla handlers
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE","OPTIONS"})

	log.Fatal(http.ListenAndServe(PORT, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
