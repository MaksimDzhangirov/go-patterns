package common

type Short struct {
	logo string
	size int
}

func NewShort(logo string, size int) Short {
	return Short{
		logo: logo,
		size: size,
	}
}

func (s *Short) SetLogo(logo string) {
	s.logo = logo
}

func (s *Short) GetLogo() string {
	return s.logo
}

func (s *Short) SetSize(size int) {
	s.size = size
}

func (s *Short) GetSize() int {
	return s.size
}