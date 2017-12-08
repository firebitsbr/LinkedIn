package views

import (
	"net/http"
	"github.com/izayacity/LinkedIn/sessions"
	"github.com/izayacity/LinkedIn/db"
	"log"
)

//ShowAllTasksFunc is used to handle the "/" URL which is the default ons
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
