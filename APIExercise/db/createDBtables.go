package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	DBname := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "polina", "password", "bookstore")
	db, err := sql.Open("postgres", DBname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE authors(
		id SERIAL PRIMARY KEY ,
		name VARCHAR ( 50 ) UNIQUE NOT NULL,
		biography VARCHAR ( 500 ) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE categories(
		id SERIAL PRIMARY KEY ,
		name VARCHAR ( 50 ) UNIQUE NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE books(
		id SERIAL PRIMARY KEY ,
		title VARCHAR ( 50 ) UNIQUE NOT NULL,
		a_id INTEGER NOT NULL references authors(id),
		c_id INTEGER NOT NULL references categories(id),
		price FLOAT

	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO categories(name) VALUES ($1)", "Fantasy")
	if err != nil {
		log.Fatal(err)
	}

}
