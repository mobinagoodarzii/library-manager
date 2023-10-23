package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func CheckInitialization() {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	//making schema
	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS library")
	if err != nil {
		log.Fatal(err)
	}

	//Use library
	_, err = db.Exec("USE library")
	if err != nil {
		log.Fatal(err)
	}

	//Checking the existence of Books and Members table in library
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS Books (
        bookId INT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        bookName VARCHAR(255) NOT NULL,
        author VARCHAR(255) NOT NULL,
        pages INT,
        price FLOAT,
        numberOfBook INT,
        dateRegistration TIMESTAMP,
        status VARCHAR(255)
    )
`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS Members (
        joinId INT PRIMARY KEY,
        firstname VARCHAR(255) NOT NULL,
        lastname VARCHAR(255) NOT NULL,
        age INT,
        dateRegistration TIMESTAMP,
        status VARCHAR(255)
    )
`)
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectToDB func returns a db connection
func ConnectToDB() (*sql.DB, error) {

	//Connect to the database
	passWord := os.Getenv("MySQL_PASSWORD")
	db, err := sql.Open("mysql", "root:"+passWord+"@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}

	//Call Ping to verify that the data source name is valid.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
