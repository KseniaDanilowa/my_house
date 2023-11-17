package houseParams

type HouseParams struct {
	Width      int
	Height     int
	Length     int
	RoomsCount int
}

func CreateHouseParams() HouseParams {
	params := HouseParams{
		Width:      10,
		Height:     8,
		Length:     12,
		RoomsCount: 1,
	}
	return params
}
