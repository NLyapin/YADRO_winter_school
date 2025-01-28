package main

import (
	"fmt"
	"time"
	"project/internal/election"
)

func main() {
	// Инициализируем узлы.
	nodes := []*election.Node{
		election.NewNode(0),
		election.NewNode(1),
		election.NewNode(2),
		election.NewNode(3),
		election.NewNode(4),
	}

	// Запускаем обработчики сообщений для каждого узла.
	for _, node := range nodes {
		go node.HandleRingMessages(nodes)
	}

	// Узел 0 инициирует выборы.
	nodes[0].StartRingElection(nodes)

	// Даем время для обработки сообщений.
	time.Sleep(2 * time.Second)

	// Печатаем информацию о лидере.
	for _, node := range nodes {
		if node.LeaderID != -1 {
			fmt.Printf("Node %d: Leader is Node %d\n", node.ID, node.LeaderID)
		}
	}

	// Подождём ещё немного, чтобы увидеть, что все узлы обработали свои сообщения.
	time.Sleep(2 * time.Second)
}
