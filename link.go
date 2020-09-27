package link

//Link impolement Link structure in go
type Link struct {
	start *Node
}

//Node is the node in Link structure
type Node struct {
	Value interface{}
	Next  *Node
}

//Push push the node to the tail of link
func (l *Link) Push(pl *Node) {
	if l.start == nil {
		l.start = pl
	} else {
		n := l.start
		for n.Next != nil {
			n = n.Next
		}
		n.Next = pl
	}
}

//Pop remove the last node of link and return the node
func (l *Link) Pop() *Node {
	n := l.start
	if n == nil {
		return nil
	}

	if n.Next == nil {
		l.start = nil
		return n
	}

	for n.Next.Next != nil {
		n = n.Next
	}

	temNode := n.Next
	n.Next = nil
	return temNode
}
