package demo_struct

import(
	"entities"
	"fmt"
)

func Demo1(){
	var student entities.Student
	student.Id = "st01"
	student.Name = "TTA"
	student.Age = 20
	student.Score = 9
	fmt.Println(student)
	fmt.Println("st Id-->",student.Id,"st Name-->",student.Name,"st Age-->",student.Age,"st Score-->",student.Score)
}

func Demo2(){
	student := entities.Student {
		Id : "st02",
		Name : "TTA2",
		Age: 20,
		Score: 8.1,
	}
	fmt.Println(student)
}

func Demo3(){
	student := entities.Student{"st03","TTA2",28,9.4}
	fmt.Println(student)
}

func Demo4(){
	var emp entities.Emp
	emp.Id="emp1"
	emp.Fullname = entities.FullName{
		Firstname:"F1",
		Middlename:"MD1",
		LastName:"L1",
	}
	emp.Address =  entities.Address{
		Street:"ST1",
		Ward:"WA1",
		District:"DS1",
	}
	emp.Salary = 20;

	fmt.Println(emp)
}

func Demo5(){
	emp := entities.Emp {
		Id : "st02",
		Fullname : entities.FullName{"F01","MD01","L01"},
		Address: entities.Address{"St01","WA01","DS01"},
		Salary: 20,
	}
	fmt.Println(emp)
}

func Demo6(){
	student := entities.Student{"st03","TTA2",28,9.4}
	var p *entities.Student = &student
	fmt.Println("Name--->",(*p).Name)
	(*p).Age = 24
	fmt.Println(student)
}

func ChangeName(p *entities.Student){
	(*p).Name = "abc"
}

func Demo7(){
	student := entities.Student{"st03","TTA2",28,9.4}
	ChangeName(&student)
	fmt.Println(student)
}

func Demo8(){
	students :=[3]entities.Student{
		{"st01","TTA2",28,9.4},
		{"st02","TTA2",25,9.4},
		{"st03","TTA2",23,9.4},
	}
	for _,student :=range students{
		fmt.Println(ToString(student))
		fmt.Println("------------------------")
	}
}

func ToString(student entities.Student) string{
	return fmt.Sprintf("id: %s\nage:%d\ncore:%f",student.Id,student.Age,student.Score)
}