/ Go program to delete the element from the given slice
package main
 
import "fmt"
 
// go 1.21 版本 slices包提供新newSlice := slices.Delete(mySlice, from, to)
func delete_at_index(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}
 
func main() {
     
    numbers := []int{10, 20, 30, 40, 50, 90, 60}
    fmt.Println("Original Slice:", numbers)
 
    var index int = 3
     
    // Get the element at a given index in the slice
    elem := numbers[index]
    //  返回新的slice 
    numbers = delete_at_index(numbers, index)
 
    fmt.Printf("The element %d was deleted.\n", elem)
    fmt.Println("Slice after deleting element:", numbers)
}
