package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	permutation "github.com/fighterlyt/permutation"
)

type relation struct {
	from string
	to   string
}

// Factory method for creating a new relation, without writing too much.
func newRelation(p1 string, p2 string) relation {
	return relation{
		from: p1,
		to:   p2,
	}
}

// For converting the map of person names to a slice of persons.
func keyOfMapToSlice(input map[string]struct{}) []string {
	result := make([]string, 0, len(input))
	for k := range input {
		result = append(result, k)
	}
	return result
}

// Calculate the maximum happiness for a given permutation of persons.
func calculateHappiness(input []string, happinessMap map[relation]int) int {
	inputLen := len(input)
	result := 0
	for i, person := range input {
		beforePerson := input[(inputLen+i-1)%inputLen]
		afterPerson := input[(i+1)%inputLen]
		result += happinessMap[newRelation(person, beforePerson)]
		result += happinessMap[newRelation(person, afterPerson)]
	}
	return result
}

// Creates permutations of the persons involved, returning the maximum
// happiness of any given permutation.
func calculateMaxHappiness(persons []string,
	happinessMap map[relation]int) int {
	perm, err := permutation.NewPerm(persons, nil)
	if err != nil {
		panic(err)
	}
	maxHappiness := 0
	for e, err := perm.Next(); err == nil; e, err = perm.Next() {
		current := e.([]string)
		happiness := calculateHappiness(current, happinessMap)
		if happiness > maxHappiness {
			maxHappiness = happiness
		}
	}
	return maxHappiness
}

func main() {
	happinessMap := make(map[relation]int)
	persons := make(map[string]struct{})

	reMatchInstruction := regexp.MustCompile(`(\S+) would (gain|lose) (\d+) happiness units by sitting next to (\S+).`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		rc := reMatchInstruction.FindStringSubmatch(scanner.Text())[1:]

		// Parse happiness value, for relation.
		happiness, err := strconv.Atoi(rc[2])
		if err != nil {
			panic(err)
		}
		if rc[1] == "lose" {
			happiness *= -1
		}

		// Add relationship to the map.
		rel := newRelation(rc[0], rc[3])
		happinessMap[rel] = happiness

		// Add person to persons, if this is the first time we encounter it.
		if _, exists := persons[rc[0]]; !exists {
			persons[rc[0]] = struct{}{}
		}
	}

	// Prepare slices of persons.
	personsSlice := keyOfMapToSlice(persons)
	persons["self"] = struct{}{}
	personsSliceWithSelf := keyOfMapToSlice(persons)

	before := time.Now()
	fmt.Println("Optimal maximum happiness:",
		calculateMaxHappiness(personsSlice, happinessMap), "took:",
		time.Now().Sub(before))
	before = time.Now()
	fmt.Println("Optimal maximum happiness with self:",
		calculateMaxHappiness(personsSliceWithSelf, happinessMap), "took:",
		time.Now().Sub(before))
}

const input = `Alice would lose 57 happiness units by sitting next to Bob.
Alice would lose 62 happiness units by sitting next to Carol.
Alice would lose 75 happiness units by sitting next to David.
Alice would gain 71 happiness units by sitting next to Eric.
Alice would lose 22 happiness units by sitting next to Frank.
Alice would lose 23 happiness units by sitting next to George.
Alice would lose 76 happiness units by sitting next to Mallory.
Bob would lose 14 happiness units by sitting next to Alice.
Bob would gain 48 happiness units by sitting next to Carol.
Bob would gain 89 happiness units by sitting next to David.
Bob would gain 86 happiness units by sitting next to Eric.
Bob would lose 2 happiness units by sitting next to Frank.
Bob would gain 27 happiness units by sitting next to George.
Bob would gain 19 happiness units by sitting next to Mallory.
Carol would gain 37 happiness units by sitting next to Alice.
Carol would gain 45 happiness units by sitting next to Bob.
Carol would gain 24 happiness units by sitting next to David.
Carol would gain 5 happiness units by sitting next to Eric.
Carol would lose 68 happiness units by sitting next to Frank.
Carol would lose 25 happiness units by sitting next to George.
Carol would gain 30 happiness units by sitting next to Mallory.
David would lose 51 happiness units by sitting next to Alice.
David would gain 34 happiness units by sitting next to Bob.
David would gain 99 happiness units by sitting next to Carol.
David would gain 91 happiness units by sitting next to Eric.
David would lose 38 happiness units by sitting next to Frank.
David would gain 60 happiness units by sitting next to George.
David would lose 63 happiness units by sitting next to Mallory.
Eric would gain 23 happiness units by sitting next to Alice.
Eric would lose 69 happiness units by sitting next to Bob.
Eric would lose 33 happiness units by sitting next to Carol.
Eric would lose 47 happiness units by sitting next to David.
Eric would gain 75 happiness units by sitting next to Frank.
Eric would gain 82 happiness units by sitting next to George.
Eric would gain 13 happiness units by sitting next to Mallory.
Frank would gain 77 happiness units by sitting next to Alice.
Frank would gain 27 happiness units by sitting next to Bob.
Frank would lose 87 happiness units by sitting next to Carol.
Frank would gain 74 happiness units by sitting next to David.
Frank would lose 41 happiness units by sitting next to Eric.
Frank would lose 99 happiness units by sitting next to George.
Frank would gain 26 happiness units by sitting next to Mallory.
George would lose 63 happiness units by sitting next to Alice.
George would lose 51 happiness units by sitting next to Bob.
George would lose 60 happiness units by sitting next to Carol.
George would gain 30 happiness units by sitting next to David.
George would lose 100 happiness units by sitting next to Eric.
George would lose 63 happiness units by sitting next to Frank.
George would gain 57 happiness units by sitting next to Mallory.
Mallory would lose 71 happiness units by sitting next to Alice.
Mallory would lose 28 happiness units by sitting next to Bob.
Mallory would lose 10 happiness units by sitting next to Carol.
Mallory would gain 44 happiness units by sitting next to David.
Mallory would gain 22 happiness units by sitting next to Eric.
Mallory would gain 79 happiness units by sitting next to Frank.
Mallory would lose 16 happiness units by sitting next to George.`
