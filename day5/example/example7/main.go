package main

import "fmt"

type student struct {
	name string
	age int
	score int
}

type integer int

func (this integer) print() {
	fmt.Println(this)
}

func (p *integer) set(value integer) {
	*p = value
}

func (this *student) init(name string, age int) {
	this.name = name
	this.age = age
}

func (this student) get() student{
	return this
}

func main() {

	var stu student

	stu.init("李微微", 28)

	fmt.Println(stu)

	stu1 := stu.get()
	fmt.Println(stu1)

	var abc integer = 10
	abc.set(20)
	abc.print()
}
