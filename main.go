package main

import "sort"

type searchFunc func(nums, searches []int) (pos []int)
type mapSearchFunc func(m map[int]int, searches []int) (pos []int)

// LinearSearch 線形探索
func LinearSearch(nums, searches []int) (pos []int) {
	pos = make([]int, len(searches))
	for si, s := range searches {
		idx := sort.Search(len(nums), func(i int) bool { return nums[i] >= s })
		if idx < len(nums) && nums[idx] == s {
			pos[si] = idx
		} else {
			pos[si] = -1
		}
	}
	return pos
}

// BinarySearch 二分探索
func BinarySearch(nums, searches []int) (pos []int) {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return binarySearch(nums, searches)
}

// numsはソート済み
func binarySearch(nums, searches []int) (pos []int) {
	pos = make([]int, len(searches))
	for si, s := range searches {
		idx := sort.Search(len(nums), func(i int) bool { return nums[i] >= s })
		if idx < len(nums) && nums[idx] == s {
			pos[si] = idx
		} else {
			pos[si] = -1
		}
	}
	return pos
}

// MapSearch ハッシュ(Map)探索
func MapSearch(nums, searches []int) (pos []int) {
	m := sliceToMap(nums)
	return mapSearch(m, searches)
}

func mapSearch(m map[int]int, searches []int) (pos []int) {
	pos = make([]int, len(searches))
	for si, s := range searches {
		if idx, ok := m[s]; ok {
			pos[si] = idx
		} else {
			pos[si] = -1
		}
	}
	return pos
}

func sliceToMap(nums []int) map[int]int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}
	return m
}
