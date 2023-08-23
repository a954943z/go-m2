package aaaa

import "fmt"

func Forfor() {

	// 同while迴圈

	a := 0
	for a < 4 {
		fmt.Println(a)
		a++
	}
println("------------")
	b := 1
	for b <= 3{
		fmt.Println(b)
		b++
	}
	println("------------")
	i := 0
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i++
	}
	println("------------")
	//一般的 for迴圈
	for x := 0; x < 3; i++ {
		fmt.Println(x)
		x++
	}
	println("------------")
	//
	bucky := []int{1, 2, 3}
	for a, v := range bucky {
		fmt.Printf("a: %v,bucky: %v", a, v)
	}

}
