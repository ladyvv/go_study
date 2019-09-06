package main

import "fmt"

type student struct {
	name string
	age uint32
	job string
}

func main() {
	var s student

	s.name = "liww"
	s.age = 20
	s.job = " 工程师"

	fmt.Println(s)

	fmt.Printf("Name:%p\n", &s.name)
	fmt.Printf("Name:%p\n", &s.age)
	fmt.Printf("Name:%p\n", &s.job)




}
