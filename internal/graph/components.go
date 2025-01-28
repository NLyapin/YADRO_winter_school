package graph

func ConnectedComponents(g *Graph) (int, map[int]int) {
	visited := make(map[int]bool)
	comp := make(map[int]int)
	count := 0

	for v := range g.Adj {
		if !visited[v] {
			count++
			var dfs func(int)
			dfs = func(u int) {
				visited[u] = true
				comp[u] = count
				for _, neighbor := range g.Adj[u] {
					if !visited[neighbor] {
						dfs(neighbor)
					}
				}
			}
			dfs(v)
		}
	}
	return count, comp
}
