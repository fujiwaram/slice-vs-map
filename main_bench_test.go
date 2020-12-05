package main

import (
	"sort"
	"testing"
)

func commonBenchmark(b *testing.B, nums, searches []int, f searchFunc) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(nums, searches)
	}
}

func commonBenchmarkFroBinarySearch(b *testing.B, nums, searches []int, f searchFunc) {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(nums, searches)
	}
}

func commonBenchmarkFroMapSearch(b *testing.B, nums, searches []int, f mapSearchFunc) {
	m := sliceToMap(nums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(m, searches)
	}
}

//
// data-1000, search-3
//

func Benchmark_LinearSearch_1000__3(b *testing.B) {
	commonBenchmark(b, nums1000, search3, LinearSearch)
}

func Benchmark_BinarySearch_1000__3(b *testing.B) {
	commonBenchmark(b, nums1000, search3, BinarySearch)
}

func Benchmark_binarySearch_1000__3(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums1000, search3, binarySearch)
}

func Benchmark_MapSearch____1000__3(b *testing.B) {
	commonBenchmark(b, nums1000, search3, MapSearch)
}

func Benchmark_mapSearch____1000__3(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums1000, search3, mapSearch)
}

//
// data-1000, search-100
//

func Benchmark_LinearSearch_1000__100(b *testing.B) {
	commonBenchmark(b, nums1000, search100, LinearSearch)
}

func Benchmark_BinarySearch_1000__100(b *testing.B) {
	commonBenchmark(b, nums1000, search100, BinarySearch)
}

func Benchmark_binarySearch_1000__100(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums1000, search100, binarySearch)
}

func Benchmark_MapSearch____1000__100(b *testing.B) {
	commonBenchmark(b, nums1000, search100, MapSearch)
}

func Benchmark_mapSearch____1000__100(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums1000, search100, mapSearch)
}

//
// data-1000, search-1000
//

func Benchmark_LinearSearch_1000__1000(b *testing.B) {
	commonBenchmark(b, nums1000, search1000, LinearSearch)
}

func Benchmark_BinarySearch_1000__1000(b *testing.B) {
	commonBenchmark(b, nums1000, search1000, BinarySearch)
}

func Benchmark_binarySearch_1000__1000(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums1000, search1000, binarySearch)
}

func Benchmark_MapSearch____1000__1000(b *testing.B) {
	commonBenchmark(b, nums1000, search1000, MapSearch)
}

func Benchmark_mapSearch____1000__1000(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums1000, search1000, mapSearch)
}

//
// data-10000, search-3
//

func Benchmark_LinearSearch_10000_3(b *testing.B) {
	commonBenchmark(b, nums10000, search3, LinearSearch)
}

func Benchmark_BinarySearch_10000_3(b *testing.B) {
	commonBenchmark(b, nums10000, search3, BinarySearch)
}

func Benchmark_binarySearch_10000_3(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums10000, search3, binarySearch)
}

func Benchmark_MapSearch____10000_3(b *testing.B) {
	commonBenchmark(b, nums10000, search3, MapSearch)
}

func Benchmark_mapSearch____10000_3(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums10000, search3, mapSearch)
}

//
// data-10000, search-100
//

func Benchmark_LinearSearch_10000_100(b *testing.B) {
	commonBenchmark(b, nums10000, search100, LinearSearch)
}

func Benchmark_BinarySearch_10000_100(b *testing.B) {
	commonBenchmark(b, nums10000, search100, BinarySearch)
}

func Benchmark_binarySearch_10000_100(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums10000, search100, binarySearch)
}

func Benchmark_MapSearch____10000_100(b *testing.B) {
	commonBenchmark(b, nums10000, search100, MapSearch)
}

func Benchmark_mapSearch____10000_100(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums10000, search100, mapSearch)
}

//
// data-10000, search-1000
//

func Benchmark_LinearSearch_10000_1000(b *testing.B) {
	commonBenchmark(b, nums10000, search1000, LinearSearch)
}

func Benchmark_BinarySearch_10000_1000(b *testing.B) {
	commonBenchmark(b, nums10000, search1000, BinarySearch)
}

func Benchmark_binarySearch_10000_1000(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums10000, search1000, binarySearch)
}

func Benchmark_MapSearch____10000_1000(b *testing.B) {
	commonBenchmark(b, nums10000, search1000, MapSearch)
}

func Benchmark_mapSearch____10000_1000(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums10000, search1000, mapSearch)
}

//
// data-10000, search-10000
//

func Benchmark_LinearSearch_10000_10000(b *testing.B) {
	commonBenchmark(b, nums10000, search10000, LinearSearch)
}

func Benchmark_BinarySearch_10000_10000(b *testing.B) {
	commonBenchmark(b, nums10000, search10000, BinarySearch)
}

func Benchmark_binarySearch_10000_10000(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums10000, search10000, binarySearch)
}

func Benchmark_MapSearch____10000_10000(b *testing.B) {
	commonBenchmark(b, nums10000, search10000, MapSearch)
}

func Benchmark_mapSearch____10000_10000(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums10000, search10000, mapSearch)
}

//
// data-50000, search-10000
//

func Benchmark_LinearSearch_50000_10000(b *testing.B) {
	commonBenchmark(b, nums50000, search10000, LinearSearch)
}

func Benchmark_BinarySearch_50000_10000(b *testing.B) {
	commonBenchmark(b, nums50000, search10000, BinarySearch)
}

func Benchmark_binarySearch_50000_10000(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums50000, search10000, binarySearch)
}

func Benchmark_MapSearch____50000_10000(b *testing.B) {
	commonBenchmark(b, nums50000, search10000, MapSearch)
}

func Benchmark_mapSearch____50000_10000(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums50000, search10000, mapSearch)
}
