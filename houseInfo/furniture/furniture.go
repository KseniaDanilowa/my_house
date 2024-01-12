package furniture

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type Furniture struct {
	Name     string
	Width    float64
	Length   float64
	Height   float64
	Material string
	Color    string
}

func CreateFurniture() []Furniture {
	furniture := []Furniture{
		{Name: "Sofa", Width: 2.5, Length: 1.5, Height: 1.0, Material: "Wood", Color: "Maroon"},
		{Name: "Bed", Width: 2.0, Length: 2.0, Height: 1.5, Material: "Wood", Color: "Dark-Brown"},
		{Name: "Dining table", Width: 1.5, Length: 1.5, Height: 0.8, Material: "Wood", Color: "Beige"},
		{Name: "Wardrobe", Width: 1.0, Length: 2.0, Height: 2.5, Material: "Wood", Color: "Brown"},
		{Name: "Bookshelf", Width: 1.0, Length: 0.5, Height: 2.0, Material: "Wood", Color: "Light-Brown"},
	}
	return furniture
}

func InsertInFurniture(furniture []Furniture, conn pgx.Conn) {
	for index := 0; index < len(furniture); index++ {
		bd, beginErr := conn.Begin(context.Background())
		if beginErr != nil {
			fmt.Println(beginErr)
			return
		}
		defer bd.Rollback(context.Background())
		_, execErr := bd.Exec(context.Background(), "insert into furniture(name, width, length, height, material, color) "+
			"values ($1, $2, $3, $4, $5, $6)", furniture[index].Name, furniture[index].Width, furniture[index].Length, furniture[index].Height, furniture[index].Material, furniture[index].Color)
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

func ShowFurniture(conn pgx.Conn) {
	var name, width, length, height, material, color string
	rows, err := conn.Query(context.Background(), "select name, width, length, height, material, color from furniture")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	} else {
		print("\n              NAME         WIDTH      LENGTH     HEIGHT      MATERIAL     COLOR\n" +
			"        _________________________________________________________________________\n")
		for rows.Next() {
			rows.Scan(&name, &width, &length, &height, &material, &color)
			fmt.Printf("%20s %10s %10s %10s %12s %12s\n", name, width, length, height, material, color)
		}
	}
}
