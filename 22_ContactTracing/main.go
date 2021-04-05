package main

import "fmt"

func contains(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}

func createMap(neighbors [][]string) map[string][]string {
	var adjacencyMap = make(map[string][]string)
	for _, pair := range neighbors {
		if !contains(adjacencyMap[pair[0]], pair[1]) {
			adjacencyMap[pair[0]] = append(adjacencyMap[pair[0]], pair[1])
		}
		if !contains(adjacencyMap[pair[1]], pair[0]) {
			adjacencyMap[pair[1]] = append(adjacencyMap[pair[1]], pair[0])
		}
	}
	return adjacencyMap
}

func main() {
	neighbors := [][]string{
		[]string{"Todd", "Kim"},
		[]string{"Todd", "Thomas"},
		[]string{"Kim", "Mike"},
		[]string{"Susan", "Kim"},
		[]string{"Phillip", "Kim"},
		[]string{"Mason", "Todd"},
		[]string{"Thomas", "Trevor"},
		[]string{"Rick", "Thomas"},
	}

	adjMap := createMap(neighbors)
	var visited []string

	var bfs func(start []string, end string, depth int) int

	bfs = func(start []string, end string, depth int) int {
		visited = append(visited, start...)

		var adj []string
		for _, person := range start {
			adj = append(adj, adjMap[person]...)
		}
		if contains(adj, end) {
			return depth + 1
		}
		depth = bfs(adj, end, depth+1)
		return depth
	}

	// start, end := "Kim", "Mike" // 1
	// start, end := "Kim", "Mason" // 2
	// start, end := "Kim", "Trevor" // 3
	start, end := "Kim", "Theo" // 0

	fmt.Printf("Distance from %s to %s is %d\n", start, end, bfs([]string{start}, end, 0))
}
