package main

import (
	"fmt"

	"github.com/DataStructure/sort"
	"github.com/DataStructure/utils"
)

func main() {
	s := utils.GenerateSlice(14, true)
	fmt.Println("s", s)
	s1 := utils.SliceCopy(s)
	fmt.Println("s1", sort.InsertSort(s1))
	s2 := utils.SliceCopy(s)
	fmt.Println("s2", sort.ShellSort(s2))
	s3 := utils.SliceCopy(s)
	fmt.Println("s3", sort.BubbleSort(s3))
	s4 := utils.SliceCopy(s)
	sort.QuickSort(s4, 0, len(s4)-1)
	fmt.Println("s4", s4)
	s5 := utils.SliceCopy(s)
	fmt.Println("s5", sort.SimpleSelectSort(s5))
	s6 := utils.SliceCopy(s)
	fmt.Println("s6", sort.HeapSort(s6))
	//s6 := utils.SliceCopy(s)
	//fmt.Println(sort.SimpleSelectSort(s5))
}
