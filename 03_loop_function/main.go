package main

import "fmt"

func printNumbers(from, to int) {
	for i := from; i <= to; i++ {
		fmt.Println(i)
	}
}

func main() {
	printNumbers(1, 10)
}
