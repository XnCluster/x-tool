package list

import (
	"fmt"
	"x-tool/container"
	"x-tool/util/arrays"
)

type ArrayList struct {
	elements []interface{}
	size     int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		elements: make([]interface{}, 0),
		size:     0,
	}
}

func NewArrayListWithCap(capacity int) *ArrayList {
	return &ArrayList{
		elements: make([]interface{}, capacity),
		size:     0,
	}
}

func NewArrayListWithContainer(container container.Container) *ArrayList {
	size := container.Size()
	list := ArrayList{
		elements: make([]interface{}, size),
		size:     0,
	}
	for i := 0; i < size; i++ {
		list.Add(container.Get(i))
	}
	return &list
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) Get(index int) interface{} {
	if index > a.Size() || index < 0 {
		return nil
	}
	return a.elements[index]
}

func (a *ArrayList) Add(e interface{}) {
	a.elements = append(a.elements, e)
	a.size++
}

func (a *ArrayList) Remove(index int) interface{} {
	a.checkRange(index)
	oldValue := a.element(index)
	a.removeFast(index)
	return oldValue
}

func (a *ArrayList) RemoveElement(e interface{}) bool {
	if e == nil {
		for i := 0; i < a.Size(); i++ {
			if a.elements[i] == nil {
				a.removeFast(i)
				return true
			}
		}
	} else {
		for i := 0; i < a.Size(); i++ {
			if a.elements[i] == e {
				a.removeFast(i)
				return true
			}
		}
	}
	return false
}

func (a *ArrayList) Set(index int, e interface{}) {
	a.checkRange(index)

}

func (a *ArrayList) removeFast(index int) {
	a.checkRange(index)
	numMoved := a.Size() - index - 1
	if numMoved > 0 {
		arrays.CopyArr(a.elements, index+1, a.elements, index, numMoved)
	}
	a.size--
	a.elements[a.size] = nil
}

func (a *ArrayList) element(index int) interface{} {
	return a.elements[index]
}

func (a *ArrayList) checkRange(index int) {
	if index < 0 || index >= a.Size() {
		errorMsg := fmt.Sprintf("index out of range index:%d, size: %d", index, a.Size())
		panic(errorMsg)
	}
}
