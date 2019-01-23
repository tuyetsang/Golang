package entities

import (
	"fmt"
)

type Cat struct {
	Name string
}

func (cat Cat) Speak() {
	fmt.Println("Cat:", cat.Name, "speak meo meo ")
}
func (cat Cat) Type() string {
	return cat.Name
}
