package entities

import (
	"fmt"
)

type Duck struct {
	Name string
}

func (duck Duck) Speak() {
	fmt.Println("Duck:", duck.Name, "speak  cap cap ")
}
func (duck Duck) Type() string {
	return duck.Name
}
