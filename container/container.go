package container

type Container interface {
	Size() int

	Get(index int) interface{}

	Add(e interface{})

	Remove(index int) interface{}

	Set(index int, e interface{}) interface{}
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
