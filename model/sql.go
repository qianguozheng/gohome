package model

import (
	"database/sql"
	"fmt"

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
	CREATE TABLE IF NOT EXISTS access_info(
		id INTERGER PRIMARY KEY,
		count INTERGER
	);
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above

	if err != nil {
		panic(err)
	}
}
func InsertCount(db *sql.DB) {

	accessId, _ := QueryCount(db)
	if accessId != -1 {

		return
	}

	stmt, err := db.Prepare("INSERT INTO access_info(id, count) values (?,?)")
	checkErr(err)

	//res, err := stmt.Exec("0", "1")
	res, err := stmt.Exec(0, 1)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id=", id)
}
func UpdateCount(db *sql.DB) {
	sql := `
	UPDATE access_info SET count=count+1
	`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}

func QueryCount(db *sql.DB) (int, int) {

	rows, err := db.Query("SELECT * FROM access_info")

	if err != nil {
		panic(err)
	}

	var accessNumber int = -1
	var id int = -1

	for rows.Next() {
		err = rows.Scan(&id, &accessNumber)

		if err != nil {
			panic(err)
		}

		fmt.Println(accessNumber)
	}
	rows.Close()
	return id, accessNumber
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//func main() {
//	db := InitDB("test.db")
//	Migrate(db)
//	InsertCount(db)
//	fmt.Println("Cread table done")
//	QueryCount(db)
//	UpdateCount(db)
//	UpdateCount(db)
//	QueryCount(db)
//	db.Close()
//}
