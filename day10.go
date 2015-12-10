package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func lookAndSay(input string) string {
	var currentNumber rune
	var count int
	var result bytes.Buffer

	for _, char := range input {
		if char == currentNumber {
			// If we encounter the same char again, increment the counter.
			count++
		} else {
			// If the char is new, check to see if the currentNumber is the
			// initial value.
			if currentNumber != '\x00' {
				result.WriteString(strconv.Itoa(count))
				result.WriteRune(currentNumber)

			}
			// Set new current number, and set counter to 1.
			currentNumber = char
			count = 1
		}
	}
	// If currentNumber is defined after we're done iterating, add the rest to
	// the result.
	if currentNumber != '\x00' {
		result.WriteString(strconv.Itoa(count))
		result.WriteRune(currentNumber)
	}
	return result.String()
}

func main() {
	input := "3113322113"

	before := time.Now()
	for i := 1; i <= 50; i++ {
		input = lookAndSay(input)
		// When we've iterated 40 times, print the length of the output for
		// solving part1.
		if i == 40 {
			fmt.Println("After 40 iterations:", len(input), "took:",
				time.Now().Sub(before))
		}
	}
	// After 50 iterations we have the result for part2.
	fmt.Println("After 50 iterations:", len(input), "took:",
		time.Now().Sub(before))
}
