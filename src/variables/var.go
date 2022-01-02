package main

import "fmt"

func main() {

	//三种初始化
	var i int
	fmt.Println(i)

	var k, v, s = true, false, "no!"
	fmt.Println(i, k, v, s)

	x := 3
	y, z := "y", "z"
	fmt.Println(x, y, z)

	//默认初始化
	var f float64
	var b bool
	var str string
	fmt.Printf("%v %v %q\n", f, b, str)

	//类型转换
	o := 33
	d := float64(o)
	u := uint(d)
	fmt.Println(o, d, u)

}

/*
基本数据类型
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
