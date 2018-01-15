package views

import (
	"net/http"
	"log"
	"strconv"
	"github.com/izayacity/LinkedIn/sessions"
	"github.com/gorilla/mux"
	"github.com/izayacity/LinkedIn/db"
	"github.com/izayacity/LinkedIn/types"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func printProfile(skills types.Skills) {
	log.Print("Skills size: ", len(skills))
	for _, skill := range skills {
		log.Print("Skill ID: ", skill.Id)
		log.Print("Skill name: ", skill.Name)
		log.Print("Skill count: ", skill.Count)
		log.Print("Skill senders: ")
		for _, sender := range skill.Sender {
			log.Print(sender)
		}
	}
}

func GetUserName(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	token, _:= jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	username := fmt.Sprint(claims["username"])
	return username
}

func ShowMyProfile(w http.ResponseWriter, r *http.Request) {
	username := GetUserName(r)
	userId := strconv.Itoa(UserIdMap[username])
	skills := db.GetSkills(userId)
	printProfile(skills)

	log.Print("ShowMyProfile: username: ", username, " and userId: ", userId)
	w.WriteHeader(200)
}

// Abandoned
func ShowMyProfileBySession(w http.ResponseWriter, r *http.Request) {
	userId, username := sessions.GetCurrentUser(r)
	if userId == -1 {
		log.Print("Empty userId in ShowMyProfile")
		w.WriteHeader(401)
		return
	}
	uid := strconv.Itoa(userId)
	skills := db.GetSkills(uid)
	printProfile(skills)

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
	skills := db.GetSkills(uid)
	printProfile(skills)
	log.Print("ShowUserProfile: user ID is ", uid)
	w.WriteHeader(200)
}

func AddSkill(w http.ResponseWriter, r *http.Request) {
	username := GetUserName(r)
	skillName := "Golang"
	err := db.AddSkill(UserIdMap[username], skillName)
	checkErr(err, w)
	log.Print("Added skill(name) ", skillName, " to user ", username)
	w.WriteHeader(201)
}

// Abandoned
func AddSkillBySession(w http.ResponseWriter, r *http.Request) {
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
	username := GetUserName(r)
	userId := UserIdMap[username]

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
