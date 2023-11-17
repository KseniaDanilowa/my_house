package appliances

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
