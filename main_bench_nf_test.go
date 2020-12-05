package main

import (
	"testing"
)

//
// data-1000, search-3
//

func Benchmark_LinearSearch_1000__3_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search3NotFound, LinearSearch)
}

func Benchmark_BinarySearch_1000__3_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search3NotFound, BinarySearch)
}

func Benchmark_binarySearch_1000__3_NotFound(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums1000, search3NotFound, binarySearch)
}

func Benchmark_MapSearch____1000__3_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search3NotFound, MapSearch)
}

func Benchmark_mapSearch____1000__3_NotFound(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums1000, search3NotFound, mapSearch)
}

//
// data-1000, search-100
//

func Benchmark_LinearSearch_1000__100_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search100NotFound, LinearSearch)
}

func Benchmark_BinarySearch_1000__100_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search100NotFound, BinarySearch)
}

func Benchmark_binarySearch_1000__100_NotFound(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums1000, search100NotFound, binarySearch)
}

func Benchmark_MapSearch____1000__100_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search100NotFound, MapSearch)
}

func Benchmark_mapSearch____1000__100_NotFound(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums1000, search100NotFound, mapSearch)
}

//
// data-1000, search-1000
//

func Benchmark_LinearSearch_1000__1000_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search1000NotFound, LinearSearch)
}

func Benchmark_BinarySearch_1000__1000_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search1000NotFound, BinarySearch)
}

func Benchmark_binarySearch_1000__1000_NotFound(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums1000, search1000NotFound, binarySearch)
}

func Benchmark_MapSearch____1000__1000_NotFound(b *testing.B) {
	commonBenchmark(b, nums1000, search1000NotFound, MapSearch)
}

func Benchmark_mapSearch____1000__1000_NotFound(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums1000, search1000NotFound, mapSearch)
}

//
// data-10000, search-3
//

func Benchmark_LinearSearch_10000_3_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search3NotFound, LinearSearch)
}

func Benchmark_BinarySearch_10000_3_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search3NotFound, BinarySearch)
}

func Benchmark_binarySearch_10000_3_NotFound(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums10000, search3NotFound, binarySearch)
}

func Benchmark_MapSearch____10000_3_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search3NotFound, MapSearch)
}

func Benchmark_mapSearch____10000_3_NotFound(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums10000, search3NotFound, mapSearch)
}

//
// data-10000, search-100
//

func Benchmark_LinearSearch_10000_100_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search100NotFound, LinearSearch)
}

func Benchmark_BinarySearch_10000_100_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search100NotFound, BinarySearch)
}

func Benchmark_binarySearch_10000_100_NotFound(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums10000, search100NotFound, binarySearch)
}

func Benchmark_MapSearch____10000_100_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search100NotFound, MapSearch)
}

func Benchmark_mapSearch____10000_100_NotFound(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums10000, search100NotFound, mapSearch)
}

//
// data-10000, search-1000
//

func Benchmark_LinearSearch_10000_1000_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search1000NotFound, LinearSearch)
}

func Benchmark_BinarySearch_10000_1000_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search1000NotFound, BinarySearch)
}

func Benchmark_binarySearch_10000_1000_NotFound(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums10000, search1000NotFound, binarySearch)
}

func Benchmark_MapSearch____10000_1000_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search1000NotFound, MapSearch)
}

func Benchmark_mapSearch____10000_1000_NotFound(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums10000, search1000NotFound, mapSearch)
}

//
// data-10000, search-10000
//

func Benchmark_LinearSearch_10000_10000_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search10000NotFound, LinearSearch)
}

func Benchmark_BinarySearch_10000_10000_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search10000NotFound, BinarySearch)
}

func Benchmark_binarySearch_10000_10000_NotFound(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums10000, search10000NotFound, binarySearch)
}

func Benchmark_MapSearch____10000_10000_NotFound(b *testing.B) {
	commonBenchmark(b, nums10000, search10000NotFound, MapSearch)
}

func Benchmark_mapSearch____10000_10000_NotFound(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums10000, search10000NotFound, mapSearch)
}

//
// data-50000, search-10000
//

func Benchmark_LinearSearch_50000_10000_NotFound(b *testing.B) {
	commonBenchmark(b, nums50000, search10000NotFound, LinearSearch)
}

func Benchmark_BinarySearch_50000_10000_NotFound(b *testing.B) {
	commonBenchmark(b, nums50000, search10000NotFound, BinarySearch)
}

func Benchmark_binarySearch_50000_10000_NotFound(b *testing.B) {
	commonBenchmarkFroBinarySearch(b, nums50000, search10000NotFound, binarySearch)
}

func Benchmark_MapSearch____50000_10000_NotFound(b *testing.B) {
	commonBenchmark(b, nums50000, search10000NotFound, MapSearch)
}

func Benchmark_mapSearch____50000_10000_NotFound(b *testing.B) {
	commonBenchmarkFroMapSearch(b, nums50000, search10000NotFound, mapSearch)
}
