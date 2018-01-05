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

	// Use Gorilla mux router to handle parameterized routing with methods
	r := mux.NewRouter()
	r.HandleFunc("/v1/login", views.Login).Methods("POST")
	r.HandleFunc("/v1/register", views.Register).Methods("POST")
	r.HandleFunc("/v1/logout", views.RequiresLogin(views.Logout)).Methods("GET")
	r.HandleFunc("/v1/me", views.RequiresLogin(views.ShowMyProfile)).Methods("GET")
	r.HandleFunc("/v1/users/{uid}", views.ShowUserProfile).Methods("GET")
	r.HandleFunc("/v1/me/skills", views.RequiresLogin(views.AddSkill)).Methods("POST")
	r.HandleFunc("/v1/me/skills/{sid}", views.RequiresLogin(views.RemoveSkill)).Methods("DELETE")
	r.HandleFunc("/v1/users/{uid}/skills/{sid}/endorse", views.RequiresLogin(views.Endorse)).Methods("PUT")

	// deal with CORS issue by using Gorilla handlers
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:5000", "http://evil.com/"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE","OPTIONS"})
	credentialsOk := handlers.AllowCredentials()

	// http listen and serve on the address and port
	log.Fatal(http.ListenAndServe(PORT, handlers.CORS(originsOk, headersOk, methodsOk, credentialsOk)(r)))
}
