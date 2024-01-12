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
		{Name: "Sergei", Age: 21, Status: "Step-son", Room: "living_room"},
		{Name: "TAGIR", Age: 98, Status: "Brother", Room: "bathroom"},
		{Name: "Danechka", Age: 20, Status: "family_friend", Room: "living_room"},
	}
	return familyMembers
}

func InsertInFamily(members []Family, conn pgx.Conn) {
	for index := 0; index < len(members); index++ {
		bd, beginErr := conn.Begin(context.Background())
		if beginErr != nil {
			fmt.Println(beginErr)
			return
		}
		defer bd.Rollback(context.Background())
		_, execErr := bd.Exec(context.Background(), "insert into family_members(name, age, status, room) values ($1, $2, $3, $4)",
			members[index].Name, members[index].Age, members[index].Status, members[index].Room)
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
}

func ShowFamily(conn pgx.Conn) {
	var name, id, age, room, status string
	rows, err := conn.Query(context.Background(), "select id, name, age, room, status from family_members")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	} else {
		print("        ID      NAME     AGE       ROOM             STATUS\n" +
			"       ____________________________________________________\n")
		for rows.Next() {
			rows.Scan(&id, &name, &age, &room, &status)
			fmt.Printf("%10s %10s %5s %15s %15s\n", id, name, age, room, status)
		}
	}
}
