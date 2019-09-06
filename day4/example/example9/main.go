package main

import (
	"fmt"
	"sort"
)

func testIntSort() {

	s1 := []int {2,5,3,1}
	sort.Ints(s1)

	fmt.Println(s1)

}

func testStrings() {
	arr := [...]string {"a", "x", "b", "AAAA", "c"}

	sort.Strings(arr[:])
	fmt.Println(arr)

}

func testIntSearch() {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	sort.Ints(a[:])
	fmt.Println(a)
	index := sort.SearchInts(a[:], 348)
	fmt.Println(index)
}

func main() {
	//testIntSort()
	//testStrings()

	testIntSearch()
}