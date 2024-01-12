package PrintHouseFromDb

import (
	"context"
	"my_house/connection"
	"my_house/houseInfo/appliances"
	"my_house/houseInfo/family"
	"my_house/houseInfo/furniture"
)

func PrintHouse() {
	ctx := context.Background()
	conn := connection.MakeConn(ctx)
	family.ShowFamily(conn)
	furniture.ShowFurniture(conn)
	appliances.ShowAppliances(conn)
}
