package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type reindeer struct {
	name         string
	speed        int
	movesecs     int
	restsecs     int
	location     int
	currMovesecs int
	currRestsecs int
	score        int
}

func newReindeer(name string, speed int, movesecs int, restsecs int) *reindeer {
	return &reindeer{
		name:         name,
		speed:        speed,
		movesecs:     movesecs,
		restsecs:     restsecs,
		currMovesecs: movesecs,
	}
}

func (r *reindeer) tick() {
	if r.currMovesecs > 0 {
		// We're flying.
		r.currMovesecs--
		r.location += r.speed
		if r.currMovesecs == 0 {
			// We're done flying, start resting.
			r.currRestsecs = r.restsecs
		}
	} else {
		// We're resting.
		r.currRestsecs--
		if r.currRestsecs == 0 {
			// We're done resting, start flying again.
			r.currMovesecs = r.movesecs
		}
	}
}

// Race for a given number of ticks (or seconds) with the reindeers given.
// Returns the reindeer with the highest location and score.
func race(reindeers []*reindeer, ticks int) (*reindeer, *reindeer) {
	// 3 2 1 race !
	for i := 0; i < ticks; i++ {
		// Move all the reindeers one second into the future.
		for _, r := range reindeers {
			r.tick()
		}
		// Find the max locations.
		var maxLocation int
		for _, r := range reindeers {
			if r.location > maxLocation {
				maxLocation = r.location
			}
		}
		// Give one point to each reindeer at that location.
		for _, r := range reindeers {
			if r.location == maxLocation {
				r.score++
			}
		}
	}

	// Find the reindeers with the maximum location/score.
	var maxDistance, maxScore *reindeer
	for _, r := range reindeers {
		if maxDistance == nil && maxScore == nil {
			maxDistance = r
			maxScore = r
			continue
		}

		if r.location > maxDistance.location {
			maxDistance = r
		}
		if r.score > maxScore.score {
			maxScore = r
		}
	}
	return maxDistance, maxScore
}

// Parse input as an interger, or panic().
func toIntOrPanic(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	// Parse the reindeers into a slice.
	reInput := regexp.MustCompile(
		`(\S+?) can fly (\d+?) km/s for (\d+?) seconds, but then must rest for (\d+?) seconds.`)
	var list []*reindeer
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		input := scanner.Text()
		rc := reInput.FindStringSubmatch(input)
		reindeer := newReindeer(rc[1], toIntOrPanic(rc[2]),
			toIntOrPanic(rc[3]), toIntOrPanic(rc[4]))
		list = append(list, reindeer)
	}

	// Race !
	before := time.Now()
	maxDistance, maxScore := race(list, 2503)
	fmt.Println(maxDistance.name, "wins part1 with a distance of", maxDistance.location)
	fmt.Println(maxScore.name, "wins part2 with a score of", maxScore.score)
	fmt.Println("took:", time.Now().Sub(before))
}

const input = `Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
Rudolph can fly 3 km/s for 15 seconds, but then must rest for 28 seconds.
Donner can fly 19 km/s for 9 seconds, but then must rest for 164 seconds.
Blitzen can fly 19 km/s for 9 seconds, but then must rest for 158 seconds.
Comet can fly 13 km/s for 7 seconds, but then must rest for 82 seconds.
Cupid can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.
Dasher can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.
Dancer can fly 3 km/s for 16 seconds, but then must rest for 37 seconds.
Prancer can fly 25 km/s for 6 seconds, but then must rest for 143 seconds.`
