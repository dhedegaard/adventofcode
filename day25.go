package main

import (
	"fmt"
	"time"
)

// Returns a function that iterates on the value.
func calc() func() int {
	value := 20151125
	return func() int {
		value = ((value * 252533) % 33554393)
		return value
	}
}

func main() {
	iter := calc()
	row := 1
	col := 1
	var value int
	before := time.Now()
	for {
		if row == 1 {
			// Check to see if we reset columns and iterate rows.
			row, col = col+1, 1
		} else {
			col++
			row--
		}
		value = iter()
		if row == inputRow && col == inputCol {
			break
		}
	}
	fmt.Println("Value at row", inputRow, "and col", inputCol, "is:",
		value, "took:", time.Now().Sub(before))
}

const (
	inputRow = 2981
	inputCol = 3075
)
