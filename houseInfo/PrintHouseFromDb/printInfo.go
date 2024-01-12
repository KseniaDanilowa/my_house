package PrintHouseFromDb

import (
	"context"
	"my_house/connection"
	"my_house/houseInfo/family"
	"my_house/houseInfo/furniture"
)

func PrintHouse() {
	ctx := context.Background()
	conn := connection.Make_conn(ctx)
	family.ShowFamily(conn)
	furniture.ShowFurniture(conn)

}
