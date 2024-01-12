package family

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type Family struct {
	Name   string
	Age    int
	Status string
	Room   string
}

func CreateFamily() []Family {
	familyMembers := []Family{
		{Name: "Elena", Age: 45, Status: "Mother", Room: "bedroom1"},
		{Name: "Ksenia", Age: 20, Status: "Daughter", Room: "bedroom2"},
	}
	return familyMembers
}

func InsertInFamily(member Family, conn pgx.Conn) {
	bd, err := conn.Begin(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	defer bd.Rollback(context.Background())
	bd.Exec(context.Background(), "insert into family_members(name, age, status, room) values ($1, $2, $3, $4)", member.Name, member.Age, member.Status, member.Room)
	bd.Commit(context.Background())
}

func ShowFamily(conn pgx.Conn) {
	var name, id, age, room string
	rows, err := conn.Query(context.Background(), "select id, name, age, room from family_members")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	} else {
		print("         ID     NAME     AGE       ROOM\n       ______________________________________\n")
		for rows.Next() {
			rows.Scan(&id, &name, &age, &room)
			fmt.Printf("%10s %10s %s %15s\n", id, name, age, room)
		}
	}
}
