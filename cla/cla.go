package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	arguments := os.Args
	for i := 0; i < len(arguments); i++ {
		temp, _ := strconv.Atoi(arguments[i])
		sum += temp
	}
	fmt.Println("Sum: ", sum)
}
