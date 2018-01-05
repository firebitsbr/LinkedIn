package db

import (
	"github.com/izayacity/LinkedIn/types"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"errors"
)

var db orm.Ormer

func init() {
	// register model
	orm.RegisterModel(new(types.User))
	// set default database
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:8889)/linkedin")
	checkErr(err)
	db = orm.NewOrm()
}

func checkErr(err error) {
	if err != nil {
		log.Print("ERROR: ", err)
		panic(err)
	}
}

// ValidUser will check if the user exists in db; If exists,
// then check if the username password combination is valid
func ValidUser(username, password string) bool {
	if username == "" || password == "" {
		log.Print("Empty field")
		return false
	}
	user, err := GetUser(username)
	if err != nil {
		return false
	}
	//If the password matches, return true
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Print("ERROR: ", err)
		return false
	}
	return true
}

// return true when there is not record in the database, so the account could be registered
func ValidEmail(email string) bool {
	if email == "" {
		log.Print("Empty email")
		return false
	}
	user := types.User{}
	err := db.QueryTable("user").Filter("email", email).One(&user)

	if err == orm.ErrNoRows {
		return true
	}
	return false
}

// return true when there is not record in the database, so the account could be registered
func ValidUsername(username string) bool {
	if username == "" {
		log.Print("Empty username")
		return false
	}
	user := types.User{}
	err := db.QueryTable("user").Filter("username", username).One(&user)

	if err == orm.ErrNoRows {
		return true
	}
	return false
}

func CreateAccount(username, email, password string) error {
	if username == "" || email == "" || password == "" {
		return errors.New("empty field")
	}
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	hash := string(hashBytes)
	user := types.User{Username:username, Email:email, Password:hash}
	_, err = db.Insert(&user)
	return err
}

func GetUser(username string) (types.User, error) {
	user := types.User{}
	if username == "" {
		return user, errors.New("empty username")
	}
	cond := orm.NewCondition()
	cond1 := cond.And("username", username).Or("email", username)

	err := db.QueryTable("user").SetCond(cond1).One(&user)
	if err == orm.ErrNoRows {
		log.Print("No result found.")
	} else if err == orm.ErrMultiRows {
		log.Print("Returned Multi Rows Not One")
	}
	return user, err
}