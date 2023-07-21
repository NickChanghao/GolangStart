package main

import (
	"fmt"
	"github.com/ecodeclub/ekit/slice"
)

func main() {
	mySlice := make([]int, 0, 5)
	for i := 0; i < 20; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Len: %d, Cap: %d\n", len(mySlice), cap(mySlice))
	}
	sumResult := slice.Sum[int](mySlice)
	fmt.Println(sumResult)
}
