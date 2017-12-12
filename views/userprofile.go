package views

import (
	"net/http"
	"github.com/izayacity/LinkedIn/sessions"
	"log"
	"github.com/gorilla/mux"
	"github.com/izayacity/LinkedIn/db"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Print("Internal Server ERROR: ", err)
		panic(err)
	}
}

// TODO skills
func ShowMyProfile(w http.ResponseWriter, r *http.Request) {
	userId, username := sessions.GetCurrentUser(r)
	if userId == -1 {
		log.Print("Empty userId in ShowMyProfile")
		w.WriteHeader(401)
		return
	}
	log.Print("ShowMyProfile: user ID: ", userId, ", Username: ", username)
	w.WriteHeader(200)
}

// TODO skills
func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	if &uid == nil {
		log.Print("Empty uid in ShowUserProfile")
		w.WriteHeader(400)
		return
	}
	log.Print("ShowUserProfile: user ID is ", uid)
	w.WriteHeader(200)
}

func AddSkill(w http.ResponseWriter, r *http.Request) {
	userId, _ := sessions.GetCurrentUser(r)
	if userId == -1 {
		log.Print("Empty userId in AddSkill")
		return
	}
	skillName := "Golang"
	err := db.AddSkill(userId, skillName)
	checkErr(err)
	log.Print("Added skill(name) ", skillName, " to user ", userId)
}

func RemoveSkill(w http.ResponseWriter, r *http.Request) {
	userId, _ := sessions.GetCurrentUser(r)
	if userId == -1 {
		log.Print("Empty userId in RemoveSkill")
		return
	}
	vars := mux.Vars(r)
	skillId, skillErr := strconv.Atoi(vars["sid"])

	if skillErr != nil || &skillId == nil {
		log.Print("Empty sid in RemoveSkill")
		return
	}
	err := db.RemoveSkill(userId, skillId)
	checkErr(err)
	log.Print("Removed skill(id) ", skillId, " from user ", userId)
}
