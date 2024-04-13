package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(DegreesOfSeparation([]string{"Alice:Bob", "Bob:Charlie", "Charlie:David", "David:Eve"}, "Alice", "Charliestring"))
}
func DegreesOfSeparation(connections []string, person1 string, person2 string) int {
	// Create a map to represent the network of friends
	network := make(map[string][]string)
	for _, connection := range connections {
		friends := strings.Split(connection, ":")
		network[friends[0]] = append(network[friends[0]], friends[1])
		network[friends[1]] = append(network[friends[1]], friends[0])
	}

	// Perform BFS to find the degrees of separation
	queue := []string{person1}
	visited := map[string]struct{}{
		person1: struct{}{},
	}
	degrees := map[string]int{
		person1: 0,
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == person2 {
			return degrees[current]
		}

		for _, friend := range network[current] {
			if _, ok := visited[friend]; !ok {
				queue = append(queue, friend)
				visited[friend] = struct{}{}
				degrees[friend] = degrees[current] + 1
			}
		}
	}

	// If no connection is found
	return -1
}
