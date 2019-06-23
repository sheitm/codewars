package main

type jNode struct {
	val  int
	next *jNode
	prev *jNode
}

func (n *jNode) setNext(next *jNode) {
	n.next = next
}

func (n *jNode) setPrev(prev *jNode) {
	n.prev = prev
}

// JosephusSurvivor solve it
func JosephusSurvivor(n, k int) int {
	if n == 1 {
		return 1
	}
	ring := getJNodeRing(n)
	for {
		ring = skipNodes(ring, k)
		ring.prev.setNext(ring.next)
		ring.next.setPrev(ring.prev)
		if ring.next == ring {
			break
		}
	}
	return ring.val
}

func skipNodes(n *jNode, k int) *jNode {
	nn := n
	for i := 0; i < k; i++ {
		nn = nn.next
	}
	return nn
}

func getJNodeRing(n int) *jNode {
	nodes := []jNode{}
	for i := 0; i < n; i++ {
		n := jNode{val: i + 1}
		nodes = append(nodes, n)
	}
	for i := 0; i < len(nodes); i++ {
		if i == 0 {
			nodes[i].setNext(&nodes[i+1])
			nodes[i].setPrev(&nodes[len(nodes)-1])
			continue
		}
		if nodes[i].val == n {
			nodes[i].setNext(&nodes[0])
			nodes[i].setPrev(&nodes[i-1])
			continue
		}
		nodes[i].setNext(&nodes[i+1])
		nodes[i].setPrev(&nodes[i-1])
	}
	nodes[len(nodes)-1].setNext(&nodes[0])
	ring := &jNode{}
	ring.setNext(&nodes[0])

	return ring
}
