package main

import "fmt"

type stu1 struct {
	name string
	age int
	score int
}

type inter1 interface {
	print()
	sleep()
}
type inter2 interface {
	p1()
}

func (p stu1) print() {
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.score)
}

func (p stu1) sleep() {
	fmt.Printf("%s is sleeping", p.name)
}
func (p stu1) p1() {
	fmt.Printf("%s is p1", p.name)
}


func main() {

	var t1 inter1
	var t2 inter2
	//t1.print()


	var a stu1

	a.name = "李微微"
	a.age = 30
	a.score = 100

	t1 = a
	t2 = a
	t2.p1()
	t1.print()
	//
	//t1.sleep()





}
