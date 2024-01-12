package houseParams

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type HouseParams struct {
	Width      int
	Height     int
	Length     int
	RoomsCount int
}

func CreateHouseParams() HouseParams {
	params := HouseParams{
		Width:      10,
		Height:     8,
		Length:     12,
		RoomsCount: 1,
	}
	return params
}

func InsertInHouseParams(params HouseParams, conn pgx.Conn) {
	bd, beginErr := conn.Begin(context.Background())
	if beginErr != nil {
		fmt.Println(beginErr)
		return
	}
	defer bd.Rollback(context.Background())
	_, execErr := bd.Exec(context.Background(), "insert into houseparams(width, height, length, roomscount) values ($1, $2, $3, $4)",
		params.Width, params.Height, params.Length, params.RoomsCount)
	if execErr != nil {
		fmt.Println(execErr)
		return
	}
	err2 := bd.Commit(context.Background())
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}

func ShowHouseParams(conn pgx.Conn) {
	var width, height, length, roomscount string
	rows, err := conn.Query(context.Background(), "select width, height, length, roomscount from houseparams")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	} else {
		print("\n       WIDTH      HEIGHT     LENGTH    ROOMSCOUNT\n" +
			"       ____________________________________________\n")
		for rows.Next() {
			rows.Scan(&width, &height, &length, &roomscount)
			fmt.Printf("%10s %10s %10s %10s\n", width, height, length, roomscount)
		}
	}
}
