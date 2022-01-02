package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Hello, Go !")

	fmt.Println("The time is : ", time.Now())

	fmt.Println("My favorite number is ", rand.Intn(10))

	fmt.Printf("Now you hava %g problems.\n", math.Sqrt(7))

}

// go语言数据类型
// 布尔型
// 数字类型
// 字符串类型
// 派生类型:
// 	指针类型
// 	数组类型
// 	结构化类型
// 	Channel类型
// 	函数类型
// 	切片类型
// 	接口类型
// 	Map类型
