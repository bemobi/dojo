package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("insert into tasks (name) values ($1) returning id;", "Renan")
	if err != nil {
		log.Fatalf("unable to run db.Exec: %s", err)
	}

	rows, err := db.Query("select id, name from tasks;")
	if err != nil {
		log.Fatalf("unable to select: %s", err)
	}

	for rows.Next() {
		task := struct {
			ID   int64
			Name string
		}{}
		rows.Scan(
			&task.ID,
			&task.Name,
		)
		log.Printf("Task found: %+v\n", task)
	}
}
