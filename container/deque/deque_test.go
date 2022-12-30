package deque

import "testing"

func TestQueue(t *testing.T) {
	queue := NewArrayDeque()
	for i := 0; i < 10; i++ {
		queue.Offer(i)
	}
	t.Logf("elements: %v, head: %d, tail: %d", queue.elements, queue.head, queue.tail)

}
