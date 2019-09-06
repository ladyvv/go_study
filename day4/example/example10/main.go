package main

import "fmt"

func testMap() {
	//var a map[string]string = map[string]string{
	//	"key": "value",
	//}
	////a := make(map[string]string, 10)
	//a["abc"] = "efg"
	//a["abc"] = "efg"
	//a["abc1"] = "efg"
	//
	//fmt.Println(a)
	//



	a := map[int]string {1:"name", 2:"title"}

	a[3] = "age"
	fmt.Println(a)
}

func modify(a map[int]string) {
	a[2] = "30"
}

func modify2(a []int) {
	a[2] = 30
}

func testMap2() {
	a := make(map[int]map[string]string , 2)

	a[1] = make(map[string]string, 2)
	a[1]["name"] = "liww"
	a[1]["age"] ="18"
	fmt.Println(a)
}

func testMap3() {
	a := make(map[int]string, 2)
	a[1] = "liww"
	a[2] ="18"

	modify(a)
	fmt.Println(a)

	b := []int{1,2,3,4}
	modify2(b)
	fmt.Println(b)
}

func main() {
	//testMap()
	//testMap2()
	testMap3()

}
