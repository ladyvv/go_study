package main

import (
	"bufio"
	"fmt"
	"os"
)

// 统计字母数量，数字数量，空格数量，其他数量
func stat(str string) (betterNum , noNum, spaceNum, otherNum int){
	arr := []rune(str)
	//betterNum, noNum, spaceNum, otherNum = 0

	for _,v := range arr {
		fmt.Println(v)
		if v >= 'a' && v <= 'z' {
			betterNum++
		} else if v >= '0' && v <= '9' {
			noNum++
		} else  if v == ' ' {
			spaceNum++
		} else {
			otherNum++
		}

	}

	return
}

func main() {

	//var str string
	//fmt.Scanf("%s\n", &str)

	reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()

	//fmt.Println(str, isPrefix, err)

	betterNum , noNum, spaceNum, otherNum := stat(string(str))
	//
	fmt.Printf("betterNum=%d noNum=%d spaceNum=%d otherNum=%d ", betterNum , noNum, spaceNum, otherNum)
}
