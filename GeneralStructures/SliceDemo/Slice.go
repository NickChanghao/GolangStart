package main

import (
	"BasicTools/Slice"
	"fmt"
	"github.com/ecodeclub/ekit/slice"
)

func main() {
	mySlice := make([]int, 0, 5)
	for i := 0; i < 20; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Len: %d, Cap: %d\n", len(mySlice), cap(mySlice))
	}
	sumResult := Slice.SumUp[int](mySlice) // 求Slice之和
	sumResult1 := slice.Sum[int](mySlice)
	smallest := Slice.Min[int](mySlice)
	fmt.Println(sumResult)
	fmt.Println(sumResult1)
	fmt.Println(smallest)
}
