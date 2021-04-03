package main

import (
	"fmt"
)

func main() {
	input := 1277271272771
	odds := 0
	evens := 0

	for {
		num := input % 10
		input = input / 10

		if num%2 == 0 {
			evens += 1
		} else {
			odds += 1
		}
		if input == 0 {
			break
		}
	}
	if odds > evens {
		fmt.Printf("Odd (%d odd digits, %d even digits)\n", odds, evens)
	} else {
		fmt.Println("Even  (%d odd digits, %d even digits)\n", odds, evens)
	}
}
