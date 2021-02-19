package code

import "fmt"

/**
 * 冒泡排序
 * 说明：就是第一个位置上的数与他相邻第二个位置上的数比较，
 * 如果比他相邻的数大，则两者交换位置，否则不交换。
 * 接着第二个位置上的数与第三个位置上的数比较大小，也是大则交换，
 * 一直到和最后一个位置的数比较交换完毕。
 * 然后，是下一个循环，就是第二个位置上的数重复上面的比较交换操作，
 * 直到把整个数组变成是一个从小到大的有序序列。
 */
func BubbleSort(arr []int) []int {
	l := len(arr)

	for i := 0; i < l; i++ {
		// 标记此轮是否有交换，没有表示已经有序
		btn := true
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				btn = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)

		if btn {
			break
		}
	}

	return arr
}

/**
 * 插入排序
 * 说明：从待排序的数组中选出来一个最小值(可以认为第一个数就是已排序的数组)，
 * 然后从剩余数中选出来最小值有序放到已排序的数组中，
 * 依次操作，直到最后的数组都是一个从小到大的有序数组为止
 */
func InsertSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		v := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > v {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = v
		fmt.Println(arr)
	}

	return arr
}

/**
 * 选择排序
 * 说明：从待排序的数组中选出来一个最小值，放到新的数组的第一个位置，
 * 继续从剩余的数组中选取最小值放入到数组中，重复上面的步骤，将数字都取出来排成新的有序数组
 */
func SelectSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		tmp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = tmp
		fmt.Println(arr)
	}

	return arr
}

/**
 * 归并排序
 */
func MergeSort(arr []int) []int {
	l := len(arr)
	if l < 2 {
		return arr
	}
	num := l / 2
	left := MergeSort(arr[:num])
	right := MergeSort(arr[num:])
	return merge(left, right)
}

// 有点双指针的意思
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	if l < len(left) {
		result = append(result, left[l:]...)
	}

	if r < len(right) {
		result = append(result, right[r:]...)
	}
	fmt.Println(result)
	return
}

/**
 * 快速排序
 */
func QuickSort(arr []int) []int {
	return separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) []int {
	if start >= end {
		return arr
	}

	i := partition(arr, start, end)
	arr = separateSort(arr, start, i-1)
	arr = separateSort(arr, i+1, end)
	return arr
}

// 分成左右两部分，进行比较，大的在右边，小的在左边，并最终返回分隔位置i
func partition(arr []int, start, end int) int {
	// 选取最后一位为分隔基准值
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[j], arr[i] = arr[i], arr[j]
			}
			i++
		}
	}

	// 基准和i交换位置
	arr[i], arr[end] = arr[end], arr[i]
	fmt.Println(arr)
	return i
}


//堆的特性
//1.是完全二叉树
//2.每一个节点都小于子节点(最小堆)
//3.根节点最小
//4.左子节点2i+1, 右子节点2i+2, 父节点(i-1)/2
func HeapSort(a []int) []int {
	heapLength := len(a)
	a = buildHeap(a, heapLength)
    fmt.Println(a)
    for i := len(a) - 1; i >= 0; i-- {
        a[i], a[0] = a[0], a[i]
        heapLength--
        a = minHeap(a, 0, heapLength)
	}
	return a
}

// 最小堆
// 大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]  
// 小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]  
func parentNode(i int)  int{
    return (i - 1)/2
}
 
//左节点
func leftNode(i int) int{
    return 2*i + 1
}
//右节点
func rightNode(i int) int{
    return 2*i + 2
}
//创建heap
func buildHeap(heap []int, length int) []int{
    for i := length/2 - 1; i >= 0; i-- {
        heap = minHeap(heap, i, length)
    }
    return heap
}
 
func minHeap(heap []int, i int, length int) []int{
    left := leftNode(i)
    right := rightNode(i)
    largest := 0
    if left < length && heap[left] > heap[i] {
        largest = left
    } else {
        largest = i
    }
    if right < length && heap[right] > heap[largest] {
        largest = right
    }
    if largest != i {
        heap[i], heap[largest] = heap[largest], heap[i]
        //需要继续比较其父节点
        heap = minHeap(heap, largest, length)
	}
	// fmt.Println(heap)
	return heap
}