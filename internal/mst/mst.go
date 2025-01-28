package mst

import (
	"project/internal/unionfind"
	"sort"
)

type Edge struct {
	U, V, Weight int
}

// MST находит минимальное остовное дерево с использованием алгоритма Крускала.
func MST(n int, edges []Edge) ([]Edge, int) {
	// Сортируем рёбра по весу.
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	mst := []Edge{}
	totalWeight := 0

	// Создаем DisjointSet для проверки циклов.
	ds := unionfind.NewDisjointSet(n)

	for _, edge := range edges {
		if ds.Union(edge.U, edge.V) {
			mst = append(mst, edge)
			totalWeight += edge.Weight
			// Если остовное дерево завершено.
			if len(mst) == n-1 {
				break
			}
		}
	}

	return mst, totalWeight
}
