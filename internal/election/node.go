package election

import (
	"math/rand"
	"time"
)

type Message struct {
	Kind   string      // "ELECTION", "COORDINATOR", "COLLECT", "COLLECT_REPLY"
	FromID int         // ID отправителя
	ToID   int         // ID получателя (если адресный)
	Data   interface{} // Дополнительные данные (например, localCount)
}

type Node struct {
	ID        int
	LeaderID  int
	Alive     bool
	LocalCount int // Локальное значение
	Inbox     chan Message
}

func NewNode(id int) *Node {
	rand.Seed(time.Now().UnixNano())
	return &Node{
		ID:        id,
		LeaderID:  -1, // -1 означает, что лидер неизвестен
		Alive:     true,
		LocalCount: rand.Intn(50) + 50,
		Inbox:     make(chan Message, 10),
	}
}
