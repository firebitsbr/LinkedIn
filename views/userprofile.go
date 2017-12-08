package views

import (
	"net/http"
	"github.com/izayacity/LinkedIn/sessions"
	"log"
)

//ShowAllTasksFunc is used to handle the "/" URL which is the default ons
func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		username := sessions.GetCurrentUserName(r)
		log.Print("Current user ", username)
	}
}
