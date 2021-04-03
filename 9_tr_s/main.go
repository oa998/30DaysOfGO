package main

import (
	"fmt"
)

func main() {
	input := "thhhiss phhhrasssee hass sssommme  repeeeatted   lllettteerrs iin itt"
	letters := []byte(input)
	var output []byte
	output = append(output, letters[0])

	// "i" always starts at 0 in a range. For this reason, this is
	// the same as "for i = 1; i < len(letters)... if input[i-1] == input[i]"
	for i, letter := range letters[1:] {
		if input[i] == letter {
			continue
		}
		output = append(output, letter)
	}
	fmt.Println("input: ", input)
	fmt.Println("output: ", string(output))
}
