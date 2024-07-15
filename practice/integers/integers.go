package main

import (
	"fmt"
	"log"
)

func main() {
	var i int
	var max int
	var min int

	for {
		fmt.Scan(&i)
		if i == 0 {
			fmt.Println(min, max)
			log.Fatal("zero entered")
		}
		if i > max {
			max = i
		} else {
			min = i
		}
	}
}
