package InsertHouseData

import (
	"context"
	"my_house/connection"
	"my_house/houseInfo/appliances"
	"my_house/houseInfo/family"
	"my_house/houseInfo/furniture"
	"my_house/houseInfo/houseParams"
)

func InsertData() {

	ctx := context.Background()
	conn := connection.MakeConn(ctx)
	appliances.InsertInAppliances(appliances.CreateAppliance(), conn)
	houseParams.InsertInHouseParams(houseParams.CreateHouseParams(), conn)
	furniture.InsertInFurniture(furniture.CreateFurniture(), conn)
	family.InsertInFamily(family.CreateFamily(), conn)
}
