package views
import (
	"log"
	"net/http"
	"github.com/izayacity/LinkedIn/db"
	"github.com/izayacity/LinkedIn/types"
	"github.com/izayacity/LinkedIn/sessions"
	"encoding/json"
	"time"
	"strconv"
)

// Used for each httpHandler to check if there is any active session
func RequiresLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sessions.IsLoggedIn(r) {
			w.WriteHeader(401)
			return
		}
		handler(w, r)
	}
}

// Invalid the session information from the cookie store
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "LoginSession")
	if err == nil {
		if session.Values["authenticated"] != "false" {
			session.Values["authenticated"] = "false"
			session.Save(r, w)
		}
	}
	w.WriteHeader(401)
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "LoginSession")
	if err != nil {
		log.Print("Fail to retrieve the session in login")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		log.Print("POST request in /v1/login/")
		u := types.User{}

		// Decode the received request body in JSON format
		err = json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Print("Bad Request JSON")
			w.WriteHeader(400)
			return
		}
		username := u.Username
		password := u.Password

		// Verify with the information in the database
		if (username != "" && password != "") && db.ValidUser(username, password) {
			log.Print("user ", username, " is authenticated")

			// Get the user information and correct the username if it's an email
			u = db.GetUser(username)
			session.Values["authenticated"] = "true"
			session.Values["username"] = u.Username
			session.Values["email"] = u.Email
			session.Values["userid"] = u.Id

			// Save the session. NOTE: has to be before any writing to the response
			err = session.Save(r, w)
			if err != nil {
				log.Print("Fail to save the session", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			// Set Cookie on client for 3 months
			expiration := time.Now().Add(90 * 24 * time.Hour)
			cookieUsername := http.Cookie{Name: "username", Value: u.Username, Expires: expiration}
			cookieId := http.Cookie{Name: "userid", Value: strconv.Itoa(u.Id), Expires: expiration}
			http.SetCookie(w, &cookieUsername)
			http.SetCookie(w, &cookieId)

			w.WriteHeader(200)
			return
		}
		log.Print("Invalid user " + username)
		w.WriteHeader(401)
	default:
		log.Print("Bad Request Method")
		w.WriteHeader(400)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		log.Print("POST request in /v1/register/")
		u := types.User{}
		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			log.Print("Bad Request")
			w.WriteHeader(400)
			return
		}
		username := u.Username
		email := u.Email
		password := u.Password

		// Check if the username or email already exists in the database
		if (username != "" && email != "" && password != "") && db.ValidUsername(username) && db.ValidEmail(email) {
			err = db.CreateAccount(username, email, password)
			if err != nil {
				log.Print("Fail to create the user account")
				return
			}
			log.Print("user ", username, "'s account is created")
			w.WriteHeader(200)
			return
		}
		log.Print("Invalid username " + username + " or email " + email)
		w.WriteHeader(401)
	default:
		log.Print("Bad Request")
		w.WriteHeader(400)
	}
}
