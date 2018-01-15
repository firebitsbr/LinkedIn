//The MIT License (MIT)
//Copyright (c) 2015 Suraj Patil
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package sessions

import (
"net/http"
"github.com/gorilla/sessions"
	"log"
)

// Store the cookie store which is going to store session data in the cookie
var Store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	Store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
}

// Check if the user has an active session and return True
func IsLoggedIn(r *http.Request) bool {
	session, err := Store.Get(r, "LoginSession")
	if err == nil && (session.Values["authenticated"] == "true") {
		return true
	}
	return false
}

// Returns the username of the logged in user
func GetCurrentUser(r *http.Request) (int, string) {
	session, err := Store.Get(r, "LoginSession")
	if err == nil && session.Values["userid"] != nil && session.Values["username"] != nil {
		log.Print(session.Values["userid"])
		log.Print(session.Values["username"])
		return session.Values["userid"].(int), session.Values["username"].(string)
	}
	return -1, ""
}
