package deque

type Deque interface {
	// Push ,push an element to the stack
	Push(e interface{})
	// Pop ,pop an element from the stack
	Pop() interface{}
	// Offer ,offer an element to the queue
	Offer(e interface{})
	// Poll , poll an element from the queue
	Poll() interface{}
	// Peek , look at the top of stack element
	Peek() interface{}
	// PeekFirst , retrieves, but not remove, the first element of queue
	PeekFirst() interface{}
	// IsEmpty , return is emptny
	IsEmpty() bool

	OfferFirst(e interface{})

	PollLast() interface{}
}
