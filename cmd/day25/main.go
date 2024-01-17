package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"slices"
	"strings"
)

type Graph map[string][]string

type Set map[string]bool

type Edge [2]string

type EdgeList []Edge

func dfs(graph Graph, visited Set, currNode string) int {
	if visited[currNode] {
		return 0
	}
	visited[currNode] = true
	numComponents := 1
	for _, nei := range graph[currNode] {
		numComponents += dfs(graph, visited, nei)
	}
	return numComponents
}

func isConnectedGraph(graph Graph) bool {
	visited := Set{}
	numConnectedComponents := 0
	for node := range graph {
		if !visited[node] {
			numConnectedComponents++
			if numConnectedComponents > 1 {
				return false
			}
			dfs(graph, visited, node)
		}
	}
	return true
}

func parseIntoGraph(lines []string) (Graph, EdgeList) {
	graph := Graph{}
	edgeList := EdgeList{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		vertex := parts[0]
		neis := strings.Split(parts[1], " ")
		graph[vertex] = append(graph[vertex], neis...)
		for _, nei := range neis {
			graph[nei] = append(graph[nei], vertex)
		}
	}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		vertex := parts[0]
		neis := strings.Split(parts[1], " ")
		for _, nei := range neis {
			edgeList = append(edgeList, [2]string{vertex, nei})
		}
	}
	return graph, edgeList
}

func deleteEdge(graph Graph, edge Edge) {
	graph[edge[0]] = slices.DeleteFunc(graph[edge[0]], func(s string) bool { return s == edge[1] })
	graph[edge[1]] = slices.DeleteFunc(graph[edge[1]], func(s string) bool { return s == edge[0] })
}

func addEdge(graph Graph, edge Edge) {
	graph[edge[0]] = append(graph[edge[0]], edge[1])
	graph[edge[1]] = append(graph[edge[1]], edge[0])
}

func countConnectedComponents(graph Graph) int {
	visited := Set{}
	ans := 1
	for node := range graph {
		if !visited[node] {
			ans *= dfs(graph, visited, node)
		}
	}
	return ans
}

func solvePart1(lines []string) int {
	graph, edgelist := parseIntoGraph(lines)
	for _, edge1 := range edgelist {
		for _, edge2 := range edgelist {
			for _, edge3 := range edgelist {
				if !((edge1 != edge2) && (edge2 != edge3) && (edge1 != edge3)) {
					continue
				}
				deleteEdge(graph, edge1)
				deleteEdge(graph, edge2)
				deleteEdge(graph, edge3)

				if !isConnectedGraph(graph) {
					return countConnectedComponents(graph)
				}

				addEdge(graph, edge1)
				addEdge(graph, edge2)
				addEdge(graph, edge3)
			}
		}
	}
	return 0
}

func main() {
	lines := aoc.GetInputLinesForDay(25, false)
	ans := solvePart1(lines)
	fmt.Println(ans)
}
