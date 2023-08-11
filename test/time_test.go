package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_time(t *testing.T) {
	list1 := []int64{1, 2, 5, 7, 7}
	list2 := []int64{7, 5, 9, 2, 1, 2, 5, 7}
	list := findIntersection(list1, list2)
	fmt.Println(list)
}

func findIntersection(list1, list2 []int64) []int64 {
	// 先对两个切片进行排序
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var intersection []int64
	i, j := 0, 0

	// 遍历两个排序后的切片，找到交集
	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			i++
		} else if list1[i] > list2[j] {
			j++
		} else {
			// 找到交集元素，加入结果切片
			intersection = append(intersection, list1[i])
			i++
			j++
		}
	}

	return intersection
}
