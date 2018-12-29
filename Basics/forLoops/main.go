package main

import "fmt"

func main() {
	i := 0
	isLessThanFive := true
	for isLessThanFive {
		if i >= 5 {
			isLessThanFive = false
		}
		fmt.Print(i)
		i++
	}
}
