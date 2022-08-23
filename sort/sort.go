package sort

import "fmt"

//以下排序均按照升序排序

//插入排序
func Insert(ls []int) {
	if len(ls) <= 1 {
		return
	}
	for i := 1; i < len(ls); i++ {
		for j := i - 1; j >= 0; j-- {
			if ls[j+1] < ls[j] {
				ls[j+1], ls[j] = ls[j], ls[j+1]
			} else {
				break
			}
		}
	}
	ls = append(ls, 9999)
	fmt.Println(ls)

}

//归并排序
func Merge(ls []int) []int {
	if len(ls) <= 1 {
		return ls
	}
	ls1 := Merge(ls[0 : len(ls)/2])
	ls2 := Merge(ls[len(ls)/2:])
	return merge(ls1, ls2)
}
func merge(ls1, ls2 []int) (ls []int) {
	ls = make([]int, 0)
	for true {
		if ls1[0] < ls2[0] {
			ls = append(ls, ls1[0])
			ls1 = deleteHead(ls1)
			if len(ls1) == 0 {
				ls = append(ls, ls2...)
				break
			}
		} else {
			ls = append(ls, ls2[0])
			ls2 = deleteHead(ls2)
			if len(ls2) == 0 {
				ls = append(ls, ls1...)
				break
			}
		}
	}
	return ls
}
func deleteHead(ls []int) []int {
	if len(ls) > 1 {
		return ls[1:]
	} else {
		return []int{}
	}
}
