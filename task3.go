package i190582

import (
	"fmt"
	"strings"
)

type Student struct {
	rollNumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollNumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

// The function to print student details
func Print(students *StudentList) {
	for i := 0; i < len(students.list); i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))
		fmt.Println("Student Rollno: ", students.list[i].rollNumber)
		fmt.Println("Student Rollno: ", students.list[i].name)
		fmt.Println("Student Rollno: ", students.list[i].address)
	}
}
