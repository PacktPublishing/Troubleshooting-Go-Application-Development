package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://gopher:rdbmsftw@localhost:5432/gopher?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to make connection: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping: %s", err)
	}
	_, err = db.Exec(`
        CREATE TABLE contacts (
            id SERIAL PRIMARY KEY,
            first VARCHAR(100),
            last VARCHAR(100)
        )
    `)
	if err != nil {
		log.Fatalf("Failed to create table: %s\n", err)
	}

	type person struct {
		first string
		last  string
	}
	for _, p := range []person{
		person{first: "Rob", last: "Pike"},
		person{first: "Ken", last: "Thompson"},
		person{first: "Robert", last: "Griesemer"},
	} {

		_, err = db.Exec(`
            INSERT INTO contacts 
                (first, last)
            VALUES
                ($1, $2)
        `, p.first, p.last)
		if err != nil {
			log.Fatalf("Failed to insert record: %s\n", err)
		}
	}

	fmt.Println("done")

}
