package demo_map

import "fmt"

func Demo1() {
	student := make(map[string]string)
	student["id"] = "st001"
	student["name"] = "name1"
	student["address"] = "address1"
	fmt.Println(student)
	fmt.Println("Student Info")
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}
func Demo2() {
	student := map[string]string{
		"id":      "st001",
		"name":    "name1",
		"address": "address1",
	}
	student["phone"] = "12345"

	fmt.Println(student)

}
func Demo3() {
	student :=
		map[string]string{
			"id":      "st001",
			"name":    "name1",
			"address": "address1",
		}
	value, ok := student["name"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("error")
	}
	delete(student, "address")
	fmt.Println(student)

}
func Display(student map[string]string) {
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}
func Demo4() {
	student := map[string]string{
		"id":      "st001",
		"name":    "name1",
		"address": "address1",
	}
	Display(student)
}
func Demo5() {
	categories := map[string][]string{
		"categories1": []string{"product 1", "product 2"},
		"categories2": []string{"product 3", "product 4"},
		"categories3": []string{"product 5", "product 6"},
	}
	fmt.Println(categories)
	categories["categories4"] = []string{"product 7", "product 8"}
	fmt.Println(categories)
	categories["categories4"] = append(categories["categories4"], "product 7.1")
	for key, value := range categories {
		fmt.Println(key)
		for _, name := range value {
			fmt.Println(name)
		}
	}
}
