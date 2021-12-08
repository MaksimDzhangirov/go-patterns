package department

type computerScience struct {
	numberOfProfessors int
}

func NewComputerScience(numberOfProfessors int) *computerScience {
	return &computerScience{
		numberOfProfessors: numberOfProfessors,
	}
}

func (c *computerScience) GetNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *computerScience) GetName() string {
	return "computerScience"
}
