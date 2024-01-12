package InsertHouseData

import (
	"context"
	"my_house/connection"
	"my_house/houseInfo/family"
	"my_house/houseInfo/furniture"
)

func InsertData() {
	ctx := context.Background()
	conn := connection.Make_conn(ctx)
	furniture.InsertInFurniture(furniture.CreateFurniture(), conn)
	family.InsertInFamily(family.CreateFamily()[1], conn)
}
