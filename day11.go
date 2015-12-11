package main

import (
	"fmt"
	"time"
)

func isValid(pass []rune) bool {
	overlapCount := 0
	hasIncreasingTriplets := false
	lastWasOverlap := false
	for i, c := range pass {
		// Avoid invalid characters.
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}

		// The rune before the current rune is the same ?
		if !lastWasOverlap && overlapCount < 2 && i > 0 && c == pass[i-1] {
			overlapCount++
			lastWasOverlap = true
		} else {
			lastWasOverlap = false
		}

		// The two previous runes are -1 and -2 of the current rune ?
		if !hasIncreasingTriplets && i >= 2 {
			if int(c) == int(pass[i-1])+1 && int(c) == int(pass[i-2])+2 {
				hasIncreasingTriplets = true
			}
		}
	}

	// Make sure we have 2 overlaps and triplets at least once.
	return overlapCount >= 2 && hasIncreasingTriplets
}

func generatePassword(previous string) string {
	result := []rune(previous)
	inputLen := len(previous)
	idx := inputLen - 1
	for {
		// If current position is 'z', set it to 'a' and go left one.
		if result[idx] == 'z' {
			result[idx] = 'a'
			idx--
			continue
		}

		// Increment current position.
		result[idx]++

		// If we can go right, then go right until we hit the end or
		if idx < inputLen {
			idx = inputLen - 1
		}

		// If the password is valid, return it.
		if isValid(result) {
			return string(result)
		}
	}
}

func main() {
	before := time.Now()
	part1 := generatePassword("hxbxwxba")
	fmt.Println("Part 1:", part1, "took:", time.Now().Sub(before))
	before = time.Now()
	fmt.Println("Part 2:", generatePassword(part1), "took:",
		time.Now().Sub(before))
}
