package printHouse

import (
	"fmt"
	"my_house/houseInfo/appliances"
	"my_house/houseInfo/family"
	"my_house/houseInfo/furniture"
	"my_house/houseInfo/houseParams"
)

func PrintHouseInfo() {
	appliances := appliances.CreateAppliances()
	furniture := furniture.CreateFurniture()
	familyMembers := family.CreateFamily()
	params := houseParams.CreateHouseParams()

	fmt.Println("	House Params:")
	fmt.Printf("Width: %d, Height: %d, Length: %d, Rooms Count: %d\n", params.Width, params.Height, params.Length, params.RoomsCount)

	fmt.Println("	Appliances:")
	for _, appliance := range appliances {
		fmt.Printf("Name: %s\nBrand: %s, Power: %d\n", appliance.Name, appliance.Brand, appliance.Power)
	}

	fmt.Println("	Furniture:")
	for _, furnishing := range furniture {
		fmt.Printf("Name: %s\nWidth: %0.1f, Length: %0.1f, Heght: %0.1f, Color: %s, Material: %s\n", furnishing.Name,
			furnishing.Width, furnishing.Length, furnishing.Height, furnishing.Color, furnishing.Material)
	}

	fmt.Println("	Family Members:")
	for _, member := range familyMembers {
		fmt.Printf("Name: %s\nStatus: %s, Age: %d\n", member.Name, member.Status, member.Age)
	}

}
