package main

import "fmt"

type cart1 struct {
	name string
	nom int
}

type cart2 struct {
	name string
	nom int
}


type train struct {
	cart1
	cart2
	cartnum int
}

func main() {

	var t  = train{

	}

	t.cart1.name = "车厢"
	t.cart1.nom = 1
	t.cartnum=5

	fmt.Println(t)


}