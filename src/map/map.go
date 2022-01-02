package main

import "fmt"

func main() {
	var citymap map[string]string
	citymap = make(map[string]string)

	citymap["巴黎"] = "法国"
	citymap["北京"] = "中国"
	citymap["佛罗伦萨"] = "意大利"

	for city := range citymap {
		fmt.Println(city, "国家是", citymap[city])
	}

	v, is := citymap["佛罗伦萨"]

	if is {
		fmt.Println("佛罗伦萨的国家是", v)

	} else {
		fmt.Println("佛罗伦萨的国家不存在")
	}
}
