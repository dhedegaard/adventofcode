package main

import (
	"fmt"
	"time"
)

func permutateContainers(sum int, count int, limit int, containers []int,
	solutions *map[int]int) int {

	if sum == limit {
		// Register the number of containers used in the solution.
		sol := *solutions
		sol[count]++

		// Return a solution.
		return 1
	}

	// We're above the limit we want, or there are no more containers left.
	if sum > limit || len(containers) == 0 {
		return 0
	}

	result := 0
	result += permutateContainers(sum+containers[0], count+1, limit, containers[1:], solutions)
	result += permutateContainers(sum, count, limit, containers[1:], solutions)
	return result
}

func main() {
	before := time.Now()

	// Map for keeping the solutions.
	solutions := make(map[int]int)
	numberOfSolutions := permutateContainers(0, 0, limit, input, &solutions)
	fmt.Println("Total number of unique solutions:", numberOfSolutions)

	// Find the lowest number of containers.
	minKey := 9999999
	count := 0
	for key, value := range solutions {
		if key < minKey {
			minKey = key
			count = value
		}
	}
	fmt.Println("Minimum number of containers used is", count, "with a total number of", count, "permutations")

	// All done, register time.
	fmt.Println("Took:", time.Now().Sub(before))
}

const limit = 150

var input = []int{
	11,
	30,
	47,
	31,
	32,
	36,
	3,
	1,
	5,
	3,
	32,
	36,
	15,
	11,
	46,
	26,
	28,
	1,
	19,
	3,
}
