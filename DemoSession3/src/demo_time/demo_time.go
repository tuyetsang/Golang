package demo_time

import (
	"time"
	"fmt"
)

func Demo1(){
	today := time.Now()
	fmt.Println("today--->",today)
	year := today.Year()
	month := today.Month()
	fmt.Println("year--->",year)
	fmt.Println("month--->",month)
	fmt.Println("month--->",int(month))
	day := today.Day()
	fmt.Println("day--->",int(day))
}
func Demo2(){
	today := time.Now()
	fmt.Println("today----->",today.Format("02/01/2006 15:04:05"))
}

func Demo3(){
	s := "20/10/2019"
	value,err := time.Parse("02/01/2006",s)
	if err != nil {
		fmt.Println(err)
	}else
	{
		fmt.Println(value.Format("2006-01-02"))
	}
}

func Demo4(){
	date1,_ := time.Parse("02/01/2006","10/10/2018")
	date2,_ := time.Parse("02/01/2006","15/10/2018")
	result := date2.Sub(date1).Hours() / 24
	fmt.Println("result---->",result)
	result2 := date2.AddDate(0,0,3)
	fmt.Println("result2----->",result2)
}