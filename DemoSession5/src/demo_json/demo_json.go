package demo_json

import (
	"demo_file/entities"
	"encoding/json"
	"fmt"
)

func Demo1() {
	product := entities.Product{"p01", "name1", 2, 6}
	str_json, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(str_json))
	}
}
func Demo2() {
	str := `{"Id":"p01","Name":"name1","Price":2,"Quantity":6}`
	var product entities.Product
	json.Unmarshal([]byte(str), &product)
	fmt.Println(product.ToString())
}
func Demo3() {
	str := `[
		{"Id":"p01","Name":"name1","Price":1,"Quantity":4},
		{"Id":"p02","Name":"name2","Price":2,"Quantity":5},
		{"Id":"p03","Name":"name3","Price":3,"Quantity":6}
	]`
	var products []entities.Product
	json.Unmarshal([]byte(str), &products)
	for _, product := range products {
		fmt.Println(product.ToString())

	}
}
func Demo4() {
	str := `{"Id":"st01","Name":"sang","Address":{"Street":"123","Ward":"P10"}}`
	var student entities.Student
	err := json.Unmarshal([]byte(str), &student)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Student Info")
		fmt.Println("Id:", student.Id)
		fmt.Println("Name:", student.Name)
		fmt.Println("Street:", student.Address.Street)
		fmt.Println("Ward:", student.Address.Ward)
	}
}
func Demo5(){
	file, err := os.Open("demo_json/invoke.json")
	if err != nil {
		fmt.Println(err)
	} else {
		var invokes []entities.Invoke
		
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			result := strings.Split(line, ",")
			product := entities.Product{
				Id:       result[0],
				Name:     result[1],
				Price:    price,
				Quantity: quantity,
			}
			products = append(products, product)
		}
		fmt.Println("product list")
		for _,invoke:=range invokes{
			err := json.Unmarshal([]byte(str), &invoke)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Invoke Info")
				fmt.Println("Id:", invoke.Id)
				fmt.Println("Name:", invoke.Name)
				fmt.Println("Street:", student.Address.Street)
				fmt.Println("Ward:", student.Address.Ward)
			}
		}
		
		}
	}
	file.Close()
}
