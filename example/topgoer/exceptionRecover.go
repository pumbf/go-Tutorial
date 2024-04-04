package main

import "fmt"

func test() {

	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		recover() //无效！
		println("defer inner")
	}()

	panic("test panic")
}

func main() {
	test()
}
