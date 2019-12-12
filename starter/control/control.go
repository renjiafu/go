package main

import (
	"time"
	"runtime"
	"math"
	"fmt"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow (x,n,lim float64) float64{
	if v:=math.Pow(x,n);v<lim{
		return v
	}else{
		fmt.Printf("%g >= %g",v,lim)
	}
	return lim
}

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	for sum < 10000 {
		sum += sum
	}
	fmt.Println(sum)

	/*
		for{
			fmt.Println("死循环!")
		}
	*/

	fmt.Println(sqrt(2),sqrt(-4))

	fmt.Println(pow(3,3,20))

	switch os:= runtime.GOOS;os{
	case "sarwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default :
	fmt.Printf("%s.",os)
	}
	ti:=time.Now()
	switch {
	case ti.Hour()<12:
		fmt.Println("Good morning!")
	case ti.Hour()<17:
		fmt.Println("Good afternoon.")
	default :
		fmt.Println("Good evening")

	}

/* 	
	延迟计算
	defer fmt.Println("world!")
	fmt.Println("hello ")

 */	
 	fmt.Println("counting")
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")


}
