package deque

type MyCircularQueue struct {
	elements             []int
	head, tail, capacity int
}

func Constructor(k int) MyCircularQueue {
	capp := calculateSize(k)
	return MyCircularQueue{
		elements: make([]int, capp),
		head:     0,
		tail:     0,
		capacity: k,
	}
}

func calculateSize(k int) int {
	initCapacity := 8
	if k >= initCapacity {
		initCapacity = k
		initCapacity |= initCapacity >> 1
		initCapacity |= initCapacity >> 2
		initCapacity |= initCapacity >> 4
		initCapacity |= initCapacity >> 8
		initCapacity |= initCapacity >> 16
		initCapacity++
		if initCapacity < 0 {
			initCapacity >>= 1
		}
	}
	return initCapacity
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.elements[this.tail] = value
	this.tail = (this.tail + 1) & (cap(this.elements) - 1)
	if this.tail == this.head {
		this.grow()
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) & (cap(this.elements) - 1)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	t := (this.tail - 1) & (cap(this.elements) - 1)
	return this.elements[t]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Size() == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Size() == this.capacity
}

func (this *MyCircularQueue) Size() int {
	return (this.tail - this.head) & (cap(this.elements) - 1)
}

func (this *MyCircularQueue) grow() {
	p := this.head
	n := cap(this.elements)
	r := n - p
	newCapacity := n << 1
	a := make([]int, newCapacity)
	copyArray(this.elements, p, a, 0, r)
	copyArray(this.elements, 0, a, r, p)
	this.elements = a
	this.head = 0
	this.tail = n
}

func copyArray(src []int, srcPos int, dest []int, desPos int, len int) {
	for i := srcPos; i < len+srcPos; i++ {
		dest[desPos] = src[i]
		desPos++
	}
}
