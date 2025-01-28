package graph

import "project/internal/queue"

func BFS(g *Graph, start int) []int {
	visited := make(map[int]bool)
	var order []int
	q := queue.Queue{}

	q.Enqueue(start)
	visited[start] = true

	for !q.IsEmpty() {
		u, _ := q.Dequeue()
		order = append(order, u)

		for _, v := range g.Adj[u] {
			if !visited[v] {
				visited[v] = true
				q.Enqueue(v)
			}
		}
	}
	return order
}
