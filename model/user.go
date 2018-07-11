package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	//If we don't get any errors but somehow still don't get a db connection
	//we exit as well

	if db == nil {
		panic("db nil")
	}

	return db
}

func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS user_info(
		id			INTEGER PRIMARY KEY AUTOINCREMENT,
		permid		INTEGER NOT NULL,
		name 		VARCHAR NOT NULL,
		password 	VARCHAR NOT NULL,
		phone		VARCHAR,
		create_time	DATETIME,
		login_time	DATETIME,
		last_login	DATETIME,
		login_times	INTEGER
		
	);
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above

	if err != nil {
		panic(err)
	}
}

//Insert Admin user
func InsertUser(db *sql.DB, permid int, name, password string) {

	userId, _ := QueryUser(db, name)
	if userId == 1 {
		return
	}

	stmt, err := db.Prepare("INSERT INTO user_info(permid, name, password, create_time) values (?,?,?,?)")
	checkErr(err)

	//res, err := stmt.Exec("0", "1")
	res, err := stmt.Exec(permid, name, password, time.Now().Format("2006-01-02 15:04:05"))
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)
}
func UpdatePassword(db *sql.DB, name, password string) {
	//	sql := `
	//	UPDATE access_info SET count=count+1
	//	`
	//	_, err := db.Exec(sql)

	//	if err != nil {
	//		panic(err)
	//	}
	stmt, err := db.Prepare("UPDATE user_info SET password=? where name=?")
	checkErr(err)

	res, err := stmt.Exec(password, name)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)
}

func DeleteUser(db *sql.DB, name string) {
	stmt, err := db.Prepare("DELETE from user_info where name=?")
	checkErr(err)

	res, err := stmt.Exec(name)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)
}
func QueryUser(db *sql.DB, name string) (int, error) {

	var err error
	//var rows *sql.Rows

	var id, permid int
	var password string

	if name != "" {
		rows := db.QueryRow("select id,permid,password from user_info where name=?", name)
		err = rows.Scan(&id, &permid, &password)
		fmt.Printf("id=%d, permid=%d, name=%s, password=%s\n",
			id, permid, name, password)
	} else {
		return -1, errors.New("No name assigned")
	}

	return id, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db := InitDB("test.db")
	Migrate(db)

	fmt.Println("Cread table done")

	QueryUser(db, "admin")
	InsertUser(db, 7, "admin", "admin")
	InsertUser(db, 7, "admintest", "admintest")
	QueryUser(db, "admin")
	UpdatePassword(db, "admin", "password")
	QueryUser(db, "admin")
	DeleteUser(db, "admintest")
	QueryUser(db, "admintest")

	db.Close()
}
