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
func (l *Link) Push(value interface{}) {
	node := &Node{Value: value}
	if l.start == nil {
		l.start = node
	} else {
		n := l.start
		for n.Next != nil {
			n = n.Next
		}
		n.Next = node
	}
}

//Pop remove the last node of link and return the node
func (l *Link) Pop() interface{} {
	n := l.start
	if n == nil {
		return nil
	}

	if n.Next == nil {
		l.start = nil
		return n.Value
	}

	for n.Next.Next != nil {
		n = n.Next
	}

	temNode := n.Next
	n.Next = nil
	return temNode.Value
}
