package views
import (
	"log"
	"net/http"
	"github.com/izayacity/LinkedIn/db"
	"github.com/izayacity/LinkedIn/types"
	"github.com/izayacity/LinkedIn/sessions"
	"encoding/json"
)

//RequiresLogin is a middleware which will be used for each httpHandler to check
// if there is any active session
func RequiresLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sessions.IsLoggedIn(r) {
			http.Redirect(w, r, "/v1/login/", 302)
			return
		}
		handler(w, r)
	}
}

//LogoutFunc Implements the logout functionality. WIll delete the session information from the cookie store
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "LoginSession")
	if err == nil {
		if session.Values["authenticated"] != "false" {
			session.Values["authenticated"] = "false"
			session.Save(r, w)
		}
	}
	http.Redirect(w, r, "/v1/login/", 302) //redirect to login irrespective of error or not
}

//LoginFunc implements the login functionality,
// will add a cookie to the cookie store for managing authentication
func Login(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "LoginSession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		log.Print("POST request in /v1/login/")
		u := types.User{}

		err = json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Print("Bad Request JSON")
			w.WriteHeader(400)
			return
		}
		username := u.Username
		password := u.Password

		if (username != "" && password != "") && db.ValidUser(username, password) {
			session.Values["authenticated"] = "true"
			session.Values["username"] = username
			err = session.Save(r, w)
			if err != nil {
				log.Print("Fail to save the session", err)
			}

			log.Print("user ", username, " is authenticated")
			http.Redirect(w, r, "/", 302)
			return
		}
		log.Print("Invalid user " + username)
		http.Redirect(w, r, "/v1/login/", 401)
	default:
		log.Print("Bad Request Method")
		http.Redirect(w, r, "/v1/login/", 400)
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
		http.Redirect(w, r, "/v1/register/", http.StatusUnauthorized)
	default:
		log.Print("Bad Request")
		http.Redirect(w, r, "/v1/register/", http.StatusUnauthorized)
	}
}
