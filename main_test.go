package main

import (
	"reflect"
	"runtime"
	"testing"
)

var (
	nums1000  = makeSeqNums(1000)
	nums10000 = makeSeqNums(10000)
	nums50000 = makeSeqNums(50000)

	search3     = []int{1000, 1000, 1000}
	search100   = makeSeqNums(100)
	search1000  = makeSeqNums(1000)
	search10000 = makeSeqNums(10000)
	search50000 = makeSeqNums(50000)

	search3NotFound     = []int{-1, -1, -1}
	search100NotFound   = makeSpecificNums(-1, 100)
	search1000NotFound  = makeSpecificNums(-1, 1000)
	search10000NotFound = makeSpecificNums(-1, 10000)
	search50000NotFound = makeSpecificNums(-1, 50000)
)

// 順番に並んだ数値のスライスを作成
func makeSeqNums(count int) []int {
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = i + 1
	}
	return nums
}

// 順番に並んだ数値のスライスを作成
func makeSeqNumsFromZero(count int) []int {
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = i
	}
	return nums
}

// 順番に並んだ数値のスライスを作成
func makeSpecificNums(num, count int) []int {
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = num
	}
	return nums
}

type args struct {
	nums     []int
	searches []int
}

type test struct {
	name    string
	args    args
	wantPos []int
}

type tests []test

var testPattern = tests{
	{
		name: "found_all",
		args: args{
			nums:     nums1000,
			searches: []int{1, 500, 1000},
		},
		wantPos: []int{0, 499, 999},
	},
	{
		name: "found_all_10000",
		args: args{
			nums:     nums10000,
			searches: nums10000,
		},
		wantPos: makeSeqNumsFromZero(10000),
	},
	{
		name: "include_not_found",
		args: args{
			nums:     nums1000,
			searches: []int{0, 500, 1001},
		},
		wantPos: []int{-1, 499, -1},
	},
	{
		name: "not_found_all",
		args: args{
			nums:     nums1000,
			searches: []int{0, -1, 1001},
		},
		wantPos: []int{-1, -1, -1},
	},
	{
		name: "not_found_all_10000",
		args: args{
			nums:     nums10000,
			searches: search10000NotFound,
		},
		wantPos: makeSpecificNums(-1, 10000),
	},
}

func Test_LinearSearch(t *testing.T) {
	commonTest(t, LinearSearch)
}

func Test_BinarySearch(t *testing.T) {
	commonTest(t, BinarySearch)
}

func Test_MapSearch(t *testing.T) {
	commonTest(t, MapSearch)
}

func commonTest(t *testing.T, f searchFunc) {
	fv := reflect.ValueOf(f)
	name := runtime.FuncForPC(fv.Pointer()).Name()

	for _, tt := range testPattern {
		t.Run(tt.name, func(t *testing.T) {
			if gotPos := f(tt.args.nums, tt.args.searches); !reflect.DeepEqual(gotPos, tt.wantPos) {
				t.Errorf(name+"() = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}
