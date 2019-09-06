package main

import "fmt"

func testslice1() {

	var s1 []int
	arr := [5]int {1,2,3,4,5}

	s1 = arr[1:2]
	s2 := arr[1:]
	s3 := s2[3:]
	//
	fmt.Println(s1, s2, s3)
	//
	//fmt.Println(len(s2), cap(s3))

	fmt.Printf("%p", s2)
	fmt.Printf("%p", &arr[1])

}
func main() {
	testslice1()

}
