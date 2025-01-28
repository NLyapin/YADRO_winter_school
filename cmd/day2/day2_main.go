package main

import (
	"fmt"
	"project/internal/unionfind"
	"project/internal/mst"
	"project/internal/dijkstra"
)

func main() {
	// Пример работы с DisjointSet.
	ds := unionfind.NewDisjointSet(6)
	ds.Union(0, 1)
	ds.Union(1, 2)
	ds.Union(3, 4)
	ds.Union(2, 3)
	fmt.Println("Find results:", ds.Find(0), ds.Find(1), ds.Find(2), ds.Find(3), ds.Find(4), ds.Find(5))

	// Пример работы MST.
	edges := []mst.Edge{
		{U: 0, V: 1, Weight: 1},
		{U: 1, V: 2, Weight: 2},
		{U: 0, V: 2, Weight: 3},
		{U: 2, V: 3, Weight: 4},
		{U: 3, V: 4, Weight: 5},
	}
	mstResult, totalWeight := mst.MST(5, edges)
	fmt.Println("MST edges:", mstResult)
	fmt.Println("Total MST weight:", totalWeight)

	// Пример работы Dijkstra.
	graph := dijkstra.Graph{
		Adj: map[int][]struct {
			To     int
			Weight int
		}{
			0: {{To: 1, Weight: 1}, {To: 2, Weight: 4}},
			1: {{To: 2, Weight: 2}, {To: 3, Weight: 6}},
			2: {{To: 3, Weight: 3}},
			3: {},
		},
	}
	dist, parent := dijkstra.Dijkstra(&graph, 0)
	fmt.Println("Dijkstra distances:", dist)
	fmt.Println("Dijkstra parents:", parent)
}
