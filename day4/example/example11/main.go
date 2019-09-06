package main

import (
	"fmt"
	"sort"
)

func testMapSort() {
	m1 := make(map[int]int, 5)
	m1[1] =1
	m1[3] =3
	m1[0] = 0
	m1[4] =4
	m1[2] =2


	var key  []int

	for k := range m1 {
		key = append(key, k)
	}

	sort.Ints(key)

	for _, v := range key {
		fmt.Println(m1[v])
	}
	fmt.Println(key)


}


func main() {
	testMapSort()
}
