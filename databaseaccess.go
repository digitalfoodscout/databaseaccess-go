package main

import(
	 "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

)
var (
	id int
	name string
)

func main() {
	db,err := sql.Open("mysql", "user:password@tcp(localhost:port)/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM foods WHERE id <= 100 ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}