package graph

type Graph struct {
	Adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{Adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
	g.Adj[v] = append(g.Adj[v], u) // Для неориентированного графа
}

func HasEdge(g *Graph, u, v int) bool {
	for _, neighbor := range g.Adj[u] {
		if neighbor == v {
			return true
		}
	}
	return false
}
