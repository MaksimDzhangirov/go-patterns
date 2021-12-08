package department

type nullDepartment struct {
	numberOfProfessors int
}

func NewNullDepartment() *nullDepartment {
	return &nullDepartment{}
}

func (d *nullDepartment) GetNumberOfProfessors() int {
	return 0
}

func (d *nullDepartment) GetName() string {
	return "nullDepartment"
}
