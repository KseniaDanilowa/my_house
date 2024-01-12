package appliances

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type Appliance struct {
	Name  string
	Brand string
	Power int
}

func CreateAppliances() []Appliance {
	appliances := []Appliance{
		{Name: "Refrigerator", Brand: "Samsung", Power: 200},
		{Name: "Microwave", Brand: "LG", Power: 1000},
		{Name: "Washing machine", Brand: "Bosch", Power: 500},
		{Name: "Air conditioner", Brand: "Haier", Power: 1500},
		{Name: "TV", Brand: "Sony", Power: 200},
	}
	return appliances
}

func InsertInAppliances(appliance []Appliance, conn pgx.Conn) {
	for index := 0; index < len(appliance); index++ {
		bd, beginErr := conn.Begin(context.Background())
		if beginErr != nil {
			fmt.Println(beginErr)
			return
		}
		defer bd.Rollback(context.Background())
		_, execErr := bd.Exec(context.Background(), "insert into appliances(name, brand, power) "+
			"values ($1, $2, $3)", appliance[index].Name, appliance[index].Brand, appliance[index].Power)
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

func ShowAppliances(conn pgx.Conn) {
	var name, brand, power string
	rows, err := conn.Query(context.Background(), "select name, brand, power from appliances")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	} else {
		print("\n                NAME            BRAND        POWER\n" +
			"        ___________________________________________\n")
		for rows.Next() {
			rows.Scan(&name, &brand, &power)
			fmt.Printf("%25s %12s %12s\n", name, brand, power)
		}
	}
}
