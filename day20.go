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

	// Find the lowest housenumber, where the value is higher than goal.
	minKey := -1
	for key, value := range houses {
		if key != 1 && value >= goal && (minKey == -1 || key < minKey) {
			minKey = key
		}
	}
	fmt.Println("Smallest housenumber is:", minKey)

	// Part 2, elfs only deliver for 50 houses at a time.
	before = time.Now()
	giftcount := make([]int, maxIter)
	for elf := 1; elf < maxIter; elf++ {
		counter := 0
		for house := elf; house < maxIter; house += elf {
			giftcount[house] += 11 * elf
			counter++
			if counter == 50 {
				break
			}
		}
	}
	// Find the first house, that satisfies the goal.
	for house, gifts := range giftcount {
		if gifts >= goal {
			fmt.Println("The first housenumber for part2:", house,
				"took:", time.Now().Sub(before))
			break
		}
	}
}

const goal = 33100000
