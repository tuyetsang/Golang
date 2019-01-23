package array_2d

import "fmt"

func Demo1() {
	var a [2][3]int
	a[0][0] = 3
	a[0][1] = -2
	a[0][2] = 11
	a[1][0] = -6
	a[1][1] = 2
	a[1][2] = 7
	fmt.Println(a)
	fmt.Println("len:", len(a))
	for r := 0; r < len(a); r++ {
		for c := 0; c < len(a[r]); c++ {
			fmt.Println("\t", a[r][c])
		}
		fmt.Println()
	}
}
func Demo2() {
	a := [2][3]int{
		{4, 2, -9},
		{11, 8, 60},
	}
	fmt.Println(a)
	for i, row := range a {
		fmt.Println("Row:", i)
		fmt.Println("\t", row)
		for j, v := range row {
			fmt.Println("j:", j, ",v:", v)
		}
	}
}
