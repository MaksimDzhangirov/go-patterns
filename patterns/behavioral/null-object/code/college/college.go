package college

import (
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/null-object/code/college/department"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/null-object/code/interfaces"
)

type College struct {
	departments []interfaces.Department
}

func NewCollege() *College {
	return &College{}
}

func (c *College) AddDepartment(departmentName string, numOfProfessors int) {
	if departmentName == "computerScience" {
		computerScienceDepartment := department.NewComputerScience(numOfProfessors)
		c.departments = append(c.departments, computerScienceDepartment)
	}
	if departmentName == "mechanical" {
		mechanicalDepartment := department.NewMechanical(numOfProfessors)
		c.departments = append(c.departments, mechanicalDepartment)
	}
	return
}

func (c *College) GetDepartment(departmentName string) interfaces.Department {
	for _, currentDepartment := range c.departments {
		if currentDepartment.GetName() == departmentName {
			return currentDepartment
		}
	}
	// Возвращаем nullDepartment, если такого факультета не существует
	return department.NewNullDepartment()
}
