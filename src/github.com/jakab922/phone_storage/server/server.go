package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jakab922/phone_storage/utils"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strings"
)

func createTableIfNotExists(db *sql.DB) error {
	createString := `CREATE TABLE IF NOT EXISTS phone_data (
		phone_number varchar(20) PRIMARY KEY,
		name varchar(30) NOT NULL
	);`

	_, err := db.Query(createString)
	log.Printf("The error after table creation: %v", err)
	return err
}

func insertData(db *sql.DB, data []utils.PhoneData) error {
	queryString := "INSERT INTO phone_data(name, phone_number) VALUES"
	tmp := make([]string, 0)

	for _, el := range data {
		// I know this is not SQL injection safe, but this is the best I can get now
		tmp = append(tmp, fmt.Sprintf("('%v', '%v')", el.Name, el.PhoneNumber))
	}

	queryString = fmt.Sprintf("%v %v;", queryString, strings.Join(tmp, ", "))
	log.Printf("queryString: %v\n", queryString)
	_, err := db.Query(queryString)
	log.Printf("The error after inserting: %v", err)
	return err
}

func main() {
	dbAddress := os.Getenv("DB_ADDRESS")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSchema := os.Getenv("DB_SCHEMA")

	connectionString := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbUser, dbPassword, dbAddress, dbSchema)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only the POST method is allowed at this endpoint", 400)
			return
		}

		if r.Body == nil {
			http.Error(w, "No data received", 400)
			return
		}

		var data []utils.PhoneData
		err := json.NewDecoder(r.Body).Decode(&data)
		log.Printf("Received data from a client: %v", data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = createTableIfNotExists(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		err = insertData(db, data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})

	serverPort := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil))
}
