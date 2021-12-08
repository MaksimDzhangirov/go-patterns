package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/null-object/code/college"
)

func main() {
	c1 := createCollege1()
	c2 := createCollege2()
	totalProfessors := 0
	departmentArray := []string{"computerScience", "mechanical", "civil", "electronics"}

	for _, departmentName := range departmentArray {
		d := c1.GetDepartment(departmentName)
		totalProfessors += d.GetNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)

	// Обнуляем общее число профессоров
	totalProfessors = 0
	for _, departmentName := range departmentArray {
		d := c2.GetDepartment(departmentName)
		totalProfessors += d.GetNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}

func createCollege1() *college.College {
	c := college.NewCollege()
	c.AddDepartment("computerScience", 4)
	c.AddDepartment("mechanical", 5)
	return c
}

func createCollege2() *college.College {
	c := college.NewCollege()
	c.AddDepartment("computerScience", 2)
	return c
}
