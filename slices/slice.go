package main

import (
	"fmt"
	"strings"
)

func main() {

	primes := [6]int{2, 3, 1, 5, 8, 0}

	// 切片左闭右开
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)

	str := []int{2, 3, 5, 7, 10}
	printSlice(str)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	//nil slices
	var s1 []int
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 5)
	printSlice(s2)

	s3 := make([]int, 0, 5)
	printSlice(s3)

	s4 := s3[:2]
	printSlice(s4)

	s5 := s4[2:5]
	printSlice(s5)

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
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}

	var s6 []int
	printSlice(s6)

	s6 = append(s6, 1)
	printSlice(s6)

	s6 = append(s6, 2, 3, 4)
	printSlice(s6)

}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d , %v \n", len(s), cap(s), s)
}
