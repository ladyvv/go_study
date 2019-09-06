package main

import "fmt"

func test1() {
	a := [5]int {1,2,3,4,5}
	fmt.Println(a)

	//a[1] = 5
	//fmt.Println(a)
	//
	//b := a
	//fmt.Println(b)
	//
	//b[2] = 100
	//fmt.Println(b, a)

	test2(&a)
	fmt.Println(a)

}

func test2(arr *[5]int) {
	(*arr)[0] = 100
}


func main() {
	// 遍历数组
	test1()
}
