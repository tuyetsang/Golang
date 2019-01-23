package main

import "fmt"

func main(){
	fmt.Println("Hello")
	// Demo1()
	// Demo2()
	// Demo3()
	// Demo4()
	// Demo5()
	display1("001","name001")
	//r1,r2,r3,r4:=calculate2(7,4)
	r1,_,r3,_:=calculate2(7,4)
	//fmt.Print("value1-->",r1,"value2-->",r2,"value3-->",r3,"value4-->",r4)
	fmt.Print("value1-->",r1,"value3-->",r3)

	display3(1,2)

	display3(1,2,3)
}

//void 
func Demo1(){
	var age int
	var fullName string = "Nguyen Van A"
	var status bool = true
	var price float32
	age = 20
	price = 4.5
	fmt.Println("age:",age,"full name:",fullName)
	fmt.Println("status:",status)
	fmt.Println("Price:",price)
}

func Demo2(){
	var age=20
	var address ="ABC"
	var a1,a2,a3 = 20,"ABBB",7
	fmt.Println("age:",age,"address:",address,"a1:",a1,"a2:",a2,"a3:",a3)
}

func Demo3(){
	age:= 20
	id,name,price := "p01","name01",4.5
	fmt.Println("age--->",age,"item_id",id,"item_name",name,"item_price",price)
}

func Demo4(){
	var a int =2
	var b float32=4.5
	var result = float32(a)+b
	fmt.Println("result:",result)
}

func Demo5(){
	var age int = 20
	fmt.Printf("age: %d",age)
	var status bool = true
	fmt.Printf("\n status: %t",status)
	var fullName string ="ABC"
	fmt.Printf("\n Full name: %s",fullName)
	var price float32 = 4.5
	fmt.Printf("\n Price: %f",price)
}

func Demo6(){
	var age int
	fmt.Print("age:")
	fmt.Scanf("%d\n",&age)
	var price float32
	fmt.Print("price:")
	fmt.Scanf("%f\n",&price)
	var fullname string
	fmt.Print("fullname:")
	fmt.Scanf("%s\n",&fullname)
	var status bool
	fmt.Print("status:")
	fmt.Scanf("%t",&status)
	fmt.Println("age:",age,"price:",price,"fullname:",fullname,"status:",status)
}

func Demo7(){
	var a = 5
	if a >2{
		fmt.Printf("a >2")
	} else {
		fmt.Println("a<2")
	}
}

func Demo8(){
	if a:=5; a>2{
		fmt.Printf("a >2")
	} else {
		fmt.Println("a<2")
	}
}

func Demo9(){
	menu:=1
	if menu == 1{
		fmt.Println("Menu 1 selected")
	} else if menu==2{
		fmt.Println("Menu 2 selected")
	} else {
		fmt.Println("Invalid")
	}
}

func Demo10(){
	
	switch menu:=1; menu {
		case 1:
			fmt.Println("Menu 1 selected")
		case 2:
			fmt.Println("Menu 2 selected")
		default:
			fmt.Println("Invalid")
	}
}

func Demo11(){
	switch a:=1 ;a{
		case 1,2,3:
			fmt.Println("a=1,2,3")
		default:
			fmt.Println("Invalid")
	}
}

func Demo12(){
	a:=5
	switch {
		case a>=1 && a<=10:
			fmt.Println("a>=1 && a<=10")
		case a<=20:
			fmt.Println("a<=20")
		default:
			fmt.Println("Invalid")
	}
}

func Demo13(){
	for i:=1;i<=10;i++{
		fmt.Print(" ",i)
	}
}

func Demo14(){
	i:=1
	for i<=10{
		fmt.Print(" ",i)
		i++
	}
}

func Demo15_Do_While(){
	i:=1
	for{
		fmt.Print(" ",i)
		if i>10 {break}
		i++
	}
}

func Hello(fullName string){
	fmt.Println("Hello",fullName)
}


func display2(id string, name string){
	fmt.Println("id:",id,"name:",name)
}

func display1(id, name string){
	fmt.Println("id:",id,"name:",name)
}

func sum(a,b int) int{
	return a+b
}

func calculate(a,b int)(int,int,int,float32){
	result1:=a+b
	result2:= a-b
	result3:=a*b
	result4:= float32(a) / float32(b)
	return result1,result2,result3,result4
}

func calculate2(a,b int)( result1 int, result2 int, result3 int, result4 float32){
	result1=a+b
	result2= a-b
	result3=a*b
	result4= float32(a) / float32(b)
	return
}

func display3(args ...int){
	fmt.Print("Size--->",len(args))
	for i:=0;i<len(args);i++{
		fmt.Print("item is -------->",args[i])
	}
}