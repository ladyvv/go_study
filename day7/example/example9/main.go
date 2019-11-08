package main

import (
	"encoding/json"
	"fmt"
)

type stu struct{
	Name string
	Age int
	Score float32
}
func testStruct() {
	//stu1 := &stu {
	//	Name:"李微微",
	//	Age:18,
	//	Score:125,
	//}

	var stu2 stu
	stu2.Name = "李晴"
	stu2.Age = 20
	stu2.Score = 100



	result, err := json.Marshal(stu2)
	ret := string(result)
	fmt.Println(ret, err)
}

func testMap() {
	//var m1  map[string]int

	m1 := make(map[string]interface{})

	m1["Name"] = "ladyvv"
	m1["age"] = 18
	m1["score"]=100.05

	str, err := json.Marshal(m1)

	ret := string(str)
	fmt.Println(ret, err)

	var m2  map[string]interface{}
	err = json.Unmarshal(str, &m2)

	fmt.Println(m2)



}



func main() {
	//testStruct()

	testMap()
}
