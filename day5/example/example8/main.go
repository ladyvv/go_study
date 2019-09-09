package main

import "fmt"

type car struct {
	weight int
	name string
}

func (p *car) run() {
	fmt.Println("runngin")
}

type bike struct {
	car
	lunzi int
}

func main() {

	var a bike

	a.weight = 100
	a.name = "飞哥"
	a.lunzi = 2

	fmt.Println(a)

	a.run()
}
