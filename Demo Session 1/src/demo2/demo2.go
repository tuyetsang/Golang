package main

import "fmt"

func main(){
	fmt.Println("[1] Tinh tong cac so tu a den b")
	fmt.Println("[2] Dem cac so chan tu a den b")
	fmt.Println("[3] Tinh tong cac so tu a den b")
	choosen:= 0 
	fmt.Scanf("%d\n",&choosen)
	switch choosen {
		case 1:
			result:= handle1()
			fmt.Println("Tong cac so la",result)
		case 2:
			result:= hanble2()
			fmt.Println("So chan la",result)
		case 3:
			choose:=true
			for{
				choose =handle3()
				if choose==false{break}
			}
		default:
			fmt.Println("Invalid")
	}
}

func handle1() int {
	a:=0
	b:=0
	fmt.Println("Nhap vao so a")
	fmt.Scanf("%d\n",&a)
	fmt.Println("Nhap vao so b")
	fmt.Scanf("%d\n",&b)
	result:=0
	for i:=a;i<b;i++{
		result = result + i
	}
	return result
}

func hanble2() int {
	a:=0
	b:=0
	fmt.Println("Nhap vao so a")
	fmt.Scanf("%d\n",&a)
	fmt.Println("Nhap vao so b")
	fmt.Scanf("%d\n",&b)
	counter:=0
	for i:=a;i<b;i++{
		if i%2==0{
			counter = counter + 1
		}	
	}
	return counter
}

func handle3() bool {
	a:=0
	b:=0
	fmt.Println("Nhap vao khoang nho nhat a")
	fmt.Scanf("%d\n",&a)
	fmt.Println("Nhap vao khoang lon nhat b")
	fmt.Scanf("%d\n",&b)
	var x int
	fmt.Println("Nhap vao so bat ky")
	fmt.Scanf("%d\n",&x)
	if x>a && x<b{
		fmt.Println("x nam trong khoang vua nhap vao")
	}
	var choose string
	fmt.Println("Ban co muon tiep tuc khong")
	fmt.Scanf("%s\n",&choose)
	if choose=="Y"{
		return true
	} else {
		return false
	}
}

