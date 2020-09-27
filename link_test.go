package link

import (
	"testing"
)

func TestLink(t *testing.T) {
	//test Push
	t.Run("testPush", func(t *testing.T) {
		var l *Link = &Link{}
		l.Push("1")
		if Value, ok := l.start.Value.(string); !ok {
			t.Fatal("Type error, want string but not.")
			if Value != "1" {
				t.Fatalf("Value error, want %q but %s", "1", Value)
			}
		}
		l.Push(2)
		if Value, ok := l.start.Next.Value.(int); !ok {
			t.Fatal("Type error, want int but not.")
			if Value != 2 {
				t.Fatalf("Value error, want %d but %d", 2, Value)
			}
		}
		l.Push(true)
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
		l.Push(1)
		l.Push(true)
		l.Push(3)
		l.Push("4")

		//string
		value := l.Pop()
		if v, _ := value.(string); v != "4" {
			t.Fatalf("Pop value error, expect %q but %s", "4", v)
		}

		//int
		value = l.Pop()
		if v, _ := value.(int); v != 3 {
			t.Fatalf("Pop value error, expect %d but %d", 3, v)
		}

		//boolean
		value = l.Pop()
		if v, _ := value.(bool); v != true {
			t.Fatalf("Pop value error, expect %v but %v", true, v)
		}

		//int
		value = l.Pop()
		if v, _ := value.(int); v != 1 {
			t.Fatalf("Pop value error, expect %d but %d", 1, v)
		}

		//nil
		value = l.Pop()
		if value != nil {
			t.Fatalf("Node expect nil ,but %v.", value)
		}
	})
}
