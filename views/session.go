package views
import (
	"log"
	"net/http"
	"github.com/izayacity/LinkedIn/db"
	"github.com/izayacity/LinkedIn/types"
	"github.com/izayacity/LinkedIn/sessions"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"time"
	"strconv"
	"io/ioutil"
	"fmt"
)

const (
	privKeyPath = "config/app.rsa"
	pubKeyPath = "config/app.rsa.pub"
)

var SignKey []byte

// Load RSA key into SignKey
func init(){
	var err error

	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
}

// Use for server error checking, panic when error is found
func checkErr(err error, w http.ResponseWriter) {
	if err != nil {
		log.Print("Internal Server ERROR: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

// Use for JSON serialization
func JsonResponse(response interface{}, w http.ResponseWriter) {
	data, err :=  json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Used for each httpHandler to check if there is any active session
func RequiresLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isValidToken(r) && !sessions.IsLoggedIn(r) {
			log.Print("Not logged in")
			w.WriteHeader(401)
			return
		}
		handler(w, r)
	}
}

// Check if the JWT token in the request header is valid
func isValidToken(r *http.Request) bool {
	// read JWT token from the Authorization key of the request header
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkVGltZSI6MTUxNTMzNTg1NiwidXNlcm5hbWUiOiJ0ZXN0In0.3xyr8vczAD34ZTZH3eUPjRQGUpj164gUhnKfYglAji8"
	tokenString := r.Header.Get("Authorization")
	fmt.Println("tokenString: ", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SignKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Username: ", claims["username"])
		fmt.Println("Created Time: ", claims["createdTime"])
		return true
	}
	fmt.Println("Error of token: ", err)
	return false
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
	log.Print("User Logged out")
	w.WriteHeader(200)
}

// Write user data into session and cookie
func initSession(w http.ResponseWriter, r *http.Request, u types.User) {
	// Initialize session and user model
	session, err := sessions.Store.Get(r, "LoginSession")
	if err != nil {
		log.Print("Fail to retrieve the session in login")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["authenticated"] = "true"
	session.Values["username"] = u.Username
	session.Values["email"] = u.Email
	session.Values["userid"] = u.Id

	// Save the session. NOTE: has to be before any writing to the response
	err = session.Save(r, w)
	checkErr(err, w)

	// Set Cookie on client for 3 months
	expiration := time.Now().Add(90 * 24 * time.Hour)
	cookieUsername := http.Cookie{Name: "username", Value: u.Username, Expires: expiration}
	cookieId := http.Cookie{Name: "userid", Value: strconv.Itoa(u.Id), Expires: expiration}
	http.SetCookie(w, &cookieUsername)
	http.SetCookie(w, &cookieId)
}

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	log.Print("POST  /v1/login/")
	u := types.User{}

	// Decode the received request body in JSON format
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Print("Bad Request JSON")
		w.WriteHeader(400)
		return
	}
	username := u.Username
	password := u.Password

	// Verify with the information in the database
	if username == "" || password == "" || !db.ValidUser(username, password) {
		log.Print("Invalid user " + username)
		w.WriteHeader(401)
		return
	}
	log.Print("user ", username, " is authenticated")

	// Get the user information and correct the username if it's an email
	u, err = db.GetUser(username)
	if err != nil {
		w.WriteHeader(401)		// not sure
		return
	}

	// Important!
	initSession(w, r, u)
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"createdTime": time.Now().Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(SignKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error signing token: %v\n", err)
		return
	}
	log.Print("Token created: ", tokenString)

	//create a token instance using the token string
	response := types.Token{Token: tokenString}
	JsonResponse(response, w)
}

// Register handler
func Register(w http.ResponseWriter, r *http.Request) {
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
		checkErr(err, w)

		log.Print("user ", username, "'s account is created")
		w.WriteHeader(201)
		return
	}
	log.Print("Invalid username " + username + " or email " + email)
	w.WriteHeader(401)
}
