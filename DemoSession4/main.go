package main
import(
	"fmt"
	"entities"
)

func main(){
	//Demo1()
	//Demo2()
	//Demo4()
	Demo6()
}
func Demo1(){
	student1 := entities.Student{"st01","n01",9}
	fmt.Println(student1.ToString())
	
}

func Demo2(){
	pro1:= entities.Product{"pro01","name01",10,12.5}
	fmt.Println(pro1.ToString())
	fmt.Println("Total",pro1.Total())
}
func Demo3(){
	student1 :=entities.NewStudent("st02","n902",20)
	fmt.Println(student1.ToString())
}

func Demo4(){
	student1 :=entities.NewStudent("st02","n902",20)
	student1.SetName("abccc")
	fmt.Println(student1.ToString())
	(&student1).SetAge(23)
	fmt.Println(student1.ToString())
	student1.SetAge(25)
	fmt.Println(student1.ToString())
}

func Demo5(){
	cate1:=entities.Category{"cat01","name01",[]entities.Product{entities.Product{"pro01","name01",12,12}}}
	result:= cate1.ToString()
	for _,item := range cate1.Products{
		result += item.ToString()
	}
	fmt.Println("Result",result)
}

func Demo6(){
	cate1:=entities.Category{"cat01","name01",[]entities.Product{
			{"pro01","name01",12,12},
			{"pro011","name011",10,122},
	}}
	cate2:=entities.Category{"cat02","name02",[]entities.Product{
		{"pro021","name21",12,12},
		{"pro022","name22",12,12},		
	}}
	// Tong tien danh muc thu nhat
	var tongtien=float32(0);
	for _,item:=range cate1.Products{
			tongtien = float32(tongtien) + item.Total()
	}
	for _,item:=range cate2.Products{
		tongtien = float32(tongtien) + item.Total()
	}
	fmt.Println("Tong tien 2 danh sach",tongtien)

}

