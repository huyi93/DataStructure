package main

import (
	"fmt"

	"github.com/DataStructure/sort"
	"github.com/DataStructure/utils"
)

func main() {
	s := utils.GenerateSlice(13, true)
	fmt.Println(sort.HeapSort(s))
	a := map[string]int{}
	fmt.Println(a["w"])
}
