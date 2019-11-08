package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("arg[%d] = %s \n", i , v)
	}
}
