package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

func mineDigitCoin(input string, expectedOutput string) int {
	expectedOutputLen := len(expectedOutput)
	for i := 0; true; i++ {
		hash := md5.New()
		io.WriteString(hash, fmt.Sprintf("%s%d", input, i))
		output := fmt.Sprintf("%x", hash.Sum(nil))
		if output[:expectedOutputLen] == expectedOutput {
			return i
		}
	}
	return 0
}

func main() {
	before := time.Now()
	fmt.Println("Answer for 5 digits is:", mineDigitCoin("bgvyzdsv", "00000"),
		"took:", time.Now().Sub(before))
	before = time.Now()
	fmt.Println("Answer for 6 digits is:", mineDigitCoin("bgvyzdsv", "000000"),
		"took.", time.Now().Sub(before))
}
