package election

import "fmt"

// Узел инициирует выборы, отправляя "ELECTION" узлам с ID > его ID.
func (n *Node) StartBullyElection(nodes []*Node) {
	if !n.Alive {
		return
	}
	fmt.Printf("Node %d: Starting Bully Election\n", n.ID)
	hasHigher := false

	for _, node := range nodes {
		if node.ID > n.ID && node.Alive {
			fmt.Printf("Node %d: Sending ELECTION to Node %d\n", n.ID, node.ID)
			node.Inbox <- Message{Kind: "ELECTION", FromID: n.ID}
			hasHigher = true
		}
	}

	if !hasHigher {
		n.DeclareAsLeader(nodes)
	}
}

// Узел обрабатывает входящие сообщения для Bully-алгоритма.
func (n *Node) HandleBullyMessages(nodes []*Node) {
	for msg := range n.Inbox {
		switch msg.Kind {
		case "ELECTION":
			if n.Alive && n.ID > msg.FromID {
				fmt.Printf("Node %d: Responding OK to Node %d\n", n.ID, msg.FromID)
				nodes[msg.FromID].Inbox <- Message{Kind: "OK", FromID: n.ID}
				n.StartBullyElection(nodes)
			}
		case "OK":
			fmt.Printf("Node %d: Received OK from Node %d\n", n.ID, msg.FromID)
		case "COORDINATOR":
			fmt.Printf("Node %d: Node %d is the new leader\n", n.ID, msg.FromID)
			n.LeaderID = msg.FromID
		}
	}
}

// Узел объявляет себя лидером.
func (n *Node) DeclareAsLeader(nodes []*Node) {
	fmt.Printf("Node %d: Declaring myself as leader\n", n.ID)
	n.LeaderID = n.ID
	for _, node := range nodes {
		if node.Alive {
			node.Inbox <- Message{Kind: "COORDINATOR", FromID: n.ID}
		}
	}
}
