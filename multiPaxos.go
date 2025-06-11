package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
const (
	Timeout    = 1500 * time.Millisecond
	TotalNodes = 5
)
type Node struct {
	mu        sync.Mutex
	nodeID    int
	state     string
	term      int
	accepted  bool
	peers     []int
	timeout   time.Duration
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
		nodeID:    nodeID,
		state:     "FOLLOWER",
		term:      0,
		accepted:  false,
		peers:     peers,
		timeout:   Timeout,
		throughput: 0,
	}
}
func (n *Node) StartProposal() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.state = "PROPOSER"
	n.term++
	n.accepted = false
	fmt.Printf("Node %d starts proposal for term %d\n", n.nodeID, n.term)
	for _, peer := range n.peers {
		n.RequestProposal(peer)
	}
}

func (n *Node) RequestProposal(peer int) {
	fmt.Printf("Node %d sends proposal to Node %d\n", n.nodeID, peer)
}

func (n *Node) OnReceiveProposal(accepted bool) {
	n.mu.Lock()
	defer n.mu.Unlock()
	if accepted {
		n.accepted = true
		n.throughput++
		if n.throughput > TotalNodes/2 {
			n.state = "ACCEPTED"
			fmt.Printf("Node %d proposal accepted for term %d\n", n.nodeID, n.term)
		}
	}
}
func (n *Node) SendHeartbeat() {
	if n.state == "ACCEPTED" {
		for _, peer := range n.peers {
			fmt.Printf("Leader Node %d sending heartbeat to Node %d\n", n.nodeID, peer)
		}
	}
}

func (n *Node) ProposalTimeout() {
	if n.state != "ACCEPTED" {
		n.StartProposal()
	}
}
func (n *Node) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if n.state == "PROPOSER" {
			n.ProposalTimeout()
		} else if n.state == "ACCEPTED" {
			n.SendHeartbeat()
		}
		time.Sleep(n.timeout)
	}
}

func main() {
	var wg sync.WaitGroup
	nodes := make([]*Node, TotalNodes)
	for i := 0; i < TotalNodes; i++ {
		nodes[i] = NewNode(i)
		wg.Add(1)
		go nodes[i].Run(&wg)
	}
	time.Sleep(10 * time.Second)
	wg.Wait()

	for _, node := range nodes {
		fmt.Printf("Node %d achieved throughput: %d\n", node.nodeID, node.throughput)
	}
}
