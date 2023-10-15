package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/olekukonko/tablewriter"
	_ "github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type Book struct {
	ID     int64
	Author string
	Title  string
	Price  int64
}

func main() {
	psqlCfg := fmt.Sprintf("host = %s port = %d user = %s "+
		"password = %s dbname = %s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlCfg)

	db, err := sql.Open("postgres", psqlCfg)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, _ := db.Query("SELECT * FROM books")

	var books [][]string

	for rows.Next() {
		var bk Book

		if err := rows.Scan(&bk.ID, &bk.Author, &bk.Title, &bk.Price); err != nil {
			panic(err)
		}

		bookSlice := []string{strconv.FormatInt(bk.ID, 10), bk.Author, bk.Title, strconv.FormatInt(bk.Price, 10)}
		books = append(books, bookSlice)
	}

	table := tablewriter.NewWriter(os.Stdout)

	bkColumns, _ := rows.Columns()
	table.SetHeader(bkColumns)
	table.SetBorder(true)
	table.AppendBulk(books)
	table.Render()
}
