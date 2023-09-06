package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	list *list.List
}

// New returns new construct queue
func New() *Queue {
	return &Queue{list.New()}
}

// Push inserts element to the queue
func (queue *Queue) Push(value interface{}) {
	queue.list.PushBack(value)
}

// Front returns first element of the queue
func (queue *Queue) Front() interface{} {
	it := queue.list.Front()
	if it != nil {
		return it.Value
	}
	return nil
}

// Back returns last element of the queue
func (queue *Queue) Back() interface{} {
	it := queue.list.Back()
	if it != nil {
		return it.Value
	}
	return nil
}

// Pop returns and deletes first element of the queue
func (queue *Queue) Pop() interface{} {
	it := queue.list.Front()
	if it != nil {
		queue.list.Remove(it)
		return it.Value
	}
	return nil
}

// Size returns size of the queue
func (queue *Queue) Size() int {
	return queue.list.Len()
}

// Empty returns whether queue is empty
func (queue *Queue) Empty() bool {
	return queue.list.Len() == 0
}

// Clear clears the queue
func (queue *Queue) Clear() {
	for !queue.Empty() {
		queue.Pop()
	}
}

func main() {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	size := q.Size()
	fmt.Printf("queue 1,2,3,4 \n")
	if size != 4 {
		fmt.Errorf("queue.Size() test failed! Got %d, expected 4.\n", size)
	}
	value := q.Front().(int)
	if value != 1 {
		fmt.Errorf("queue.Front() test failed! Got %d, expected 1.\n", value)
	} else {
		fmt.Printf("queue.Front() as expected 1, queue:1,2,3,4\n")
	}

	value = q.Pop().(int)
	if value != 1 {
		fmt.Errorf("queue.Pop() test failed! Got %d, expected 1.\n", value)
	} else {
		fmt.Printf("queue.Pop() as expected 1, queue:2,3,4\n")
	}

	size = q.Size()
	if size != 3 {
		fmt.Errorf("queue.Size() test failed! Got %d, expected 3.\n", size)
	}

	value = q.Pop().(int)
	if value != 2 {
		fmt.Errorf("queue.Pop() test failed! Got %d, expected 2.\n", value)
	}
	value = q.Back().(int)
	if value != 4 {
		fmt.Errorf("queue.Back() test failed! Got %d, expected 4.\n", value)
	}
	empty := q.Empty()
	if empty {
		fmt.Errorf("queue.Empty() test failed! Got %t, expected false.\n", empty)
	}
	value = q.Pop().(int)
	if value != 3 {
		fmt.Errorf("queue.Pop() test failed! Got %d, expected 3.\n", value)
	}
	empty = q.Empty()
	if empty {
		fmt.Errorf("queue.Empty() test failed! Got %t, expected false.\n", empty)
	}
	value = q.Pop().(int)
	if value != 4 {
		fmt.Errorf("queue.Pop() test failed! Got %d, expected 4.\n", value)
	}
	empty = q.Empty()
	if !empty {
		fmt.Errorf("queue.Empty() test failed! Got %t, expected true.\n", empty)
	}
	nilValue := q.Front()
	if nilValue != nil {
		fmt.Errorf("queue.Front() test failed! Got %d, expected nil.", nilValue)
	}
	nilValue = q.Back()
	if nilValue != nil {
		fmt.Errorf("queue.Back() test failed! Got %d, expected nil.", nilValue)
	}
	nilValue = q.Pop()
	if nilValue != nil {
		fmt.Errorf("queue.Pop() test failed! Got %d, expected nil.", nilValue)
	}
	size = q.Size()
	if size != 0 {
		fmt.Errorf("queue.Size() test failed! Got %d, expected 0.", size)
	}
}
