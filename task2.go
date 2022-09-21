package i190582

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}
type company struct {
	companyName string
	employees   []employee
}

// A function to print company's details
func PrintCompany() {
	comp := CreateCompany()
	fmt.Println("Company Name: ", comp.companyName)
	fmt.Println("Employees Details:\n")
	for i := 0; i < 3; i++ {
		fmt.Println("Employee #: ", i+1)
		fmt.Println("Name: ", comp.employees[i].name)
		fmt.Println("Salary: ", comp.employees[i].salary)
		fmt.Println("Position: ", comp.employees[i].position)
		fmt.Println()
	}
}

func CreateEmployees() []employee {
	// Making three employees
	emp1 := employee{"Ansar", 100000, "UI Designer"}
	emp2 := employee{"Salar", 100000, "Database Admin"}
	emp3 := employee{"Nisar", 100000, "Head QA"}

	// Creating emplys array
	emplys := []employee{emp1, emp2, emp3}
	return emplys
}

func CreateCompany() company {
	emplys := CreateEmployees()
	// Creating a company struct
	comp := company{"Microsoft", emplys}
	return comp
}
