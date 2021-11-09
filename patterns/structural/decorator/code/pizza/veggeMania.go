package pizza

type veggeMania struct {
}

func NewVeggeMania() *veggeMania {
	return &veggeMania{}
}

func (p *veggeMania) GetPrice() int {
	return 15
}
