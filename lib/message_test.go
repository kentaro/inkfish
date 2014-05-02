package inkfish

import (
	"testing"
)

func TestMessage(t *testing.T) {
	message := NewMessage("join", "test", "test")

	{
		actual := message.Command
		expected := "join"

		if actual != expected {
			t.Fatalf("Message.Command is not correctly set: %v", actual)
		}
	}

	{
		actual := message.Channel
		expected := "test"

		if actual != expected {
			t.Fatalf("Message.Channel is not correctly set: %v", actual)
		}
	}

	{
		actual := message.Body
		expected := "test"

		if actual != expected {
			t.Fatalf("Message.Body is not correctly set: %v", actual)
		}
	}
}
