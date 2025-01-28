package main

import (
	"fmt"
	"project/internal/graph"
)

func main() {
	g := graph.NewGraph()

	// Добавление рёбер
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 6)
	g.AddEdge(3, 7)

	// Вывод структуры графа
	fmt.Println("Graph adjacency list:", g.Adj)

	// Проверка существования рёбер
	fmt.Println("HasEdge(0, 1):", graph.HasEdge(g, 0, 1))
	fmt.Println("HasEdge(0, 3):", graph.HasEdge(g, 0, 3))

	// Обходы
	fmt.Println("BFS order:", graph.BFS(g, 0))
	fmt.Println("DFS order:", graph.DFS(g, 0))

	// Поиск компонент связности
	count, comp := graph.ConnectedComponents(g)
	fmt.Println("Number of connected components:", count)
	fmt.Println("Components:", comp)
}
