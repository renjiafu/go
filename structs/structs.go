package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {

	fmt.Println(Student{"xiaoming", 12})

	v := Student{"xiaohua", 11}
	v.name = "xiaoli"
	fmt.Println(v)

	p := &v
	p.name = "xiaohong"
	fmt.Println(v)
	fmt.Println(p)
	fmt.Println(*p)

}
