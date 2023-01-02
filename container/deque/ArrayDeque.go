package deque

import (
	"fmt"
	"x-tool/util/arrays"
)

type ArrayDeque struct {
	elements   []interface{}
	head, tail int
}

const minCapacity = 8

func NewArrayDeque() *ArrayDeque {
	return &ArrayDeque{
		elements: make([]interface{}, minCapacity),
		head:     0,
		tail:     0,
	}
}

func (a *ArrayDeque) Push(e interface{}) {
	a.addLast(e)
}

func (a *ArrayDeque) Pop() interface{} {
	return a.removeLast()
}

func (a *ArrayDeque) Offer(e interface{}) {
	a.addLast(e)
}

func (a *ArrayDeque) Poll() interface{} {
	return a.removeFirst()
}

func (a *ArrayDeque) Peek() interface{} {
	return a.getLast()
}

func (a *ArrayDeque) PeekFirst() interface{} {
	return a.getFist()
}

func (a *ArrayDeque) OfferFirst(e interface{}) {
	a.addFist(e)
}

func (a *ArrayDeque) PollLast() interface{} {
	return a.removeLast()
}

func (a *ArrayDeque) IsEmpty() bool {
	return a.head == a.tail
}

func (a *ArrayDeque) Size() int {
	return (a.tail - a.head) & (cap(a.elements) - 1)
}

func (a *ArrayDeque) addFist(e interface{}) {
	if e == nil {
		panic(fmt.Sprintf("null point e: %v", e))
	}
	a.head = (a.head - 1) & (cap(a.elements) - 1)
	a.elements[a.head] = e
	if a.head == a.tail {
		a.doubleCapacity()
	}
}

func (a *ArrayDeque) addLast(e interface{}) {
	if e == nil {
		panic(fmt.Sprintf("null point e: %v", e))
	}
	a.elements[a.tail] = e
	a.tail = (a.tail + 1) & (cap(a.elements) - 1)
	if a.tail == a.head {
		a.doubleCapacity()
	}
}

func (a *ArrayDeque) getFist() interface{} {
	h := a.head
	result := a.elements[h]
	return result
}

func (a *ArrayDeque) getLast() interface{} {
	result := a.elements[(a.tail-1)&(cap(a.elements)-1)]
	return result
}

func (a *ArrayDeque) removeFirst() interface{} {
	h := a.head
	result := a.elements[h]
	a.elements[h] = nil
	a.head = (h + 1) & (cap(a.elements) - 1)
	return result
}

func (a *ArrayDeque) removeLast() interface{} {
	a.tail = (a.tail - 1) & (cap(a.elements) - 1)
	result := a.elements[a.tail]
	a.elements[a.tail] = nil
	return result
}

func (a *ArrayDeque) doubleCapacity() {
	if a.head != a.tail {
		panic("ArrayDeque grow fail")
	}
	p := a.head
	n := cap(a.elements)
	r := n - p
	newCapacity := n << 1
	if newCapacity < 0 {
		panic("Sorry, ArrayDeque too big")
	}
	newArr := make([]interface{}, newCapacity)
	arrays.CopyArr(a.elements, p, newArr, 0, r)
	arrays.CopyArr(a.elements, 0, newArr, r, p)
	a.elements = newArr
	a.head = 0
	a.tail = n
}

func (a *ArrayDeque) iterator() *DeqIterator {
	return newIter(a)
}

func newIter(a *ArrayDeque) *DeqIterator {
	return &DeqIterator{a, a.head, a.tail, -1}
}

type DeqIterator struct {
	deque                  *ArrayDeque
	cursor, fence, lastRet int
}

func (i *DeqIterator) HasNext() bool {
	return i.cursor != i.fence
}

func (i *DeqIterator) Next() interface{} {
	if i.cursor == i.fence {
		panic("")
	}
	deque := i.deque
	res := deque.elements[i.cursor]
	if deque.tail != i.fence || res == nil {
		panic("")
	}
	i.lastRet = i.cursor
	i.cursor = (i.cursor + 1) & (cap(deque.elements) - 1)
	return res
}
