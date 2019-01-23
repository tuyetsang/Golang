package demo_pointer

import "fmt"

func Demo1() {
	var a int = 5
	fmt.Println(a)
	fmt.Println(&a)
}
func Demo2() {
	var a int = 5
	var p *int = &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(a)
	fmt.Println(*p)
	*p = 6
	fmt.Println(a)
	q := &a
	fmt.Println(p)
	fmt.Println(*q)
	*q = 7
}
func Change1(a int) {
	a = 11
}
func Change2(p *int) {
	*p = 33
}
func Demo3() {
	a := 2
	Change1(a)
	fmt.Println(a)
	Change2(&a)
	fmt.Println(a)
}
func Swap(p *int, q *int) {
	tmp := *p
	*p = *q
	*q = tmp

}
func Demo4() {
	a, b := 5, 10
	Swap(&a, &b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func Modify1(a [3]int) {
	a[1] = 11
}
func Modify2(p *[3]int) {
	(*p)[1] = 22
}

func Demo5() {
	var a = [3]int{5, 9, 1}
	Modify1(a)
	fmt.Println(a)
	Modify2(&a)
	fmt.Println(a)
}

//Xay dung hamf swap doi 2 gia tri kieu nguyn a,b.su dung con tro
