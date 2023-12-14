package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {
	//printHouse.PrintHouseInfo()

	urlExample := "postgres://my_house_project:123@localhost:5436/house_db"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	//var name string
	//var id int64
	//err = conn.QueryRow(context.Background(), "select id, name from users").Scan(&id, &name)
	//rows, err := conn.Query(context.Background(), "select id, name from users")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//for rows.Next() {
	//	rows.Scan(&id, &name)
	//	fmt.Println("id = ", id, "\nname = ", name)
	//}
	nameInsert := "Ivan"

	tx, err := conn.Begin(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	defer tx.Rollback(context.Background())

	tx.Exec(context.Background(), "insert into users(name) values ($1)", nameInsert)

	tx.Commit(context.Background())
}
