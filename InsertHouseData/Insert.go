package InsertHouseData

import (
	"context"
	"my_house/connection"
	"my_house/houseInfo/appliances"
)

func InsertData() {
	ctx := context.Background()
	conn := connection.MakeConn(ctx)
	appliances.InsertInAppliances(appliances.CreateAppliances(), conn)
	//furniture.InsertInFurniture(furniture.CreateFurniture(), conn)
	//family.InsertInFamily(family.CreateFamily()[1], conn)
}
