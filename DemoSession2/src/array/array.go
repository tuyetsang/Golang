package array
import 
(
	"fmt"
	"sort"
)
func Demo1(){
	var a[5]int
	a[0]=5
	a[1]=-1
	a[2]=2
	a[3]=4
	a[4]=-6

	fmt.Println(a)
	fmt.Println("Size:",len(a))
	for i:=0;i<len(a);i++{
		fmt.Println("Item",i+1,"value",a[i])
	}
}

func Demo2(){
	a:=[]int {5,10,12,33,-1,-12}
	fmt.Println(a)
	fmt.Println("Size:",len(a))
	for i:=0;i<len(a);i++{
		fmt.Println("Item",i+1,"value",a[i])
	}
}

func Demo3(a[]int){
	fmt.Println("Size:",len(a))
	for i:=0;i<len(a);i++{
		fmt.Println("Item",i+1,"value",a[i])
	}
	
}


func Demo4(){
	a:=[]string {"A","B","C"}
	a = append(a,"D")
	fmt.Println("Size:",len(a))
	for i:=0;i<len(a);i++{
		fmt.Println("Item",i+1,"value",a[i])
	}
	
}

func Demo5(a[]int)[]int{
	fmt.Println("Size:",len(a))
	var result[]int
	for i:=0;i<len(a);i++{
		if a[i] >0{
			result = append(result,a[i])
		}
	}
	return result	
}

func Demo6(a[]int)(result1[]int, result2[]int,result3[]int, result4[]int){
	fmt.Println("Size:",len(a))
	for i:=0;i<len(a);i++{
		if a[i] >0{
			result1 = append(result1,a[i])
		}
		if a[i]<0{
			result2 =append(result2,a[i])
		}
		if a[i]%2==0 && a[i]>0{
			result3 =append(result3,a[i])
		}
		if a[i]%2!=0 && a[i]>0{
			result4 =append(result4,a[i])
		}
	}
	return result1,result2,result3,result4	
}

// range will get value and index of item
func Demo7(){
	a:=[]string {"B","C","A"}
	for i,v:=range a {
		fmt.Println("index",i,"Value",v)
	}
	for _,v:=range a {
		fmt.Println("Value",v)
	}
}

func Demo8(){
	a:=[2][3]int{
		{5,-2,9},
		{6,1,8},
	}
	fmt.Println(a)
	fmt.Println("Size:",len(a))

	for r:=0; r<len(a);r++{
		for c:=0;c<len(a[r]);c++{
			fmt.Print("\t", a[r][c])
		}
	}
}

func Demo9(){
	a:=[]int {5,10,12,33,-1,-12}
	sort.Ints(a)
	fmt.Println(a)
	b:=[]string {"A","B","C"}
	sort.Strings(b)
	fmt.Println(b)
}