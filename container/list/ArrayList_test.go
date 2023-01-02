package list

import (
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	var arrayList = NewArrayList()
	for i := 0; i < 100; i++ {
		arrayList.Add(i)
	}
	if arrayList.Size() != 100 {
		t.Errorf("expceptig size = %d, bug got size = %d", 100, arrayList.Size())
	}
}

func TestGet(t *testing.T) {
	var arrayList = NewArrayList()
	for i := 0; i < 10; i++ {
		arrayList.Add(i)
	}
	for i := 0; i < 10; i++ {
		get := arrayList.Get(i).(int)
		if i != get {
			t.Errorf("execpting i == get, bug i = %d, get = %d", i, get)
		}
	}
}

func TestArrayList_Remove(t *testing.T) {
	var arrayList = NewArrayList()
	var N = 100
	for i := 0; i < N; i++ {
		arrayList.Add(i)
	}
	randInt := rand.Intn(N)
	remove := arrayList.Remove(randInt)
	if remove == nil {
		t.Errorf("execptiong remove must not be nil, bug got a nil")
	}
	arrayList.Remove(-1)
}

func TestArrayList_RemoveElement(t *testing.T) {
	var arrayList = NewArrayListWithCap(1000)
	var N = 100
	for i := 0; i < N; i++ {
		arrayList.Add(i)
	}
	element := arrayList.RemoveElement(10)
	if element {
		t.Logf("expecting remove element 10, bot remove fail")
	}
}
