package main

import (
	"container/heap"
	"fmt"
	"time"
)

// Each holiday has a name and a date
type Holiday struct {
	name string
	date time.Time
}

// for convenience, we can create a string representation to print the holiday instance
func (h Holiday) String() string {
	return "[" + h.name + ", " + h.date.Format("Jan 2") + "]"
}

// We can create a new type to represent a list of holidays
type Holidays []Holiday

// Most of the changes here would be in the `Less` method
func (h Holidays) Less(i, j int) bool {
	// Here, we choose to order by comparing the calendar dates of
	// two holiday instances
	return h[i].date.Before(h[j].date)
}

// The remaining methods are more or less the same as the previous examples
func (h Holidays) Len() int { return len(h) }

func (h Holidays) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Holidays) Push(x interface{}) {
	*h = append(*h, x.(Holiday))
}

// remove and return element Len() - 1
func (h *Holidays) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	holidayHeap1 := myHeap()
	fmt.Println("holidayHeap1: ", holidayHeap1)
	fmt.Println("holidayHeap1.Pop(), get the last item from heap")
	for holidayHeap1.Len() > 0 {
		fmt.Printf("%v \n", holidayHeap1.Pop())
	}
	fmt.Println()

	holidayHeap2 := myHeap()
	fmt.Println("holidayHeap2: ", holidayHeap2)
	holidayHeap2.Pop()
	fmt.Println("after pop one, holidayHeap2: ", holidayHeap2)

	holidayHeap2.Pop()
	fmt.Println("after pop two, holidayHeap2: ", holidayHeap2)
	fmt.Println()

	holidayHeap3 := myHeap()
	fmt.Println("holidayHeap3: ", holidayHeap3)

	// func heap.Pop(h Interface) any Pop removes and returns the minimum element (according to Less) from the heap.
	// The complexity is O(log n) where n = h.Len(). Pop is equivalent to Remove(h, 0).

	fmt.Println("heap.Pop(holidayHeap3), get the the minimum element from heap")
	for holidayHeap3.Len() > 0 {
		fmt.Printf("%v \n", heap.Pop(holidayHeap3))
	}
}

func myHeap() *Holidays {
	holidayHeap := &Holidays{}
	heap.Init(holidayHeap)
	heap.Push(holidayHeap, Holiday{name: "Christmas", date: time.Date(2023, time.December, 25, 0, 0, 0, 0, time.Local)})
	heap.Push(holidayHeap, Holiday{name: "Labour Day", date: time.Date(2023, time.May, 1, 0, 0, 0, 0, time.Local)})
	heap.Push(holidayHeap, Holiday{name: "Diwali", date: time.Date(2023, time.November, 23, 0, 0, 0, 0, time.Local)})

	return holidayHeap
}

/*运行结果
holidayHeap1:  &[[Labour Day, May 1] [Christmas, Dec 25] [Diwali, Nov 23]]
holidayHeap1.Pop(), get the last item from heap
[Diwali, Nov 23]
[Christmas, Dec 25]
[Labour Day, May 1]

holidayHeap2:  &[[Labour Day, May 1] [Christmas, Dec 25] [Diwali, Nov 23]]
after pop one, holidayHeap2:  &[[Labour Day, May 1] [Christmas, Dec 25]]
after pop two, holidayHeap2:  &[[Labour Day, May 1]]

holidayHeap3:  &[[Labour Day, May 1] [Christmas, Dec 25] [Diwali, Nov 23]]
heap.Pop(holidayHeap3), get the the minimum element from heap
[Labour Day, May 1]
[Diwali, Nov 23]
[Christmas, Dec 25]
*/
