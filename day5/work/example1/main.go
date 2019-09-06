package main

import "fmt"

func maopao(s1 []int, maxIndex int) {

	for i:=0; i<= maxIndex; i++ {
		if i == len(s1)-1 {
			break
		}
		if s1[i] <= s1[i+1] {
			continue
		}
		s1[i], s1[i+1] = s1[i+1], s1[i]
		maopao(s1, i)
	}
}

func sort1() {
	s1 := []int {5,4,3,2,1}
	fmt.Println(s1)
	maopao(s1, len(s1)-1)
	fmt.Println(s1)
}


// 冒泡排序
func bubblesort(s1 []int) {
	for i:=0; i< len(s1); i++ {
		for j:=1; j<len(s1)-i; j++ {
			if s1[j-1] > s1[j] {
				s1[j-1], s1[j] = s1[j], s1[j-1]
			}
		}
	}
}

// 插入排序
func innsertsort(s1 []int) {
	for i := 1; i< len(s1); i++ {
		for j:=i; j>0; j-- {
			if s1[j] < s1[j-1] {
				s1[j], s1[j-1] = s1[j-1], s1[j]
			}
		}
	}
}

//
func choosesort(s1 []int) {
	for i := 0; i< len(s1); i++ {
		min := i
		for j:=i+1; j < len(s1); j++ {
			if s1[min] > s1[j] {
				min = j
			}
		}
		s1[min], s1[i] = s1[i], s1[min]
	}
}

func main() {
	//sort1()

	s1 := []int {5,4,3,2,1,6}
	s2 := []int{12, 15, 9, 20, 6, 31, 24}

	//bubblesort(s1)
	//fmt.Println(s1)

	//innsertsort(s2)
	//fmt.Println(s2)

	choosesort(s1)
	fmt.Println(s1)


	choosesort(s2)
	fmt.Println(s2)
}
