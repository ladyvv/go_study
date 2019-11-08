package main

import "fmt"



func main() {


	var ch1 chan int
	ch1 = make(chan int, 100)

	//var ch2 chan int

	for i:=0; i<100; i++ {
		ch1 <- i
	}

	close(ch1)

	var a int

	for i:=0; i<100; i++ {
		a = <- ch1
		fmt.Println(a)
	}

}
