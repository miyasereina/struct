package sort

import (
	"math/rand"
	"time"
)

//以下排序均按照升序排序

//插入排序
//时间复杂度O(n^2)
//空间复杂度O(1) 原地排序
//稳定，但受输入影响
//描述：在有序数组的基础上插入值
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			//待插入值为arr[i]
			//若arr[i]>arr[j]说明原地不动，否则一直向前插入
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			} else {
				break
			}
		}
	}
	return arr
}

//归并排序
//时间复杂度O(nlgn) 递归树或主定理
//空间复杂度O(n)
//稳定,与输入无关
//描述：
//分：分成n/2的两个子序列
//治：对子序列进行归并排序
//合：对排好的子序列进行合并
func MergeSort(arr []int) []int {
	//递归结束开始返回的条件
	if len(arr) <= 1 {
		return arr
	}
	//分治
	arr1 := MergeSort(arr[0 : len(arr)/2])
	arr2 := MergeSort(arr[len(arr)/2:])
	return merge(arr1, arr2)
}
func merge(arr1, arr2 []int) (arr []int) {
	arr = make([]int, 0)
	for true {
		//合并
		//合并后删除已经合并的元素，不容易越界
		if arr1[0] < arr2[0] {
			arr = append(arr, arr1[0])
			arr = deleteHead(arr1)
			if len(arr1) == 0 {
				arr = append(arr, arr2...)
				break
			}
		} else {
			arr = append(arr, arr2[0])
			arr2 = deleteHead(arr2)
			if len(arr2) == 0 {
				arr = append(arr, arr1...)
				break
			}
		}
	}
	return arr
}
func deleteHead(arr []int) []int {
	if len(arr) > 1 {
		return arr[1:]
	} else {
		return []int{}
	}
}

//快速排序
//时间复杂度O(nlgn) 递归树或主定理
//空间复杂度O(nlgn)
//不稳定,与输入有关
//描述：
//分：将数组分为两个子数组left元素全部小于arr[i]right元素全部大于arr[i]
//治：递归调用将子数组排序
//合：原址排序，治的时候已经有序，不再合并
//最好shuffle一下，快排非常依赖输入
func QuickSort(arr []int) []int {
	return divide(arr, 0, len(arr)-1)
}
func divide(arr []int, l, r int) []int {
	if l < r {
		q := partition(arr, l, r)
		divide(arr, l, q-1)
		divide(arr, q+1, r)
	}
	return arr
}

///注意这里利用的循环不变：
///1. if l<=i<=j:arr[i]<=pivot
///2. if j+1=<i<=r-1:arr[i]>pivot
func partition(arr []int, l, r int) int {
	j := l - 1
	for i := l; i < r; i++ {
		//小的在前面所以遇到小的就扔到前面去
		//基于1
		if arr[i] <= arr[r] {
			j += 1
			arr[j], arr[i] = arr[i], arr[j]
		}

	}
	//基于2
	arr[j+1], arr[r] = arr[r], arr[j+1]
	return j + 1
}

//一个简单的shuffle
func randomizedPartition(arr []int, l, r int) int {
	i := rand.Intn(r - l)
	i += l
	arr[i], arr[r] = arr[r], arr[i]
	return partition(arr, l, r)
}

//堆排序
//时间复杂度O(nlgn)
//空间复杂度O(nlgn)
//不稳定,与输入有关
//原地排序
//描述：	1.建立最大堆
//		2.将根节点放到最后
//		3.维护剩下元素为最大堆
//		4.重复2,3

func HeapSort(arr []int) []int {
	lo := 0
	hi := len(arr)
	// Build heap with greatest element at top.
	for i := hi/2 - 1; i >= 0; i-- {
		siftDown(arr, i, hi)
	}
	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		siftDown(arr, lo, i)
	}
	return arr
}

//维护最大堆（可递归，可迭代）
func siftDown(arr []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		//边界条件
		if child >= hi {
			break
		}
		//注意不要越界，选出左子节点和右子节点里的较大值
		if child+1 < hi && arr[child] < arr[child+1] {
			child++
		}
		//建堆时是是上浮的，下面的已经是堆，所以可以直接return
		if arr[root] >= arr[child] {
			return
		}
		//
		arr[root], arr[child] = arr[child], arr[root]
		root = child
	}
}

////////以下是沙雕排序，写着玩
//协程排序
//时间复杂度：O(n)
//空间复杂度：O(n)
//严重受输入影响，不稳定
func sleep(i int, ch chan int) {
	time.Sleep(time.Duration(int64(i)) * time.Second)
	ch <- i
}
func CoroutineSort(arr []int) []int {
	ans := make([]int, len(arr))
	ch := make(chan int)
	for i := 0; i < len(arr); i++ {
		go sleep(arr[i], ch)
	}
	for i := 0; i < len(arr); i++ {
		ans[i] = <-ch
	}
	return ans
}
