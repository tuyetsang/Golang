package demo_string

import (
	"fmt"
	"sort"
	"strings"
)

func Demo1() {
	s := "abc"
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("\t%c", s[i])
	}

}
func Demo2() {
	s := "     abc           "
	fmt.Println(len(s))
	s = strings.TrimSpace(s)
	fmt.Println(len(s))

}
func Demo3() {
	var ch rune = 'a'
	fmt.Printf("%d", ch)
	fmt.Printf("\n%c\n", ch)
	s := "123abc456"
	result := strings.TrimFunc(s, func(c rune) bool {
		return c >= 48 && c <= 57
	})
	fmt.Println(result)
}
func Demo4() {
	s1 := "abc"
	s1 = strings.ToUpper(s1)
	fmt.Println(s1)

	s2 := "DEF"
	s2 = strings.ToLower(s2)
	fmt.Println(s2)
}
func Demo5() {
	s1 := "nokia"
	s2 := "No"
	result1 := strings.HasPrefix(s1, s2)
	fmt.Println(result1)

	result2 := strings.HasPrefix(strings.ToLower(s1), strings.ToLower(s2))
	fmt.Println(result2)
	result3 := strings.HasSuffix(s1, s2)
	fmt.Println(result3)
	result4 := strings.Contains(s1, s2)
	fmt.Println(result4)
}

func Demo6() {
	name1 := "mary"
	name2 := "peter"
	result := strings.Compare(name1, name2)
	fmt.Println(result)
}

/*
Tao mang kieu chuoi co 5 phan tu chua ten sp
xay dung ham thuc hien cac yeu cau
1.xay dung ham tra ve danh sach cac sp co ten chua 1 keyword
1.sap xep cac ten theo thu tu tang dan, giam dan
*/

// func Demo7() {
// 	var sp = [5]string{"banhoc", "ghesofa", "gheluoi", "ghego", "bancafe"}
// 	keyword := "ghe"
// 	for r := 0; r < len(sp); r++ {
// 		result1 := strings.HasPrefix(sp, keyword)
// 		fmt.Println(result1)
// 	}

// 	sp1 := sp[:]
// 	sort.Slice(sp1, func(i, j int) bool {
// 		return sp1[i] >= sp1[j]
// 	})
// 	fmt.Println(sp)
// }

func Search(names []string, keyword string) (result []string) {
	for _, name := range names {
		if strings.Contains(strings.ToLower(name), strings.ToLower(keyword)) {
			result = append(result, name)
		}

	}
	return
}
func SortNameASC(names []string) []string {
	sort.Slice(names, func(i, j int) bool {
		return strings.Compare(names[i], names[j]) < 0
	})
	return names
}
