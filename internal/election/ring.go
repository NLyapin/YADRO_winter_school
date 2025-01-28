package election

import "fmt"

// Узел инициирует выборы в кольце.
func (n *Node) StartRingElection(nodes []*Node) {
	if !n.Alive {
		return
	}
	fmt.Printf("Node %d: Starting Ring Election\n", n.ID)
	nextID := (n.ID + 1) % len(nodes)
	nodes[nextID].Inbox <- Message{Kind: "ELECTION", FromID: n.ID, Data: n.ID}
}

// Узел обрабатывает сообщения кольцевого выбора.
func (n *Node) HandleRingMessages(nodes []*Node) {
	for msg := range n.Inbox {
		switch msg.Kind {
		case "ELECTION":
			maxID := msg.Data.(int)
			if n.ID > maxID {
				maxID = n.ID
			}
			nextID := (n.ID + 1) % len(nodes)
			if msg.FromID == n.ID {
				fmt.Printf("Node %d: Election complete, leader is %d\n", n.ID, maxID)
				n.LeaderID = maxID
				n.AnnounceCoordinator(nodes, maxID)
			} else {
				nodes[nextID].Inbox <- Message{Kind: "ELECTION", FromID: msg.FromID, Data: maxID}
			}
		case "COORDINATOR":
			fmt.Printf("Node %d: Node %d is the new leader\n", n.ID, msg.FromID)
			n.LeaderID = msg.FromID
		}
	}
}

// Узел объявляет лидера.
func (n *Node) AnnounceCoordinator(nodes []*Node, leaderID int) {
	for _, node := range nodes {
		if node.Alive {
			node.Inbox <- Message{Kind: "COORDINATOR", FromID: leaderID}
		}
	}
}
