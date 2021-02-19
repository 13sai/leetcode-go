package code

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println()
	fmt.Println("TestBubbleSort------")
	a := []int{9, 11, 98, 13, 41, 15, 28, 10, 2, 32, 45, 27, 20}
	fmt.Println("init", a)
	fmt.Println(BubbleSort(a))
	fmt.Println()
}

func TestInsertSort(t *testing.T) {
	fmt.Println()
	fmt.Println("InsertSort------")
	a := []int{9, 11, 98, 13, 41, 15, 28, 10, 2, 32, 45, 27, 20}
	fmt.Println("init", a)
	fmt.Println(InsertSort(a))
}

func TestSelectSort(t *testing.T) {
	fmt.Println()
	fmt.Println("SelectSort------")
	a := []int{9, 11, 98, 13, 41, 15, 28, 10, 2, 32, 45, 27, 20}
	fmt.Println("init", a)
	fmt.Println(SelectSort(a))
}

func TestMergeSort(t *testing.T) {
	fmt.Println()
	fmt.Println("MergeSort------")
	a := []int{9, 11, 98, 13, 41, 15, 28, 10, 2, 32, 45, 27, 20}
	fmt.Println("init", a)
	fmt.Println(MergeSort(a))
}

func TestQuickSort(t *testing.T) {
	fmt.Println()
	fmt.Println("QuickSort------")
	a := []int{9, 11, 98, 13, 41, 15, 28, 10, 2, 32, 45, 27, 20}
	fmt.Println("init", a)
	fmt.Println(QuickSort(a))
}

func TestHeapSort(t *testing.T) {
	fmt.Println()
	fmt.Println("HeapSort------")
	a := []int{9, 11, 98, 13, 41, 15, 28, 10, 2, 32, 45, 27, 20}
	fmt.Println("init", a)
	fmt.Println(HeapSort(a))
}
