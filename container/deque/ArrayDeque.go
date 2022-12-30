package deque

import "x-tool/util/arrays"

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

func (a *ArrayDeque) addFist(e interface{}) {
	a.head = (a.head - 1) & (cap(a.elements) - 1)
	a.elements[a.head] = e
	if a.head == a.tail {
		a.doubleCapacity()
	}
}

func (a *ArrayDeque) addLast(e interface{}) {
	a.elements[a.tail] = e
	a.tail++
	if a.tail&(cap(a.elements)-1) == a.head {
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
	p := a.head
	n := cap(a.elements)
	r := n - p
	newCapacity := n << 1
	newArr := make([]interface{}, newCapacity)
	arrays.CopyArr(a.elements, p, newArr, 0, r)
	arrays.CopyArr(a.elements, 0, newArr, r, p)
	a.elements = newArr
	a.head = 0
	a.tail = n
}
