package main

import "fmt"

type student struct {
	Name string
	height int
	score float32
}
func main() {

	str := "ladyvv 160 89.5"
	var stu student

	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.height, &stu.score)
	fmt.Println(stu)
}
