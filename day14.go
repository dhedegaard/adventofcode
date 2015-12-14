package main

import (
	"fmt"
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
		r.currMovesecs--
		r.location += r.speed
		if r.currMovesecs == 0 {
			r.currRestsecs = r.restsecs
		}
	} else {
		r.currRestsecs--
		if r.currRestsecs == 0 {
			r.currMovesecs = r.movesecs
		}
	}
}

// Race for a given number of ticks (or seconds) with the reindeers given.
// Returns the reindeer with the highest location.
func race(reindeers []*reindeer, ticks int) *reindeer {
	// 3 2 1 race !
	for i := 0; i < ticks; i++ {
		for _, r := range reindeers {
			r.tick()
		}
	}

	// Find the reindeer with the maximum location.
	var result *reindeer
	for _, r := range reindeers {
		if result == nil {
			result = r
		} else if r.location > result.location {
			result = r
		}
	}
	return result
}

func main() {
	var list []*reindeer
	list = append(list, newReindeer("comet", 14, 10, 127))
	list = append(list, newReindeer("dancer", 16, 11, 162))

	before := time.Now()
	result := race(list, 1000)
	fmt.Println(result.name, "wins with a distance of", result.location,
		"took:", time.Now().Sub(before))
}
