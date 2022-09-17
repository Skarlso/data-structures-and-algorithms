package chapter18

import (
	"math"
)

type City struct {
	Routes map[*City]int
	Name   string
}

func NewCity(name string) *City {
	return &City{
		Name:   name,
		Routes: make(map[*City]int),
	}
}

func (c *City) AddRoute(city *City, price int) {
	c.Routes[city] = price
}

func DijkstraShortestPath(start *City, goal *City) []*City {
	cheapestPricesTable := make(map[*City]int)
	// cameFrom is a lot better... but let's stick to the book
	cheapestPreviousStopoverCityTable := make(map[*City]*City)

	// the book uses an array for this, I'm not sure why... A map is a lot simpler here.
	unvisitedCities := make(map[*City]struct{})
	visited := make(map[*City]struct{})

	// Set up start as 0.
	cheapestPricesTable[start] = 0

	current := start

	// we run as long as there are cities to visit
	for current != nil {
		visited[current] = struct{}{}
		delete(unvisitedCities, current)

		for city, price := range current.Routes {
			if _, ok := visited[city]; !ok {
				unvisitedCities[city] = struct{}{}
			}

			currentPrice := cheapestPricesTable[city] + price
			if v, ok := cheapestPricesTable[city]; !ok || currentPrice < v {
				cheapestPricesTable[city] = currentPrice
				cheapestPreviousStopoverCityTable[city] = current
			}
		}

		min := math.MaxInt64
		var minCity *City
		for c := range unvisitedCities {
			if cheapestPricesTable[c] < min {
				min = cheapestPricesTable[c]
				minCity = c
			}
		}
		current = minCity
	}

	// The core is done. Let's walk backwards to find the path
	shortestPath := make([]*City, 0)

	// We go backwards from the goal to start. Traverse the path back to the origin.
	currentCity := goal

	for currentCity != start {
		shortestPath = append(shortestPath, currentCity)
		currentCity = cheapestPreviousStopoverCityTable[currentCity]
	}

	shortestPath = append(shortestPath, start)

	// I'm returning this list, but it's in reverse order.
	return shortestPath
}
