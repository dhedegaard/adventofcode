package main

import (
	"fmt"
	"time"
)

func main() {
	// Calculate number of packages for all possible solutions.
	fmt.Print("Generating giftcount for houses... ")
	before := time.Now()
	houses := make(map[int]int)
	maxIter := goal / 10
	for house := 1; house < maxIter; house++ {
		for elf := house; elf < maxIter; elf += house {
			houses[elf] += house * 10
		}
	}
	fmt.Println("Took", time.Now().Sub(before))

	minKey := -1
	for key, value := range houses {
		if key != 1 && value >= goal && (minKey == -1 || key < minKey) {
			minKey = key
		}
	}
	fmt.Println("Smallest housenumber is:", minKey)
}

const goal = 33100000
