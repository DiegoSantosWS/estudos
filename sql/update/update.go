package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ? ")
	stmt.Exec("Uoxiton Istive", 1)
	stmt.Exec("Diego dos santos", 2)

	// delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ? ")
	stmt2.Exec(3)
}
