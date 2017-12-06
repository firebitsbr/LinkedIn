package views
import (
	"log"
	"net/http"
	"github.com/izayacity/LinkedIn/db"
	"github.com/izayacity/LinkedIn/types"
	"encoding/json"
)

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		log.Print("POST request in /v1/login/")
		u := types.User{}
		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			log.Print("Bad Request")
			w.WriteHeader(400)
			return
		}
		username := u.Username
		password := u.Password

		if (username != "" && password != "") && db.ValidUser(username, password) {
			log.Print("user ", username, " is authenticated")
			w.WriteHeader(200)
			return
		}
		log.Print("Invalid user " + username)
		w.WriteHeader(404)
	default:
		log.Print("Bad Request")
		w.WriteHeader(400)
	}
}
