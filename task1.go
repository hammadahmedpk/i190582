package i190582

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// Printing the person's values
func PrintStruc(person Person) {
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)
	fmt.Println("Job: ", person.job)
	fmt.Println("Salary: ", person.salary)
	fmt.Println()
}
