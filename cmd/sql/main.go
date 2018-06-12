package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Customer struct {
	Name   sql.NullString `sql:"type:text"`
	Apikey sql.NullString `sql:"type:text"`
}

var db *sql.DB

func main() {

	var err error

	connStr := "postgres://app1@micropostgres:wz94GaAbhMBQsmalQpan@micropostgres.postgres.database.azure.com:5432/Customer"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT "name", "apikey" FROM public."Customer"`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.Name, &customer.Apikey)

		// var name sql.NullString
		// var apikey sql.NullString
		// err = rows.Scan(&name, &apikey)
		if err != nil {
			log.Fatal(err)
		}

		println(customer.Name.String + " " + customer.Apikey.String)
		//fmt.Println(name.String + " " + apikey.String)

	}

}
