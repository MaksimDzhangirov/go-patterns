package department

type mechanical struct {
	numberOfProfessors int
}

func NewMechanical(numberOfProfessors int) *mechanical {
	return &mechanical{
		numberOfProfessors: numberOfProfessors,
	}
}

func (m *mechanical) GetNumberOfProfessors() int {
	return m.numberOfProfessors
}

func (m *mechanical) GetName() string {
	return "mechanical"
}
