package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gcd(top, bottom int) int {
	// copy pasta : https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations
	for bottom != 0 {
		x := bottom
		bottom = top % bottom
		top = x
	}
	return top
}

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(seed)
	for i := 0; i < 20; i++ {
		num1, num2 := generator.Intn(2000), generator.Intn(2000)
		areCoprime := gcd(num1, num2) == 1
		var val string
		if areCoprime {
			val = "✅"
		} else {
			val = "❌"
		}
		fmt.Printf("%d and %d are %s coprime\n", num1, num2, val)
	}

}
