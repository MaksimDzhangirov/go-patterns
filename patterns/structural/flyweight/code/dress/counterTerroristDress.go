package dress

type counterTerroristDress struct {
	color string
}

func NewCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}

func (c *counterTerroristDress) GetColor() string {
	return c.color
}
