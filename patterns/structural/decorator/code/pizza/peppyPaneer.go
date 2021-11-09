package pizza

type peppyPaneer struct {
}

func NewPeppyPaneer() *peppyPaneer {
	return &peppyPaneer{}
}

func (p *peppyPaneer) GetPrice() int {
	return 20
}
