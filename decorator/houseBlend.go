package decorator

type HouseBlend struct {
	Description string
	Beverage
}

// NewHouseBlend
// Constructor 대응
func NewHouseBlend() Beverage {
	hb := &HouseBlend{
		Description: "하우스 블렌드 커피",
	}
	return hb
}

func (hb *HouseBlend) Cost() float64 {
	return 0.89
}

func (hb *HouseBlend) GetDescription() string {
	return hb.Description
}
