package link

type link struct {
	start *node
}
type node struct {
	value interface{}
	next  *node
}

func (l *link) push(pl *node) {
	if l.start == nil {
		l.start = pl
	} else {
		n := l.start
		for n.next != nil {
			n = n.next
		}
		n.next = pl
	}
}

// func (l *link) pop() *node {
// 	var n *node = l.start
// 	var p *node
// 	for n.next != nil {
// 		p = n
// 		n = n.next
// 	}
// 	rn := n
// 	p = nil
// 	return rn
// }
