package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	var s1 []int
	fmt.Println("切片初始化 --> ", "s1", s1, "type", reflect.TypeOf(s1), "len", len(s1), "cap", cap(s1))

	s2 := make([]int, 5)
	fmt.Println("切片初始化 --> ", "s2", s2, "type", reflect.TypeOf(s2), "len", len(s2), "cap", cap(s2))

	s3 := make([]int, 0, 5)
	fmt.Println("切片初始化 --> ", "s3", s3, "type", reflect.TypeOf(s3), "len", len(s3), "cap", cap(s3))

	s4 := s3[:2]
	fmt.Println("切片截取 --> ", "s4", s4, "type", reflect.TypeOf(s4), "len", len(s4), "cap", cap(s4))

	s5 := s4[2:5]
	fmt.Println("切片截取 --> ", "s5", s5, "type", reflect.TypeOf(s5), "len", len(s5), "cap", cap(s5))

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Println("type", reflect.TypeOf(board))
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}

	var s6 []int
	fmt.Println("切片初始化 --> ", "s6", s6, "type", reflect.TypeOf(s6), "len", len(s6), "cap", cap(s6))

	s6 = append(s6, 1)

	s6 = append(s6, 2, 3, 4)
	fmt.Println("s6 -->", s6)
}
