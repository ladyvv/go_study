package main

import "fmt"

type studet struct {
	name string "this is name"
	age int "this is age"
}

type stu studet

type interger int

func main() {
	var i int =  2

	var j interger = 100

	j = interger(i)
	fmt.Println(j)




}
