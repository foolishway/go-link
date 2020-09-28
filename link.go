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

//Pop remove the last node of link and return it
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

//Shift remove the first node of link and return it
func (l *Link) Shift() interface{} {
	if l.start == nil {
		return nil
	}

	temNode := l.start

	l.start = l.start.Next

	return temNode.Value
}

//GetLen return the length of link
func (l *Link) GetLen() int32 {
	var length int32
	n := l.start
	for n != nil {
		length++
		n = n.Next
	}
	return length
}

//Clear clear the node of link and return the length
func (l *Link) Clear() int32 {
	length := l.GetLen()

	l.start = nil

	return length
}

//GetValue return the value of link at indexTh
func (l *Link) GetValue(index int32) interface{} {
	if index >= l.GetLen() {
		return nil
	}

	var count int32
	n := l.start
	for index != count {
		count++
		n = n.Next
	}
	return n.Value
}

//Splice delete or replace node of link
func (l *Link) Splice(start int32, deleteCount int32, items ...interface{}) (removed []interface{}) {
	if l.GetLen() <= start {
		if len(items) > 0 {
			for _, item := range items {
				l.Push(item)
			}
		}
		return nil
	}

	newLink := &Link{}

	n := start
	var i int32

	for i = 0; i < n; i++ {
		newLink.Push(l.GetValue(int32(i)))
	}

	if len(items) > 0 {
		for _, item := range items {
			newLink.Push(item)
		}
	}

	// count := start
	for i = start + deleteCount; i < l.GetLen(); i++ {
		newLink.Push(l.GetValue(int32(i)))
	}

	//get removed node and return it
	for i = start; i < start+deleteCount; i++ {
		removed = append(removed, l.GetValue(int32(i)))
	}
	l.start = newLink.start
	return
}
