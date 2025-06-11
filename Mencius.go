package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
const (
	TotalNodes     = 5
	ProposalWindow = 200 * time.Millisecond
	RunDuration    = 10 * time.Second
)
type Node struct {
	mu         sync.Mutex
	nodeID     int
	peers      []int
	throughput int
}
func NewNode(nodeID int) *Node {
	peers := make([]int, 0, TotalNodes-1)
	for i := 0; i < TotalNodes; i++ {
		if i != nodeID {
			peers = append(peers, i)
		}
	}
	return &Node{
		nodeID:     nodeID,
		peers:      peers,
		throughput: 0,
	}
}

func (n *Node) Propose(wg *sync.WaitGroup, stopCh <-chan struct{}) {
	defer wg.Done()
	ticker := time.NewTicker(ProposalWindow)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			n.mu.Lock()
			n.throughput++
			n.mu.Unlock()
			for _, peer := range n.peers {
				fmt.Printf("Node %d proposes to Node %d\n", n.nodeID, peer)
			}
		case <-stopCh:
			return
		}
	}
}
func main() {
	var wg sync.WaitGroup
	nodes := make([]*Node, TotalNodes)
	stopCh := make(chan struct{})

	for i := 0; i < TotalNodes; i++ {
		nodes[i] = NewNode(i)
		wg.Add(1)
		go nodes[i].Propose(&wg, stopCh)
	}

	time.Sleep(RunDuration)
	close(stopCh)
	wg.Wait()

	for _, node := range nodes {
		fmt.Printf("Node %d achieved throughput: %d\n", node.nodeID, node.throughput)
	}
}
