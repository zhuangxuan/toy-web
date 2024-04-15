package main

import "fmt"

func main() {

	fmt.Println("start")
	defer func() {
		fmt.Println("aaa")
	}()

	defer func() {
		fmt.Println("bbb")
	}()

	defer func() {
		fmt.Println("ccc")
	}()
	fmt.Println("end")
}
