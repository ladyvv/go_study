package main

import "fmt"

type che interface {
	getName() string
	run()
	didi()
}

type bmw struct {
	name string
	bowl int
}

func (this bmw) run() {
	fmt.Println("bmw is running")
}

func (this bmw) didi() {
	fmt.Println("bmw is DIDI")
}

func (this bmw) getName()  string{
	return this.name
}

func main() {

	var bmw1 bmw
	var che1 che

	bmw1.name = "bmw1"
	che1 = bmw1

	che1.run()
	che1.didi()
	name := che1.getName()
	fmt.Println(name)


}
