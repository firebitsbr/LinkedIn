package main

import (
	"log"
	"net/http"
	"database/sql"
	"github.com/izayacity/LinkedIn/types"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	PORT := ":3000"
	log.Print("Running server on "+ PORT)
	http.HandleFunc("/", showUsers)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func showUsers (w http.ResponseWriter, r *http.Request) {
	// data models
	var user types.User

	// db connection
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/linkedin?charset=utf8")
	checkErr(err)

	// a test query
	getUserSQL := "SELECT * FROM  User"
	rows, err := db.Query(getUserSQL)
	checkErr(err)
	defer rows.Close()

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