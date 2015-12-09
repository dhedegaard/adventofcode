package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type connection struct {
	distance int
	cities   []*city
}

// Returns the opposing city, from the cityname.
func (c connection) opposingCity(cityname string) *city {
	for _, cityPtr := range c.cities {
		if (*cityPtr).name != cityname {
			return cityPtr
		}
	}
	panic("WTF")
}

type city struct {
	name        string
	connections map[*connection]struct{}
}

func (c city) getLowestDistanceFrom(cities []*city,
	visitedCities []string) int {
	minimalDistance := 0
	for connection, _ := range c.connections {
		opposing := *connection.opposingCity(c.name)
		opposingName := opposing.name

		// Check to see if we've already visited this city.
		skip := false
		for _, visitedCity := range visitedCities {
			if visitedCity == opposingName {
				skip = true
			}
		}
		if skip {
			continue
		}

		// Calculate lowest distance from connection.
		weight := connection.distance + opposing.getLowestDistanceFrom(
			cities, append(visitedCities, c.name))
		if minimalDistance == 0 || weight < minimalDistance {
			minimalDistance = weight
		}
	}
	return minimalDistance
}

func (c city) getHighestDistanceFrom(cities []*city,
	visitedCities []string) int {
	maximumDistance := 0
	for connection, _ := range c.connections {
		opposing := *connection.opposingCity(c.name)
		opposingName := opposing.name

		// Check to see if we've already visited this city.
		skip := false
		for _, visitedCity := range visitedCities {
			if visitedCity == opposingName {
				skip = true
			}
		}
		if skip {
			continue
		}

		// Calculate highest distance from connection.
		weight := connection.distance + opposing.getHighestDistanceFrom(
			cities, append(visitedCities, c.name))
		if weight > maximumDistance {
			maximumDistance = weight
		}
	}
	return maximumDistance
}

func main() {
	// Parse the input into a graph.
	citymap := make(map[string]*city)
	var cities []*city
	scanner := bufio.NewScanner(strings.NewReader(input))
	reInstruction := regexp.MustCompile(`^(\S+?) to (\S+?) = (\d+)$`)
	for scanner.Scan() {
		rc := reInstruction.FindStringSubmatch(scanner.Text())

		// Add new cities to the citymap.
		for _, cityname := range rc[1:3] {
			if _, exists := citymap[cityname]; !exists {
				newCity := city{
					name:        cityname,
					connections: make(map[*connection]struct{}),
				}
				citymap[cityname] = &newCity
				cities = append(cities, &newCity)
			}
		}

		// Add distance.
		distInt, err := strconv.Atoi(rc[3])
		if err != nil {
			panic(err)
		}
		cityname1 := rc[1]
		cityname2 := rc[2]
		city1ptr := citymap[cityname1]
		city2ptr := citymap[cityname2]
		newConnection := connection{
			distance: distInt,
			cities: []*city{
				city1ptr,
				city2ptr,
			},
		}

		// Add the new connection to the cities it's connected from.
		citymap[cityname1].connections[&newConnection] = struct{}{}
		citymap[cityname2].connections[&newConnection] = struct{}{}
	}

	before := time.Now()
	minimalDistance := 99999999999
	maximumDistance := 0
	// Iterate on cities, finding the highest/lowest distance for visiting all
	// the cities only once.
	for _, cityPtr := range citymap {
		city := *cityPtr

		cityLowestDistance := city.getLowestDistanceFrom(
			cities, []string{})
		cityHighestDistance := city.getHighestDistanceFrom(
			cities, []string{})

		if cityLowestDistance < minimalDistance {
			minimalDistance = cityLowestDistance
		}

		if cityHighestDistance > maximumDistance {
			maximumDistance = cityHighestDistance
		}
	}

	fmt.Println("Minimal distance:", minimalDistance)
	fmt.Println("Maximum distance:", maximumDistance)
	fmt.Println("Time taken:", time.Now().Sub(before))
}

const input = `Faerun to Norrath = 129
Faerun to Tristram = 58
Faerun to AlphaCentauri = 13
Faerun to Arbre = 24
Faerun to Snowdin = 60
Faerun to Tambi = 71
Faerun to Straylight = 67
Norrath to Tristram = 142
Norrath to AlphaCentauri = 15
Norrath to Arbre = 135
Norrath to Snowdin = 75
Norrath to Tambi = 82
Norrath to Straylight = 54
Tristram to AlphaCentauri = 118
Tristram to Arbre = 122
Tristram to Snowdin = 103
Tristram to Tambi = 49
Tristram to Straylight = 97
AlphaCentauri to Arbre = 116
AlphaCentauri to Snowdin = 12
AlphaCentauri to Tambi = 18
AlphaCentauri to Straylight = 91
Arbre to Snowdin = 129
Arbre to Tambi = 53
Arbre to Straylight = 40
Snowdin to Tambi = 15
Snowdin to Straylight = 99
Tambi to Straylight = 70`
