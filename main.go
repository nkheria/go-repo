package main

import (
	"database/sql"
	"fmt"
	"log"

	- "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	//connect to database

	conn, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=password dbname=bookings sslmode=disable")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}

	defer conn.Close()

	log.Println("Connected to database!")

	//test the connection

	err = conn.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to ping: %v\n", err))
	}

	log.Println("Pinged the database!")
	//get the rows from the table

	err = getAllRows()
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to get rows: %v\n", err))

	//insert a row

	//get rows from the table again

	//update a row

	//get rows from the table again

	//get one row by id

	//delete one row

	//get rows again

}

func getAllRows(conn *sql.DB) errot{
	rows, err := conn.Query("SELECT id,first_name,last_name FROM users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is: ", id, firstName, lastName)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error scanning rows",err)
	}

	fmt.Println("-----------------")
	
	return nil

}
