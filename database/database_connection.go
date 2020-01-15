package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Trying to connect mysql database")

	db, err := sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/testDb")

	if err!=nil{
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO USER VALUES('agrawalvishal1698@gmail.com',1234)")

	if err!=nil{
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Successfully inserted into user table")

	users, err := db.Query("select * from user")

	if err!=nil{
		panic(err.Error())
	}

	fmt.Println(users)
}
