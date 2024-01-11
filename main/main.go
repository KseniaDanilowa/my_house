package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"my_house/houseInfo/family"
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

	//InsertInFamily(family.CreateFamily()[1], conn)
	family.ShowFamily(conn)

}
