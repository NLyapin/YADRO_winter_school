package graph

import "project/internal/stack"

// Рекурсивная реализация DFS
func dfsUtil(g *Graph, v int, visited map[int]bool, order *[]int) {
	visited[v] = true
	*order = append(*order, v)
	for _, u := range g.Adj[v] {
		if !visited[u] {
			dfsUtil(g, u, visited, order)
		}
	}
}

func DFS(g *Graph, start int) []int {
	visited := make(map[int]bool)
	var order []int
	dfsUtil(g, start, visited, &order)
	return order
}

// Итеративная реализация DFS
func DFSIterative(g *Graph, start int) []int {
	visited := make(map[int]bool)
	var order []int
	s := stack.Stack{}

	s.Push(start)

	for !s.IsEmpty() {
		u, _ := s.Pop()
		if !visited[u] {
			visited[u] = true
			order = append(order, u)

			for i := len(g.Adj[u]) - 1; i >= 0; i-- { // Добавляем в обратном порядке
				s.Push(g.Adj[u][i])
			}
		}
	}
	return order
}
