package main

import "fmt"

func main() {

	var a1 [2]string
	a1[0] = "Hello"
	a1[1] = "World"
	fmt.Println("数组不带元素初始化 --> ", "a1", a1)

	a2 := [8]int{2, 3, 4, 6, 8, 9, 0, 10}
	fmt.Println("数组带元素初始化 --> ", "a2", a2)
}
