package main

import "fmt"

func testarr1() {
	var a1 [3]int
	a2 := [3]int {1, 2,3}
	var a3  = [3]int {4,5,6}


	fmt.Println(a1, a2, a3)
}

func testarr2() {
	var a1 [3][2]int

	a2 := [3][2]int {{1,2},{11,22},{33,34}}

	fmt.Println(a1, a2)

	for index, value := range a2 {
		for k, v := range value {
			fmt.Printf("index=%d key=%d value=%d \n", index, k, v)
		}
	}

}

func main() {
	//testarr1()

	testarr2()
}
