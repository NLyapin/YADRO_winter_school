package unionfind

type DisjointSet struct {
	parent []int
	rank   []int
}

// NewDisjointSet создает и инициализирует структуру DisjointSet.
func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		ds.parent[i] = i
	}
	return ds
}

// Find выполняет поиск сжатия пути для элемента x.
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // Сжатие пути.
	}
	return ds.parent[x]
}

// Union объединяет множества x и y, если они еще не объединены.
func (ds *DisjointSet) Union(x, y int) bool {
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	if rootX == rootY {
		return false // Уже в одном множестве.
	}

	// Объединяем по рангу.
	if ds.rank[rootX] > ds.rank[rootY] {
		ds.parent[rootY] = rootX
	} else if ds.rank[rootX] < ds.rank[rootY] {
		ds.parent[rootX] = rootY
	} else {
		ds.parent[rootY] = rootX
		ds.rank[rootX]++
	}

	return true
}
