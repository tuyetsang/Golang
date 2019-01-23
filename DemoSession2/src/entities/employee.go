package entities

import (
	"fmt"
)

type Address struct {
	Stress, Ward, Distric string
}
type FullName struct {
	FirstName, LastName string
}
type Employee struct {
	Human Human
	Id    string
	// FullName FullName
	// Address  Address
	Salary float64
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("Id:%s\n%s\nSalary:%f", employee.Id, employee.Human.ToString(), employee.Salary)
}
