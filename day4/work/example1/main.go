package main

import "fmt"

// 99乘法表
func mutil() {
	for i:=1; i<=9; i++ {
		for j:=1; j<= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	mutil()
}
