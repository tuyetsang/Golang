package demo_date

import (
	"fmt"
	"time"
)

func Demo1() {
	today := time.Now()
	fmt.Println("today:", today)
	fmt.Println("Year:", today.Year())
	month := today.Month()
	fmt.Println(month)
	fmt.Println(int(month))
	fmt.Println("day:", today.Day())
	fmt.Println("day:", today.Hour())
	fmt.Println("day:", today.Minute())
}
func Demo2() {
	today := time.Now()
	fmt.Println(today.Format("02/01/2006 15:04:05"))
}
func Demo3() {
	s := "27/12/2018"
	d, error := time.Parse("02/01/2006", s)
	if error != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(d.Format("2006 - 01 - 02"))
	}
}
func Demo4() {
	today := time.Now()
	time1 := today.AddDate(0, 0, 2)
	fmt.Println("time1:", time1.Format("02/01/2006"))
	time2 := today.Add(2 * 24 * time.Hour)
	fmt.Println("time2:", time2.Format("02/01/2006"))
}

func Demo5() {
	ngaysinh := "27/12/2018"
	today := time.Now()
	d, error := time.Parse("02/01/2006", ngaysinh)
	d.Format("02/01/2006")
	if error != nil {
		fmt.Println("Error")
	} else {
		if d.Day() == today.Day() && d.Month() == today.Month() && d.Year() == today.Year() {
			fmt.Println("Hom nay la sinh nhat")
		} else {
			fmt.Println("Hom nay khong phai la sinh nhat")
		}
	}

}

func Demo6() {
	var tuoi int
	ngaysinh := "27/12/2000"
	today := time.Now()
	d, error := time.Parse("02/01/2006", ngaysinh)
	d.Format("02/01/2006")
	if error != nil {
		fmt.Println("Error")
	} else {
		tuoi = today.Year() - d.Year()
		fmt.Println("tuoi:", tuoi)
	}
}

/*
cho ngay sinh kieu chuoi.
1.kiem tra hom nay phai laf sinh nhat k?
2.Khai bao ham tinh tuoi dua vao ngay sinh duoc cung cap
*/
