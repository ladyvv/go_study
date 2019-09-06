package main

import (
	"errors"
	"fmt"
	"time"
)

// 异常处理
func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}

func testerr() {

	defer func() {
		err :=recover()
		if err != nil {
			fmt.Println(err)
		}
	}()


	err := initConfig()
	if err != nil {
		panic(err)
	}
}

func main() {


	for {
		test()
		time.Sleep(time.Second)
	}

	s1 := []int {1,2,3,4}

	s1 = append(s1, 5,6,7)
	fmt.Println(s1)

	s1 = append(s1, s1...)
	fmt.Println(s1)

}