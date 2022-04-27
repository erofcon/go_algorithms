package algorithms

import "fmt"

func RunBfs() {
	//a -> b -> c -> f
	//|         |
	//d -> i -> j

	graph := map[string][]string{
		"a": {"b"},
		"b": {"a", "c", "d"},
		"c": {"b", "f"},
		"d": {"b", "i"},
		"i": {"d", "j"},
		"f": {"c", "j"},
		"j": {"f", "i"},
	}

	start := "a"
	goal := "i"

	visited := bfs(start, goal, graph)
	currentNode := goal
	fmt.Print(currentNode)
	for currentNode != start {
		currentNode = visited[currentNode]
		fmt.Print(" --> ", currentNode)
	}
}

func bfs(start string, goal string, graph map[string][]string) map[string]string {
	queue := make([]string, 0)
	queue = append(queue, start)
	visited := map[string]string{
		start: "",
	}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if currentNode == goal {
			break
		}

		nextNodes := graph[currentNode]

		for _, nextNode := range nextNodes {
			_, v := visited[nextNode]
			if !v {
				queue = append(queue, nextNode)
				visited[nextNode] = currentNode
			}
		}

	}

	return visited
}
