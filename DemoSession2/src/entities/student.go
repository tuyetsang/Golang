package entities

import (
	"fmt"
)

type Student struct {
	Id   string
	Name string
	Age  int
}

func (student Student) ToString() string {
	return fmt.Sprintf("id:%s\nname:%s\nage:%d", student.Id, student.Name, student.Age)
}
func NewStudent(id, name string, age int) Student {
	student := Student{id, name, age}
	return student
}
func (student Student) SetName(name string) {
	student.Name = name
}
func (student *Student) SetAge(age int) {
	(*student).Age = age
}
