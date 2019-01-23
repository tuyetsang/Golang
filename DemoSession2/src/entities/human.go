package entities

import (
	"fmt"
)

type Human struct {
	Name, Gender string
}

func (human Human) ToString() string {
	return fmt.Sprintf("Name:%s\nGender:%s", human.Name, human.Gender)
}
func (human Human) Eat() string {
	return fmt.Sprintf("Human eat")
}
func (human Human) Sleep() {
	fmt.Println("Human sleep")
}
func (human Human) Move() {
	fmt.Println("Human move")
}
