package demo_pointer
import(
	"fmt"
)

func Demo1(){
	var a int =5
	fmt.Println("a-->",a)
	fmt.Println("Address of a ---->",&a)
	// dua con tro p ve a
	var p *int = &a
	fmt.Println("Address of p--->",p)
	// in ra gia tri cua con tro p
	fmt.Println("p---->",*p)
	var q *int = &a
	fmt.Println("Address of p--->",q)
	fmt.Println("q---->",*q)
	*q = 10
	fmt.Println("a new ---->",a)
	fmt.Println("*p new ---->",*p)
}

func Change1(a int){
	a = 123
}
func Change2(p *int){
	*p = 456
}

func Swap(a *int, b *int){
	temp := *a
	*a = *b
	*b = temp

}
func Demo2(){
	a := 5
	Change1(a)
	fmt.Println("a-------->",a)
	c :=6
	Change2(&c)
	fmt.Println("c new-------->",c)
	d :=7
	e :=8
	Swap(&d,&e)
	fmt.Println("after swap--->",d,e)
	
}