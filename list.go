package List

//List impolement List structure in go
type List struct {
	start *Node
}

//Node is the node in List structure
type Node struct {
	Value interface{}
	Next  *Node
}

//Push push the node to the tail of List
func (l *List) Push(value interface{}) {
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

//Pop remove the last node of List and return it
func (l *List) Pop() interface{} {
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

//Shift remove the first node of List and return it
func (l *List) Shift() interface{} {
	if l.start == nil {
		return nil
	}

	temNode := l.start

	l.start = l.start.Next

	return temNode.Value
}

//GetLen return the length of List
func (l *List) GetLen() int32 {
	var length int32
	n := l.start
	for n != nil {
		length++
		n = n.Next
	}
	return length
}

//Clear clear the node of List and return the length
func (l *List) Clear() int32 {
	length := l.GetLen()

	l.start = nil

	return length
}

//GetValue return the value of List at indexTh
func (l *List) GetValue(index int32) interface{} {
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

//Splice delete or replace node of List
func (l *List) Splice(start int32, deleteCount int32, items ...interface{}) (removed []interface{}) {
	if l.GetLen() <= start {
		if len(items) > 0 {
			for _, item := range items {
				l.Push(item)
			}
		}
		return nil
	}

	newList := &List{}

	n := start
	var i int32

	for i = 0; i < n; i++ {
		newList.Push(l.GetValue(int32(i)))
	}

	if len(items) > 0 {
		for _, item := range items {
			newList.Push(item)
		}
	}

	// count := start
	for i = start + deleteCount; i < l.GetLen(); i++ {
		newList.Push(l.GetValue(int32(i)))
	}

	//get removed node and return it
	for i = start; i < start+deleteCount; i++ {
		removed = append(removed, l.GetValue(int32(i)))
	}
	l.start = newList.start
	return
}
