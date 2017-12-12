package views

import (
	"net/http"
	"log"
	"github.com/izayacity/LinkedIn/sessions"
	"github.com/izayacity/LinkedIn/db"
	"github.com/gorilla/mux"
	"strconv"
)

func Endorse(w http.ResponseWriter, r *http.Request) {
	userId, _ := sessions.GetCurrentUser(r)
	if userId == -1 {
		log.Print("Empty current userId in Endorse")
		w.WriteHeader(401)
		return
	}
	vars := mux.Vars(r)
	uid, _ := strconv.Atoi(vars["uid"])
	sid, _ := strconv.Atoi(vars["sid"])

	if &uid == nil || &sid == nil {
		log.Print("Empty uid or sid in Endorse")
		w.WriteHeader(400)
		return
	}
	// update endorse in the database
	if !db.IsEndorsed(sid, userId, uid) {
		err := db.Endorse(sid, userId, uid)
		checkErr(err, w)
		db.UpdateCount(sid)
	}
	log.Print("User ", userId, " endorsed skill ", sid, " to user ", uid)
	w.WriteHeader(204)
}
