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

func (h *Holidays) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	holidays := &Holidays{}
	heap.Init(holidays)
	heap.Push(holidays, Holiday{name: "Christmas", date: time.Date(2023, time.December, 25, 0, 0, 0, 0, time.Local)})
	heap.Push(holidays, Holiday{name: "Labour Day", date: time.Date(2023, time.May, 1, 0, 0, 0, 0, time.Local)})
	heap.Push(holidays, Holiday{name: "Diwali", date: time.Date(2023, time.November, 23, 0, 0, 0, 0, time.Local)})

	fmt.Println("holidays: ", holidays)

	holidays.Pop()
	fmt.Println("after pop one, holidays: ", holidays)

	holidays.Pop()
	fmt.Println("after pop two, holidays: ", holidays)
}

//运行结果
//holidays:  &[[Labour Day, May 1] [Christmas, Dec 25] [Diwali, Nov 23]]
//after pop one, holidays:  &[[Labour Day, May 1] [Christmas, Dec 25]]
//after pop two, holidays:  &[[Labour Day, May 1]]
