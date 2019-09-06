package main

import (
	"fmt"
	"strconv"
)

type student struct {
	name string
	age int
	num string
	next *student
}
func main() {

	var head student

	head.name ="李班长"
	head.age = 20
	head.num="no.1"

	inserttail(&head)
	//fmt.Println(head)

	//inserthead(&head)

	//bianli(&head)

	bianli(&head)

	//i := 10
	//str := strconv.Itoa(i)
	//
	//fmt.Println(str)
	//
	//str1 := fmt.Sprintf("stu%d", i)
	//fmt.Println(str1)



	del(&head)
	bianli(&head)

	add(&head)
	bianli(&head)



}

func del(p *student) {

	var prev *student = new(student)

	for p != nil {
		if p.name == "stu5" {
			prev.next = p.next
		}
		prev = p
		p = p.next
	}
}

func add(p *student) {

	var news student = student{
		name:"李微微",
		age:31,
	}

	for p != nil {
		if p.name == "stu1" {
			news.next=p.next
			p.next = &news
		}
		p = p.next
	}
}

func inserthead(p *student) {
	for i:=0; i<10; i++ {
		var stu = student{
			name:"stu"+strconv.Itoa(i),
			age : i,
			num : "no."+strconv.Itoa(i),
			next:p,
		}
		p = &stu
	}

	bianli(p)
}


func inserttail(p *student) {
	for i:=0; i< 10; i++ {
		var stu = student{
			name:"stu"+strconv.Itoa(i),
			age : i,
			num : "no."+strconv.Itoa(i),
		}
		p.next = &stu
		p = p.next
	}
}

func bianli(p *student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}
