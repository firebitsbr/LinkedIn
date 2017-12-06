package db

import (
	"github.com/izayacity/LinkedIn/types"
	"github.com/astaxie/beego/orm"
	"log"
	"strings"
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
	user := types.User{}
	cond := orm.NewCondition()
	cond1 := cond.And("username", username).Or("email", username)
	err := db.QueryTable("User").SetCond(cond1).One(&user)

	if err == orm.ErrNoRows {
		log.Print("No result found.")
		return false
	} else if err == orm.ErrMultiRows {
		log.Print("Returned Multi Rows Not One")
		return false
	}
	//If the password matches, return true
	if strings.Compare(password, user.Password) == 0 {
		return true
	}
	//by default return false
	log.Print("Password mismatch")
	return false
}
