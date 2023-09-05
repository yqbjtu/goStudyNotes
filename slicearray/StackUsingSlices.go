package stack

type StackUsingSlices struct {
	items []interface{}
}

func (sus *StackUsingSlices) Push(element interface{}) interface{} {
	sus.items = append(sus.items, element)
	return element
}

func (sus *StackUsingSlices) Pop() interface{} {
	l := len(sus.items)
	element := sus.items[l-1]
	sus.items = sus.items[:l-1]
	return element
}

func (sus *StackUsingSlices) Peek() interface{} {
	return sus.items[len(sus.items)-1]
}

func (sus *StackUsingSlices) Len() int {
	return len(sus.items)
}
