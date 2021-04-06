package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle(slice []int) {
	seed := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(seed)
	fmt.Println(slice)
	for i := range slice {
		r := generator.Intn(len(slice))
		slice[i], slice[r] = slice[r], slice[i]
	}
}

func weave(slice1, slice2 []int) []int {
	combined := make([]int, len(slice1)+len(slice2))
	var (
		shorter []int
		longer  []int
	)
	if len(slice1) > len(slice2) {
		shorter = slice2
		longer = slice1
	} else {
		shorter = slice1
		longer = slice2
	}
	var i int
	for i = 0; i < len(shorter); i++ {
		combined[i*2] = shorter[i]
		combined[i*2+1] = longer[i]
	}
	// append the rest of the longer slice to the combined slice
	for i = 0; i < len(longer)-len(shorter); i++ {
		combined[i+len(shorter)*2] = longer[len(shorter)+i]
	}
	return combined
}

func main() {
	slice0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	shuffle(slice0)
	fmt.Println(slice0)
	fmt.Println("Weaved: ", weave(slice1, slice2))
}
