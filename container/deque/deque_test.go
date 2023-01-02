package deque

import (
	"container/list"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	stack := list.New()
	for i := 0; i < 10; i++ {
		stack.PushBack(i)
	}
	for stack.Len() > 0 {
		i := stack.Remove(stack.Back()).(int)
		fmt.Println(i)
	}
}

func TestQueue1(t *testing.T) {
	deque := NewArrayDeque()
	list := list.New()
	for i := 0; i < 10000; i++ {
		deque.Offer(i)
		list.PushBack(i)
	}
	if deque.Size() != list.Len() {
		t.Errorf("deque size = %d, list len = %d", deque.Size(), list.Len())
	}
	for !deque.IsEmpty() {
		poll := deque.Poll().(int)
		i := list.Remove(list.Front()).(int)
		if i != poll {
			t.Errorf("错误, poll = %d, i = %d", poll, i)
		}
	}
	if deque.Size() != list.Len() {
		t.Errorf("deque size = %d, list len = %d", deque.Size(), list.Len())
	}
}

func TestAsStack(t *testing.T) {
	defer func() {
		i := recover()
		t.Log(i)
	}()
	deque := NewArrayDeque()
	deque.Push(nil)
}

func TestFor(t *testing.T) {
	deque := NewArrayDeque()
	deque.Push(1)
	deque.Push(2)
	deque.Push(3)
	deque.OfferFirst(4)
	deque.Push(6)
	inter := deque.iterator()
	for inter.HasNext() {
		e := inter.Next()
		t.Log(e)
	}
}
