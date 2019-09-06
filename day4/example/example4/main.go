package main

import "fmt"

//
func add() func(a int) int {

	x := 0
	f := func (a int) int {
		x += a
		return x
	}

	return f
}
func main() {

	f := add()
	fmt.Println(f)


	fmt.Println(f(1))
	fmt.Println(f(10))
	fmt.Println(f(100))


}
