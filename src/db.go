package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var defaultDBLocation string = "../db/main.db"
var globalDB *sql.DB = nil

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
	globalDB = db
	return db
}
//ADD PUBLIC USERKEYS ASSOCIATED WITH A USERNAME 

func formatDB(db *sql.DB) *sql.DB{
	format,_:= db.Prepare("CREATE TABLE IF NOT EXISTS grps (id INTEGER PRIMARY KEY, grp TEXT, key TEXT)")
	format.Exec()
	format,_= db.Prepare("CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY, message TEXT, grp TEXT, FOREIGN KEY(grp) REFERENCES grps(grp))")
	format.Exec()
	fmt.Println("Data table created")
	fmt.Println("DB Formatted")
	return db
}

func testDB(db *sql.DB) {
	insert,_:=db.Prepare("INSERT INTO grps (grp, key) VALUES (?,?)")
	insert.Exec("testing", "1")
	insert,_= db.Prepare("INSERT INTO messages (message, grp) VALUES (?,?)")
	insert.Exec("message","testgroup")
	messages,_:= db.Query("SELECT id, grp, key FROM grps")
	var id int
	var grp string
	var key string

	for messages.Next() {
		messages.Scan(&id,&grp,&key)
		fmt.Println(id,grp,key)
	}

	fmt.Println(getGrpKey("testing"))


}

func getGrpsMessages(grp string, key string) {
	queryKey,_:= globalDB.Prepare("SELECT key FROM messages WHERE grp=?")
	queryKey.Exec(grp)
}

func getGrpKey(grp string) string {
	var key string = "22"
	query,_:=globalDB.Query("SELECT key FROM grps WHERE grp=$1", grp)
	for query.Next() {
		query.Scan(&key)
	}
	return key
}