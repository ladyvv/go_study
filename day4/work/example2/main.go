package main

import "fmt"

func wanshuInner(n int) bool {
	sum := 0
	for j:=1; j< n; j++ {
		if n%j == 0 {
			sum += j
		}
	}

	return sum == n

}
// 完数 如果一个数恰好等于它的因子之和，则称该数为“完全数”
func wanshu() {
	n := 100
	for i:=1; i< n; i++ {
		if wanshuInner(i) {
			fmt.Println(i)
		}
	}
}
func main() {
	wanshu()
}
