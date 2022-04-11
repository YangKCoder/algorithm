package heap

import (
	"fmt"
)

func ExampleHeapSort() {
	arr := []int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0}
	heapSort(arr)
	fmt.Println(arr)
	// output:
	// [0 1 2 3 5 6 7 8 9 10]
}
