package dijkstra

import (
	"math"
)

type Graph struct {
	Adj map[int][]struct {
		To     int
		Weight int
	}
}

// Dijkstra находит кратчайшие пути из start до всех вершин графа.
func Dijkstra(g *Graph, start int) ([]int, []int) {
	n := len(g.Adj)
	dist := make([]int, n)
	parent := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		parent[i] = -1
	}
	dist[start] = 0

	for i := 0; i < n; i++ {
		// Находим вершину с минимальным расстоянием.
		u := -1
		for v := 0; v < n; v++ {
			if !visited[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}

		if dist[u] == math.MaxInt32 {
			break // Остальные вершины недостижимы.
		}

		visited[u] = true

		// Обновляем расстояния для соседей.
		for _, edge := range g.Adj[u] {
			v, weight := edge.To, edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				parent[v] = u
			}
		}
	}

	return dist, parent
}
