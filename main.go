package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/izayacity/LinkedIn/types"
	"github.com/astaxie/beego/orm"
)

func init() {
	// register model
	orm.RegisterModel(new(types.User))
	// set default database
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:8889)/linkedin?charset=utf8")
	checkErr(err)
}

func main() {
	// bind the port of the http server
	PORT := ":5000"
	fmt.Println("Running server on "+ PORT)
	http.HandleFunc("/", showUsers)

	err := http.ListenAndServe(PORT, nil)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("ERR: %v\n", err)
		panic(err)
	}
}

func printErrorID(id int64, err error) {
	if err != nil {
		fmt.Println("ID: %d, ERR: %v", id, err)
	}
}

// test function connecting to the database and showing all the users using Orm pacakge
func showUsers (w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()
	user := types.User{Username:"Izaya", Password:"222", Email:"izayacity@gmail.com"}
	fmt.Println("showusers")
	// insert
	//id, err := o.Insert(&user)
	//printErrorID(id, err)

	// update
	//user.Username = "Izayacity"
	//id, err = o.Update(&user)
	//printErrorID(id, err)

	// query
	user = types.User{Username:"Izayacity"}
	err := o.Read(&user, "username")

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(user.Id, user.Username, user.Email)
	}
}

// test function connecting to the database and showing all the users
func dbTest (w http.ResponseWriter, r *http.Request) {
	// db connection
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/linkedin?charset=utf8")
	checkErr(err)
	defer db.Close()

	// a test query
	getUserSQL := "SELECT * FROM  User"
	rows, err := db.Query(getUserSQL)
	checkErr(err)
	defer rows.Close()

	var user types.User

	for rows.Next() {
		user = types.User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Status)
		checkErr(err)

		fmt.Println(user.Id)
		fmt.Println(user.Username)
		fmt.Println(user.Password)
		fmt.Println(user.Email)
		fmt.Println(user.Status)
	}
}
