package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ingredient struct {
	name       string
	capacity   int64
	durability int64
	flavor     int64
	texture    int64
	calories   int64
}

func strToInt(input string) int64 {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return int64(result)
}

func main() {
	// Parse the ingredients.
	reInput := regexp.MustCompile(
		`(.*?): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)
	var ingredients []ingredient
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		rc := reInput.FindStringSubmatch(scanner.Text())
		ingredients = append(ingredients, ingredient{
			name:       rc[1],
			capacity:   strToInt(rc[2]),
			durability: strToInt(rc[3]),
			flavor:     strToInt(rc[4]),
			texture:    strToInt(rc[5]),
			calories:   strToInt(rc[6]),
		})
	}

	before := time.Now()
	var maxScore, maxScoreWithCalories, a, b, c int64
	for a = 0; a <= 100; a++ {
		for b = 0; b <= 100-a; b++ {
			for c = 0; c <= 100-a-b; c++ {
				d := 100 - a - b - c
				capacity := ingredients[0].capacity*a +
					ingredients[1].capacity*b +
					ingredients[2].capacity*c +
					ingredients[3].capacity*d
				durability := ingredients[0].durability*a +
					ingredients[1].durability*b +
					ingredients[2].durability*c +
					ingredients[3].durability*d
				flavor := ingredients[0].flavor*a +
					ingredients[1].flavor*b +
					ingredients[2].flavor*c +
					ingredients[3].flavor*d
				texture := ingredients[0].texture*a +
					ingredients[1].texture*b +
					ingredients[2].texture*c +
					ingredients[3].texture*d
				if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
					continue
				}

				score := capacity * durability * flavor * texture
				if score > maxScore {
					maxScore = score
				}

				if score > maxScoreWithCalories &&
					(ingredients[0].calories*a+
						ingredients[1].calories*b+
						ingredients[2].calories*c+
						ingredients[3].calories*d) == 500 {
					maxScoreWithCalories = score
				}
			}
		}
	}
	fmt.Println("The best cookie:", maxScore)
	fmt.Println("The best cookie with 500 calories:", maxScoreWithCalories)
	fmt.Println("took:", time.Now().Sub(before))
}

const input = `Sprinkles: capacity 2, durability 0, flavor -2, texture 0, calories 3
Butterscotch: capacity 0, durability 5, flavor -3, texture 0, calories 3
Chocolate: capacity 0, durability 0, flavor 5, texture -1, calories 8
Candy: capacity 0, durability -1, flavor 0, texture 5, calories 8`
