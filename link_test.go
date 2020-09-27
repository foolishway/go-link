package link

import (
	"testing"
)

func TestLink(t *testing.T) {
	//test Push
	t.Run("testPush", func(t *testing.T) {
		var l *Link = &Link{}
		l.Push(&Node{Value: "1"})
		if Value, ok := l.start.Value.(string); !ok {
			t.Fatal("Type error, want string but not.")
			if Value != "1" {
				t.Fatalf("Value error, want %q but %s", "1", Value)
			}
		}
		l.Push(&Node{Value: 2})
		if Value, ok := l.start.Next.Value.(int); !ok {
			t.Fatal("Type error, want int but not.")
			if Value != 2 {
				t.Fatalf("Value error, want %d but %d", 2, Value)
			}
		}
		l.Push(&Node{Value: true})
		if Value, ok := l.start.Next.Next.Value.(bool); !ok {
			t.Fatal("Type error, want boolean but not.")
			if Value != true {
				t.Fatalf("Value error, want %v but %v", true, Value)
			}
		}
	})

	//test pop
	t.Run("testPop", func(t *testing.T) {
		var l *Link = &Link{}
		l.Push(&Node{Value: 1})
		l.Push(&Node{Value: true})
		l.Push(&Node{Value: 3})
		l.Push(&Node{Value: "4"})

		//string
		node := l.Pop()
		if node == nil {
			t.Fatal("Node should not be nil, but nil.")
		}
		if _, ok := node.Value.(string); !ok {
			t.Fatal("Pop type error, expect string ,but not.")

		}
		if value, _ := node.Value.(string); value != "4" {
			t.Fatalf("Pop value error, expect %q but %s", "4", value)
		}

		//int
		node = l.Pop()
		if node == nil {
			t.Fatal("Node should not be nil, but nil.")
		}
		if _, ok := node.Value.(int); !ok {
			t.Fatal("Pop type error, expect int ,but not.")

		}
		if value, _ := node.Value.(int); value != 3 {
			t.Fatalf("Pop value error, expect %d but %d", 3, value)
		}

		//boolean
		node = l.Pop()
		if node == nil {
			t.Fatal("Node should not be nil, but nil.")
		}
		if _, ok := node.Value.(bool); !ok {
			t.Fatal("Pop type error, expect int ,but not.")

		}
		if value, _ := node.Value.(bool); value != true {
			t.Fatalf("Pop value error, expect %v but %v", true, value)
		}

		//int
		node = l.Pop()
		if node == nil {
			t.Fatal("Node should not be nil, but nil.")
		}
		if _, ok := node.Value.(int); !ok {
			t.Fatal("Pop type error, expect int ,but not.")

		}
		if value, _ := node.Value.(int); value != 1 {
			t.Fatalf("Pop value error, expect %d but %d", 1, value)
		}

		//nil
		node = l.Pop()
		if node != nil {
			t.Fatal("Node expect nil ,but not.")
		}
	})
}
