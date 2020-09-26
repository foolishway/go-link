package link

import "testing"

func TestLink(t *testing.T) {
	//test push
	t.Run("testPush", func(t *testing.T) {
		var l *link = &link{}
		l.push(&node{value: "1"})
		if value, ok := l.start.value.(string); !ok {
			t.Fatal("Type error, want string but not.")
			if value != "1" {
				t.Fatalf("Value error, want %q but %s", "1", value)
			}
		}
		l.push(&node{value: 2})
		if value, ok := l.start.next.value.(int); !ok {
			t.Fatal("Type error, want int but not.")
			if value != 2 {
				t.Fatalf("Value error, want %d but %d", 2, value)
			}
		}
		l.push(&node{value: true})
		if value, ok := l.start.next.next.value.(bool); !ok {
			t.Fatal("Type error, want boolean but not.")
			if value != true {
				t.Fatalf("Value error, want %v but %v", true, value)
			}
		}
	})

	//test pop
}
