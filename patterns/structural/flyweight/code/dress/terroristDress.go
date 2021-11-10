package dress

type terroristDress struct {
	color string
}

func NewTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

func (t *terroristDress) GetColor() string {
	return t.color
}
