package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func del(x, y int) int {
	return x - y
}

//多返回
func swap(x, y string) (string, string) {
	return y, x
}

// go语言的返回值可以命名，如果这样它们被视为变量，定义在函数首部。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

	fmt.Println(add(23, 32))

	fmt.Println(del(100, 99))

	fmt.Println(swap("world", "hello "))

	fmt.Println(split(10))

}
