package views

import (
	"net/http"
	"log"
	"strconv"
	"github.com/izayacity/LinkedIn/sessions"
	"github.com/gorilla/mux"
	"github.com/izayacity/LinkedIn/db"
)

// TODO skills json response
func ShowMyProfile(w http.ResponseWriter, r *http.Request) {
	userId, username := sessions.GetCurrentUser(r)
	if userId == -1 {
		log.Print("Empty userId in ShowMyProfile")
		w.WriteHeader(401)
		return
	}
	uid := strconv.Itoa(userId)
	db.GetSkills(uid)
	log.Print("ShowMyProfile: user ID: ", userId, ", Username: ", username)
	w.WriteHeader(200)
}

// TODO skills json response
func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	if &uid == nil {
		log.Print("Empty uid in ShowUserProfile")
		w.WriteHeader(400)
		return
	}
	db.GetSkills(uid)
	log.Print("ShowUserProfile: user ID is ", uid)
	w.WriteHeader(200)
}

// TODO skill object response
func AddSkill(w http.ResponseWriter, r *http.Request) {
	userId, _ := sessions.GetCurrentUser(r)
	if userId == -1 {
		log.Print("Empty userId in AddSkill")
		w.WriteHeader(401)
		return
	}
	skillName := "Golang"
	err := db.AddSkill(userId, skillName)
	checkErr(err, w)
	log.Print("Added skill(name) ", skillName, " to user ", userId)
	w.WriteHeader(201)
}

func RemoveSkill(w http.ResponseWriter, r *http.Request) {
	userId, _ := sessions.GetCurrentUser(r)
	if userId == -1 {
		log.Print("Empty userId in RemoveSkill")
		w.WriteHeader(401)
		return
	}
	vars := mux.Vars(r)
	skillId, skillErr := strconv.Atoi(vars["sid"])

	if skillErr != nil || &skillId == nil {
		log.Print("Empty sid in RemoveSkill")
		w.WriteHeader(400)
		return
	}
	err := db.RemoveSkill(userId, skillId)
	checkErr(err, w)
	log.Print("Removed skill(id) ", skillId, " from user ", userId)
	w.WriteHeader(204)
}
