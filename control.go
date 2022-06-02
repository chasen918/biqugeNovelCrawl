package main

import "fmt"

//重构中
func getInput() int {
	var num int
	fmt.Scanln(&num)
	return num
}
