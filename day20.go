package main

import (
	"fmt"
	"time"
)

func main() {
	before := time.Now()
	for house := 1; true; house++ {
		result := 0
		for elf := 1; elf <= house; elf++ {
			// If the house was visited by the elf, calculate and add to the
			// score.
			if house%elf == 0 {
				result += 10 * (house / elf)
			}
		}
		if house%100 == 0 {
			fmt.Println("house:", house, "score:", result)
		}
		if result >= 33100000 {
			fmt.Println("First house at or above input:", house, "took:", time.Now().Sub(before))
			break
		}
	}
}
