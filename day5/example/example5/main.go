package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"student_name"`
	Age int `json:"student_age"`
	Score int `json:"student_score"`
}
func main() {

	var stu1 = Student{
		Name:"stu1",
		Age :20,
		Score:100,
	}

	arr, error := json.Marshal(stu1)
	//{"student_name":"stu1","student_age":20,"student_score":100} <nil>

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(string(arr), error)
}
