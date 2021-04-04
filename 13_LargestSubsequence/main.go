package main

import (
	"fmt"
)

func split(input int) []int {
	var digits []int
	for {
		digit := input % 10
		input /= 10
		digits = append(digits, digit)
		if input == 0 {
			break
		}
	}

	// https://github.com/golang/go/wiki/SliceTricks#reversing
	for i := len(digits)/2 - 1; i >= 0; i-- {
		opp := len(digits) - i - 1
		digits[i], digits[opp] = digits[opp], digits[i]
	}
	return digits
}

func merge(digits []int) int {
	num := 0
	for _, digit := range digits {
		num *= 10
		num += digit
	}
	return num
}

func seek(digits []int, length int) int {
	if len(digits) < length {
		panic("Not enough digits")
	}
	best := 0
	for i := 0; i < len(digits)-length+1; i++ {
		if merge(digits[i:i+length]) > best {
			best = merge(digits[i : i+length])
		}
	}
	return best
}

func main() {
	input := 82394729894299315
	fmt.Println(seek(split(input), 4))
}
