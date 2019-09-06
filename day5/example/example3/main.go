package main

import "fmt"

type family struct {
	name string
	age int
	familyno string
	left *family
	right *family
}

func main() {
	var root family

	fmt.Println(root)

	var root2 *family = new(family)
	fmt.Println(*root2)

	root.name = "祖父祖母"
	root.age = 70
	root.familyno = "第一家"


	var left1 family
	left1.name = "left1"
	left1.age = 60
	left1.familyno = "第二left家"

	var right1 family
	right1.name = "right1"
	right1.age = 65
	right1.familyno = "第二right家"


	root.left = &left1
	root.right = &right1

	trans(&root)

}

func trans(root *family) {

	if root == nil {
		return
	}
	fmt.Println(root)

	trans(root.left)
	trans(root.right)
}