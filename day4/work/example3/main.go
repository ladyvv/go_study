package main

import "fmt"

// 回文 aba格式
func huiwen(str string) bool {
	arr := []rune(str)
	length := len(arr)

	for i:=0; i< length; i++ {
		if i == length/2 {
			return true
		}
		if arr[i] != arr[length-1-i] {
			return false
		}
	}
}

func main() {
	var str string

	fmt.Scanf("%s", &str)

	if huiwen(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
