package views

import (
	"net/http"
	"github.com/izayacity/LinkedIn/sessions"
	"github.com/izayacity/LinkedIn/db"
	"log"
	"github.com/gorilla/mux"
)

func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		username := sessions.GetCurrentUserName(r)
		user := db.GetUser(username)
		log.Print("User ID: ", user.Id, ", User Email: ", user.Email, ", Username: ", user.Username)
	default:
		w.WriteHeader(400)
	}
}

func ShowUserProfileById(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		vars := mux.Vars(r)
		log.Print("User ID is ", vars["uid"])
	default:
		w.WriteHeader(400)
	}
}