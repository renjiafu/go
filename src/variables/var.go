package main

import (
	"fmt"
	"reflect"
)

func main() {

	//三种初始化
	var i int
	fmt.Println("第一种初始化方式 --> ", i)

	var isDone, isFinished, isRed = true, false, "no!"
	fmt.Println("第二种初始化方式 --> ", "isDone", isDone, "isFinished", isFinished, "isRed", isRed)

	i2 := 3
	firstName, lastName := "y", "z"
	fmt.Println("第三种初始化方式 --> ", "i2", i2, "firstName", firstName, "lastName", lastName)

	//类型转换
	f := float64(i)
	u := uint(i)
	fmt.Println("类型转换 --> ", "f", reflect.TypeOf(f), f, "u", reflect.TypeOf(u), u)

}

/*
基本数据类型

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

bool

string

complex64 complex128
*/
