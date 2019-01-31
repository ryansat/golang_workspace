package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type testTable struct {
	FirstName string
	LastName  string
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM testtable")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag testTable
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.FirstName, &tag.LastName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.FirstName)
	}

}
