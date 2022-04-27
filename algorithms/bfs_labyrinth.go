package algorithms

import "fmt"

type Key struct {
	i, k int
}

func RunBfsLabyrinth() {
	grid := [][]string{
		{"-", "-", "X", "X", "X"},
		{"-", "X", "-", "X", "X"},
		{"-", "-", "X", "-", "X"},
		{"-", "-", "X", "X", "-"},
		{"X", "-", "-", "-", "-"},
	}
	start := []int{0, 0}
	goal := []int{4, 4}

	graph := make(map[Key][][]int)

	for i, _ := range grid {
		for k, _ := range grid[i] {
			next := getNextNode(i, k, grid)
			graph[Key{i, k}] = append(graph[Key{i, k}], next...)
		}
	}

	result := search(start, goal, graph)
	currentNode := goal
	grid[start[0]][start[1]] = "+"
	for {
		if currentNode[0] == start[0] && currentNode[1] == start[1] {
			break
		}
		grid[currentNode[0]][currentNode[1]] = "+"
		currentNode = result[Key{currentNode[0], currentNode[1]}]

	}
	for i, _ := range grid {
		for k, _ := range grid[i] {
			fmt.Print(grid[i][k], " ")
		}
		fmt.Println()
	}
}

func search(start []int, goal []int, graph map[Key][][]int) map[Key][]int {

	queue := make([][]int, 0)
	queue = append(queue, start)
	visited := map[Key][]int{
		{start[0], start[1]}: {},
	}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if currentNode[0] == goal[0] && currentNode[1] == goal[1] {
			break
		}

		nextNodes := graph[Key{currentNode[0], currentNode[1]}]
		for _, next := range nextNodes {
			_, v := visited[Key{next[0], next[1]}]
			if !v {
				queue = append(queue, next)
				visited[Key{next[0], next[1]}] = currentNode
			}
		}
	}
	return visited
}

func getNextNode(x int, y int, matrix [][]string) [][]int {

	next := make([][]int, 0)
	ways := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}
	check := func(a, b int) bool {
		if 0 <= a && a < len(matrix) && 0 <= b && b < len(matrix[:]) {
			if matrix[a][b] != "X" {
				return true
			}
		}
		return false
	}
	for _, v := range ways {
		if check(x+v[0], y+v[1]) {
			next = append(next, []int{x + v[0], y + v[1]})
		}
	}
	return next
}
