package main

import "fmt"

func a1() {
	arr := []int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		defer fmt.Print(arr[i], " ")
	}
}

func main() {
	a1()
}
