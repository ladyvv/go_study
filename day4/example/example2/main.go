package main

import "fmt"

func test() {

	//var s1 []int
	//fmt.Println(s1)

	//s2 := new([]int)
	//fmt.Println(s2)
	//
	s3 := make([]int, 5)
	fmt.Println(s3)

	s3[2] = 5
	fmt.Println(s3)
	//
	//s4 := []int {1,2,3,4}
	//fmt.Println(s4)

	//*s2 = make([]int, 4)
	//fmt.Println(s2)

	//*s2 =[]int {1,2,3,4}
	//fmt.Println(s2)
	//
	//fmt.Println(*s2)

}

func main() {
	test()
}
