package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var defaultDBLocation string = "../db/main.db"

func bootstrapDB() {
	var dbFileExists bool = createDB()
	if dbFileExists {
		var db = connectDb(defaultDBLocation)
		formatDB(db)
		testDB(db)
	} else {
		createDB()
	}
}


func createDB() bool{
	if fileExists(defaultDBLocation) {
		return true
	} else {
		os.Create(defaultDBLocation)
		fmt.Println("DB Created")
		return false
	}
}

func fileExists(filename string) bool {
	test, err:= os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	} else {
		return !test.IsDir()
	}
}

func connectDb(dbFile string) *sql.DB{
	db,_:= sql.Open("sqlite3", dbFile)
	fmt.Println("Connected to DB")
	return db
}

func formatDB(db *sql.DB) *sql.DB{
	format,_ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, passwordHash TEXT)")
	format.Exec()
	fmt.Println("DB Formatted")
	return db
}

func testDB(db *sql.DB) {
		insert, _:= db.Prepare("INSERT INTO users (username, passwordHash) VALUES (?,?)")
	insert.Exec("another", "another")

	rows,_:= db.Query("SELECT id, username, passwordHash from users")
	var id int
	var username string
	var passwordHash string

	for rows.Next() {
		rows.Scan(&id, &username, &passwordHash)
		fmt.Println(id,username,passwordHash)
	}
}