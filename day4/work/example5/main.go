package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 两个大数相加
func add(s1 string, s2 string) string {

	minL := len(s1)
	maxL := len(s2)
	longS := s2
	if len(s2) < len(s1) {
		minL = len(s2)
		maxL = len(s1)
		longS = s1
	}

	sum := ""
	left :=  0
	diff := 0
	s1t := 0
	s2t := 0

	for i:=0; i<minL; i++ {
		fmt.Println(minL-i-1)
		s1t = int(s1[maxL-i-1] - '0')
		s2t = int(s2[minL-i-1] - '0')
		addret := s2t + s1t + left

		if addret >= 10 {
			left = 1
			diff = addret-10
		} else {
			left = 0
			diff = addret
		}
		diffNew := diff + '0'
		sum = fmt.Sprintf("%c", diffNew) + sum
	}


	for i:=maxL-minL-1; i>=0; i-- {
		fmt.Println(i)
		s1t = int(longS[i] - '0')
		addret := s1t + left

		if addret >= 10 {
			left = 1
			diff = addret-10
		} else {
			left = 0
			diff = addret
		}
		diffNew := diff + '0'
		sum = fmt.Sprintf("%c", diffNew) + sum
	}


	return sum
}


func main() {

	reader := bufio.NewReader(os.Stdin)

	result, _, _ := reader.ReadLine()

	resultNew := strings.Split(string(result), "+")

	addResult := add(resultNew[0], resultNew[1])
	fmt.Println(addResult)

}
