package main

import (
	"entities"
	"fmt"
	// "sort"
	// "strings"
)

// "mylib"
// "mylib/lib1"
// "mylib/lib2"
// samename "mylib/geometry/circle"
// "mylib/geometry/rectange"
// "demo_string",
// "sort"

// Function name always begin with upper word
func Demo1() {
	var student entities.Student
	student.Id = "st001"
	student.Name = "abc"
	student.Age = 20
	fmt.Println(student.ToString())
}
func Demo2() {
	product1 := entities.Product{"p01", "name", 20, 5}

	fmt.Println(product1.ToString())
	fmt.Println(product1.Total())
}
func Demo3() {
	student := entities.NewStudent("st02", "name2", 22)
	fmt.Println(student.ToString())

}
func Demo4() {
	student := entities.NewStudent("st02", "name2", 22)
	fmt.Println(student.ToString())
	student.SetName("abc")
	fmt.Println(student.ToString())
	(&student).SetAge(25)
	fmt.Println(student.ToString())
	student.SetAge(26)
	fmt.Println(student.ToString())
}
func Demo5() {
	category1 := entities.Category{"c01", "category1", []entities.Product{
		{"p01", "tv1", 20, 3},
		{"p02", "tv2", 21, 3},
	}}
	category2 := entities.Category{"c02", "category2", []entities.Product{
		{"p03", "tv3", 23, 3},
		{"p04", "tv4", 24, 3},
	}}
	fmt.Println(category1.ToString())
	fmt.Println(category1.ToString())
	categories := []entities.Category{category1, category2}
	fmt.Println("Total:")
	fmt.Println(Q1(categories))
	fmt.Println("Max")
	fmt.Println(category1.Max().ToString())
	fmt.Println("Max All")
	fmt.Println(Q3(categories).ToString())
}
func Q1(categories []entities.Category) (result float64) {
	result = 0
	for _, category := range categories {
		result += category.Total()
	}
	return
}
func Q3(categories []entities.Category) (m entities.Product) {
	m = categories[0].Max()
	for _, category := range categories {
		if m.Price < category.Max().Price {
			m = category.Max()
		}
	}
	return
}
func Demo6() {
	employee := entities.Employee{entities.Human{"name1", "male"}, "e01", 20}
	fmt.Println(employee.ToString())
}
func Demo7() {
	var human entities.Human
	var ihuman entities.IHuman = human
	fmt.Println(ihuman.Eat())
	ihuman.Sleep()
	var ihuman2 entities.IHuman2 = human
	ihuman2.Move()
}
func Demo8() {
	animals := []entities.Animal{
		entities.Duck{"Duck1"},
		entities.Cat{"Cat 1"},
		entities.Chicken{"Chicken 1"},
		entities.Cat{"Cat 2"},
		entities.Duck{"Duck 2"},
	}
	for _, animal := range animals {
		animal.Speak()
	}
}
func main() {
	// mylib.Hello()
	// fmt.Println(mylib.Hi("ABC"))
	// fmt.Println(lib1.Show("aa@gmail.com"))
	// fmt.Println(lib2.Display("TMA"))
	// fmt.Println("Dien tich hinh tron",samename.Dt(5.6))
	// fmt.Println("Chu vi hinh tron",samename.Cv(5.6))
	// fmt.Println("Dien tich hinh chu nhat",rectange.Dt(5,6))
	// fmt.Println("Chu vi hinh chu nhat",rectange.Cv(5,6))
	// var sp = []string{"banhoc", "ghesofa", "gheluoi", "ghego", "bancafe"}
	// keyword := "ghe"
	// fmt.Println(demo_string.Search(sp, keyword))
	// fmt.Println(demo_string.SortNameASC(sp))
	// array.Demo2()
	// a:=[]int {5,10,12,33,-1,-12}
	Demo8()

	// array.Demo3(a)
	// array.Demo4()
	// fmt.Println(array.Demo5(a))
	// r1,r2,r3,r4 := array.Demo6(a)
	// fmt.Println("Duong",r1,"Am",r2,"Chan",r3,"Le",r4)

	// array.Demo7()
	// array.Demo8()
	// array.Demo9()
}

// func Demo2() {
// 	student := entities.Student{
// 		Id:    "st01",
// 		Name:  "abc",
// 		Score: 5.6,
// 		Age:   20,
// 	}
// 	fmt.Println(student)
// }

// func Demo3() {
// 	student := entities.Student("st01", "avc", 5.6, 20)
// 	fmt.Println(student)
// }
// func Demo4() {
// 	var employee entities.Employee
// 	employee.Id = "e01"
// 	employee.FullName = entities.FullName{
// 		FirstName: "FN 1",
// 		LastName:  "LN 1",
// 	}
// 	employee.Address = entities.Address{
// 		Stress:  "123 HCM",
// 		Ward:    "P10",
// 		Distric: "Q.12",
// 	}
// 	employee.Salary = 20
// 	fmt.Println(employee)
// 	fmt.Println("id:", employee.Id)
// 	fmt.Println("Fullname:", employee.FullName.FirstName)
// }

// func Demo5() {
// 	employee := entities.Employee{
// 		Id: "e02",
// 		FullName: entities.FullName{
// 			FirstName: "FN 1",
// 			LastName:  "LN 1",
// 		},
// 		Address: entities.Address{
// 			Stress:  "123 HCM",
// 			Ward:    "P10",
// 			Distric: "Q.12",
// 		},
// 		Salary: 20.0,
// 	}
// 	fmt.Println(employee)

// }

// func ToString(student entities.Student) string {
// 	return fmt.Sprintf("Id:%s\nName:%s\nScore:%f\nAge:%d", student.Id, student.Name, student.Score, student.Age)
// }
// func Demo6() {
// 	student := entities.Student{
// 		Id:    "st01",
// 		Name:  "abc",
// 		Score: 5.6,
// 		Age:   20,
// 	}
// 	fmt.Println(ToString(student))
// }
// func Demo7() {
// 	p := &entities.Student{
// 		Id:    "st01",
// 		Name:  "abc",
// 		Score: 5.6,
// 		Age:   20,
// 	}
// 	fmt.Println("id:", (*p).Id)
// 	fmt.Println("Name:", (*p).Name)

// }

// func Change1(student entities.Student) {
// 	student.Name = "ABC"
// }
// func Change2(student *entities.Student) {
// 	(*student).Name = "ABC"
// }
// func Demo8() {
// 	student := entities.Student{
// 		Id:    "st01",
// 		Name:  "abc",
// 		Score: 5.6,
// 		Age:   20,
// 	}
// Change1(student)
// fmt.Println(ToString(student))
// 	Change2(&student)
// 	fmt.Println(ToString(student))
// }

// func Demo9() {
// 	students := [3]entities.Student{
// 		{"st01", "SV1", 7.8, 20},
// 		{"st02", "SV2", 7.9, 20},
// 		{"st03", "SV3", 7.4, 20},
// 	}
// 	for _, student := range students {
// 		fmt.Println(ToString(student))
// 		fmt.Println("...................")
// 	}
// }c ToString2(product entities.Product) string {
// 	return fmt.Sprintf("Id:%s\nName:%s\nPrice:%f\nQuantity:%d", product.Id, product.Name, product.Price, product.Quantity)
// }
// func TongTien(products []entities.Product) (total float64) {
// 	total = 0
// 	for _, product := range products {
// 		total += product.Price * float64(product.Quantity)
// 	}
// 	return
// }

// // func Max(products []entities.Product) (max entities.Product, min entities.Product) {
// // 	max = products[0]
// // 	min = products[0]
// // 	for _, product := range products {
// // 		if product >= max {
// // 			product = max
// // 		} else {
// // 			product = min
// // 		}
// // 	}
// // }
// func TimKey(products []entities.Product, keyword string) (result []entities.Product) {
// 	for _, product := range products {
// 		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(keyword)) {
// 			result = append(result, product)
// 		}
// 	}
// 	return
// }

// func Demo10() {

// 	products := [5]entities.Product{
// 		{"p01", "SP1", 1, 20},
// 		{"p02", "SP2", 2, 21},
// 		{"p03", "SP3", 3, 20},
// 		{"p04", "SP4", 4, 20},
// 		{"p05", "SP5", 5.2, 20},
// 	}

// 	fmt.Println("total:", TongTien(products[:]))
// max, min := Max(products[:])
// fmt.Println("Max:")
// fmt.Println(ToString2(max))
// fmt.Println("Min:")
// fmt.Println(ToString2(min))
// 	fmt.Println("Sap xep giam dan theo gia:")
// 	product := products[:]
// 	sort.Slice(product, func(i, j int) bool {
// 		return product[i].Price >= product[j].Price
// 	})
// 	fmt.Println(products)
// 	fmt.Println("Sap xep tang dan theo gia:")
// 	sort.Slice(product, func(i, j int) bool {
// 		return product[i].Price <= product[j].Price
// 	})
// 	fmt.Println(products)
// 	result := TimKey(products[:], "SP1")
// 	fmt.Println(result)
// }
// func Demo11() {
// 	student1 := entities.Student{"st01", "SV1", 7.8, 20}
// 	student2 := entities.Student{"st01", "SV1", 7.8, 21}
// 	if student1 == student2 {
// 		fmt.Println("=")
// 	} else {
// 		fmt.Println("#")
// 	}
// }

/*Khai bao struct Product chua thong tin sp co cac thuoc tinh
1.id string
2.name string
3.price float64
4.quantity int
Khai bao 1 mang cos 5 sp. Khai bao cac ham thuc hien cac hanh dong sau:
a)Tinh tong tien tat ca sp
b)Liet ke thong tin sp cos gia lon nhat,nho nhat
c)Sap xep cac sp theo thu tu tang dan,giam dan cua gia
d)Tim danh sach cac sp co ten chua 1 keyword
*/
