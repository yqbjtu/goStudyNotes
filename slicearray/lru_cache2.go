package main

import (
	"container/list"
	"fmt"
)

type DataNode struct {
	Data   int
	KeyPtr *list.Element
}

type LRUCache struct {
	Queue    *list.List
	// map中存放数据
	Items    map[int]*DataNode
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Queue: list.New(), Items: make(map[int]*DataNode), Capacity: capacity}
}

func (l *LRUCache) Get(key int) int {
	if item, ok := l.Items[key]; ok {
		// 将该数据对应的KeyPtr，移动到list的最前面
		l.Queue.MoveToFront(item.KeyPtr)
		return item.Data
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			// 从queue中删除尾部的元素
			l.Queue.Remove(back)
			// 从存放数据的map中删除尾部对应的元素
			delete(l.Items, back.Value.(int))
		}

		// 将新插入的数据对应的KeyPtr，移动到list的最前面
		l.Items[key] = &DataNode{Data: value, KeyPtr: l.Queue.PushFront(key)}
	} else {

		item.Data = value
		// 将已有的数据value更新后，更新它在map中的值
		l.Items[key] = item
		// 将已有的数据对应的KeyPtr，移动到list的最前面
		l.Queue.MoveToFront(item.KeyPtr)
	}
}

// leetcode 146. LRU 缓存, 借助"container/list" 实现，container提供了三种不同的容器，分别是heap（堆），链表（list），环(ring)
// list.List: List represents a doubly linked list 文档 https://pkg.go.dev/container/list
func main() {
	// Test case 1
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	//	[null, null, null, 1, null, -1, null, -1, 3, 4]
	fmt.Println("Test case 1")
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

	// Test case 2
	// 	["LRUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	// [null,null,null,null,null,-1,3]
	fmt.Println("Test case 2")
	obj = Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}
