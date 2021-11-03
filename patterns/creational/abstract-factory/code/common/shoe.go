package common

type Shoe struct {
	logo string
	size int
}

func NewShoe(logo string, size int) Shoe {
	return Shoe{
		logo: logo,
		size: size,
	}
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
	return s.size
}
