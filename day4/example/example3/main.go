package main

import "fmt"

func fabInner(n int) int {
	if n ==1 || n ==2 {
		return 1
	}
	return fabInner(n-1)+fabInner(n-2)
}

// 斐波那契数列
func fab(n int) {

	// 方式一：F(1)=1，F(2)=1, F(n)=F(n-1)+F(n-2)（n>=3，n∈N*）
	for i:=1; i< n; i++ {
		fmt.Println(fabInner(i))
	}

	// 方式二：
	tmpArr := make([]int, n)
	for i:=1; i< n; i++ {
		arrKey := i-1
		if i ==1 || i ==2 {
			tmpArr[arrKey] = 1
			continue
		}
		tmpArr[arrKey] = tmpArr[arrKey-1]+tmpArr[arrKey-2]
	}

	fmt.Println(tmpArr)


}

func main() {


	fab(10)
}