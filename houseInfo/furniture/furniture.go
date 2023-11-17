package furniture

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
