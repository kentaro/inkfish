package inkfish

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	messageQueue := make(chan *Message)
	quitQueue := make(chan bool, 1)
	queue := NewQueue(messageQueue, quitQueue)

	{
		go func() { queue.Push(NewMessage("join", "test")) }()
		actual := queue.Pop()
		expected := NewMessage("join", "test")

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Message is not correctly enqueued: %+v", actual)
		}
	}
}
