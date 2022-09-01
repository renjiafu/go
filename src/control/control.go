package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// 循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//for true {
	//	fmt.Println("死循环!")
	//}
	//
	//for {
	//	fmt.Println("死循环!")
	//}

	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("windows")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")
	}

	fmt.Println("延迟调用 -->")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}
