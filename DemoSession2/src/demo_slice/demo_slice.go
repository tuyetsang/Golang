package demo_slice

import (
	"fmt"
	"sort"
)

func Demo1() {
	var a = [5]int{4, -1, 2, 5, 8}
	fmt.Println(a)
	a1 := a[0:2]
	fmt.Println(a1)
	a[1] = 999
	fmt.Println(a)
	a2 := a[0:4]
	fmt.Println(a2)
	a[1] = 99
	fmt.Println(a1)
	fmt.Println(a)
}
func Demo2() {
	var a = [5]int{4, 1, -2, 5, 8}
	a1 := a[0:2]
	fmt.Println(a)
	a1 = append(a1, 123)
	fmt.Println(a1)
}
func Demo3() {
	var a = [5]int{4, 1, -2, 5, 8}
	a1 := a[:2]
	fmt.Println(a1)
	a2 := a[2:]
	fmt.Println(a2)
	a3 := a[:]
	fmt.Println(a3)
}
func Demo4() {
	var a = [5]int{4, 1, -2, 5, 8}
	fmt.Println("len:", len(a))
	a1 := a[0:2]
	fmt.Println(a1)
	fmt.Println("\tlen:", len(a1))
	fmt.Println("\tcap:", cap(a1))
	a2 := a[1:4]
	fmt.Println(a2)
	fmt.Println("\tlen:", len(a2))
	fmt.Println("\tcap:", cap(a2))
	a3 := a[4:]
	fmt.Println(a3)
	fmt.Println("\tlen:", len(a3))
	fmt.Println("\tcap:", cap(a3))
	a3 = append(a3, 123)
	fmt.Println("\tlen:", len(a3))
	fmt.Println("\tcap:", cap(a3))
}
func Demo5() {
	a := make([]int, 3, 5)
	a[0] = 4
	a[1] = 2
	a[2] = -4
	fmt.Println(a)
	fmt.Println("\tlen:", len(a))
	fmt.Println("\tcap:", cap(a))
	a = append(a, 123, 12, 1)
	fmt.Println(a)
	fmt.Println("\tlen:", len(a))
	fmt.Println("\tcap:", cap(a))
}
func Change1(a [3]int) {
	a[1] = 999
}

func Change2(a []int) {
	a[1] = 999
}

func Demo6() {
	var a = [3]int{4, 1, -2}
	Change1(a)
	fmt.Println(a)
	Change2(a[:])
	fmt.Println(a)
}
func Demo7() {
	var a = [5]int{4, 1, -2, 5, 8}
	fmt.Println(a)
	b := a[:]
	sort.Slice(b, func(i, j int) bool {
		return b[i] >= b[j]
	})
	fmt.Println(a)
}
