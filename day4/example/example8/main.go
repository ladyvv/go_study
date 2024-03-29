package main

import "fmt"

func testSlice() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]
	fmt.Printf("before len[%d] cap[%d]\n", len(s), cap(s))
	s[1] = 100
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	fmt.Println("before a:", a)

	s = append(s, 10)
	s = append(s, 10)
	fmt.Printf("after len[%d] cap[%d]\n", len(s), cap(s))
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)

	s[1] = 1000
	fmt.Println("after a:", a)
	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
}

func testCopy() {


	a := []int {1,2,3,4,5}

	b := make([]int, 3)

	copy(b, a)

	fmt.Println(b)
}

func testModifyString() {
	s := "我hello world"
	s1 := []rune(s)
	fmt.Println(s1)
	s1[0] = 200
	s1[1] = 128
	s1[2] = 64
	str := string(s1)
	fmt.Println(str)
}


func main() {
	//testSlice()

	//testCopy()

	testModifyString()

}
