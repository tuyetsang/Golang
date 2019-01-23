package mylib
import "fmt"

func Hello(){
	fmt.Println("Hello")
}

func Hi(fullName string) string {
	return fmt.Sprintf("Hello %s",fullName)
}