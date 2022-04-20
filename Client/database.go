package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/alexeyco/simpletable"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qustavo/dotsql"
)

var db *sql.DB
var dotQueries *dotsql.DotSql

func initializeDatabase() {

	// Load the database file into memory
	os.Create("../Database/nito.db")
	db, _ = sql.Open("sqlite3", "../Database/nito.db")
	dotQueries, _ = dotsql.LoadFromFile("../Database/queries.sql")

	// Test for database access
	err := db.Ping()
	if err != nil {
		fmt.Println("[ERROR] Cannot connect to the database")
	}

	// Create necessary database tables
	_, err = dotQueries.Exec(db, "initialize-database")
	if err != nil {
		log.Fatal(err)
	}

	// DEBUG: Load Sample Data
	_, err = dotQueries.Exec(db, "sample-data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[SUCCESS] Database Initialized")
}

// ============================= Listeners =============================

func getListeners() {

	// Define return values
	var id, protocol, port string

	// Run the query
	rows, err := dotQueries.Query(db, "get-all-listeners")
	if err != nil {
		panic(err)
	}

	// Define the table and its header
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Protocol"},
			{Align: simpletable.AlignCenter, Text: "Port"},
		},
	}

	// Scan for return values and populate the table
	for rows.Next() {
		err = rows.Scan(&id, &protocol, &port)
		if err != nil {
			panic(err)
		}

		// Add row to a cell
		cell := []*simpletable.Cell{
			{Text: id},
			{Text: protocol},
			{Text: port},
		}

		table.Body.Cells = append(table.Body.Cells, cell)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
