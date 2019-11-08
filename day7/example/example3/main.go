package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	str, error := reader.ReadString('\n')

	fmt.Println(str, error)
}
