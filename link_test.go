package link

import (
	"log"
	"math/rand"
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

	//test shift
	t.Run("testShift", func(t *testing.T) {
		var l *Link = &Link{}
		l.Push(1)
		l.Push(true)
		l.Push(3)
		l.Push("4")

		//int
		value := l.Shift()
		if v, _ := value.(int); v != 1 {
			t.Fatalf("Shift value error, expect %d but %d", 1, v)
		}

		//boolean
		value = l.Shift()
		if v, _ := value.(bool); v != true {
			t.Fatalf("Shift value error, expect %v but %v", true, v)
		}

		//int
		value = l.Shift()
		if v, _ := value.(int); v != 3 {
			t.Fatalf("Shift value error, expect %d but %d", 3, v)
		}
		//string
		value = l.Shift()
		if v, _ := value.(string); v != "4" {
			t.Fatalf("Shift value error, expect %q but %s", "4", v)
		}
		//nil
		value = l.Pop()
		if value != nil {
			t.Fatalf("Node expect nil ,but %v.", value)
		}
	})

	//test getLen
	t.Run("testGetLen", func(t *testing.T) {
		var l *Link = &Link{}
		for i := 1; i <= 100; i++ {
			l.Push(i)
			if l.GetLen() != int32(i) {
				t.Fatalf("GetLen error, expect %d, but %d", i, l.GetLen())
			}
		}
	})

	//test clear
	t.Run("testClear", func(t *testing.T) {
		var l *Link = &Link{}
		for i := 1; i <= 100; i++ {
			l.Push(i)
		}

		log.Printf("The length before clear is %d.", l.GetLen())
		len := l.Clear()
		if len != 100 {
			t.Fatalf("Expect clear 100 nodes, but %d", len)
		}
		if l.GetLen() != 0 {
			t.Fatalf("After clear the length expect 0, but %d", l.GetLen())
		}
	})

	//test getValue
	t.Run("testGetValue", func(t *testing.T) {
		var l *Link = &Link{}
		for i := 0; i < 100; i++ {
			l.Push(i)
		}

		for i := 0; i < 100; i++ {
			index := rand.Intn(100)
			value, _ := l.GetValue(int32(index)).(int)
			if index != value {
				t.Fatalf("Getvalue error expect %d, but %d.", index, value)
			}
		}
	})

	//test splice
	t.Run("testSplice", func(t *testing.T) {
		l := &Link{}
		l.Push(1)
		l.Push(2)
		l.Push(3)
		l.Push(4)
		l.Splice(5, 0, 5, 6)
		v := l.GetLen()
		if v != 6 {
			t.Fatalf("Splice error expect length 6 but %d", v)
		}
		l.Splice(0, 0, 7)
		v = l.GetLen()
		if v != 7 {
			t.Fatalf("Splice error expect length 7 but %d", v)
		}
	})
}
