package connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func MakeConn(ctx context.Context) pgx.Conn {
	url := "postgres://my_house_project:123@localhost:5436/house_db"
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return *conn
}
