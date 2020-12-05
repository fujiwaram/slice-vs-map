# slice-vs-map
Performance investigation slice vs map for golang



# Usage

```
go test -bench . -benchmem
```

# Results

## go 1.14.7 - mac OS Mojave - CPU 2.2 GHz Intel Core i7 - Memory 16GB

```
go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/fujiwaram/slice-vs-map
Benchmark_LinearSearch_1000__3_NotFound-12               9323829               119 ns/op              32 B/op          1 allocs/op
Benchmark_BinarySearch_1000__3_NotFound-12                 36187             33794 ns/op              96 B/op          3 allocs/op
Benchmark_binarySearch_1000__3_NotFound-12               9893866               116 ns/op              32 B/op          1 allocs/op
Benchmark_MapSearch____1000__3_NotFound-12                 39306             31149 ns/op           41193 B/op          8 allocs/op
Benchmark_mapSearch____1000__3_NotFound-12              20573210              58.0 ns/op            32 B/op          1 allocs/op
Benchmark_LinearSearch_1000__100_NotFound-12              346309              3263 ns/op             896 B/op          1 allocs/op
Benchmark_BinarySearch_1000__100_NotFound-12               32680             37185 ns/op             960 B/op          3 allocs/op
Benchmark_binarySearch_1000__100_NotFound-12              381375              3130 ns/op             896 B/op          1 allocs/op
Benchmark_MapSearch____1000__100_NotFound-12               38694             31305 ns/op           42057 B/op          8 allocs/op
Benchmark_mapSearch____1000__100_NotFound-12              958657              1239 ns/op             896 B/op          1 allocs/op
Benchmark_LinearSearch_1000__1000_NotFound-12              37316             32809 ns/op            8192 B/op          1 allocs/op
Benchmark_BinarySearch_1000__1000_NotFound-12              19250             63447 ns/op            8256 B/op          3 allocs/op
Benchmark_binarySearch_1000__1000_NotFound-12              38010             30899 ns/op            8192 B/op          1 allocs/op
Benchmark_MapSearch____1000__1000_NotFound-12              28604             43717 ns/op           49353 B/op          8 allocs/op
Benchmark_mapSearch____1000__1000_NotFound-12              96390             11900 ns/op            8192 B/op          1 allocs/op
Benchmark_LinearSearch_10000_3_NotFound-12               7593306               160 ns/op              32 B/op          1 allocs/op
Benchmark_BinarySearch_10000_3_NotFound-12                  2755            441585 ns/op              96 B/op          3 allocs/op
Benchmark_binarySearch_10000_3_NotFound-12               7576921               151 ns/op              32 B/op          1 allocs/op
Benchmark_MapSearch____10000_3_NotFound-12                  3676            326457 ns/op          322319 B/op         13 allocs/op
Benchmark_mapSearch____10000_3_NotFound-12              20296476              58.8 ns/op            32 B/op          1 allocs/op
Benchmark_LinearSearch_10000_100_NotFound-12              271489              4326 ns/op             896 B/op          1 allocs/op
Benchmark_BinarySearch_10000_100_NotFound-12                2736            452755 ns/op             960 B/op          3 allocs/op
Benchmark_binarySearch_10000_100_NotFound-12              301408              4147 ns/op             896 B/op          1 allocs/op
Benchmark_MapSearch____10000_100_NotFound-12                3348            334387 ns/op          323190 B/op         13 allocs/op
Benchmark_mapSearch____10000_100_NotFound-12              924733              1255 ns/op             896 B/op          1 allocs/op
Benchmark_LinearSearch_10000_1000_NotFound-12              28459             44140 ns/op            8192 B/op          1 allocs/op
Benchmark_BinarySearch_10000_1000_NotFound-12               2502            481863 ns/op            8256 B/op          3 allocs/op
Benchmark_binarySearch_10000_1000_NotFound-12              29974             40520 ns/op            8192 B/op          1 allocs/op
Benchmark_MapSearch____10000_1000_NotFound-12               3456            328581 ns/op          330483 B/op         13 allocs/op
Benchmark_mapSearch____10000_1000_NotFound-12             100328             12673 ns/op            8192 B/op          1 allocs/op
Benchmark_LinearSearch_10000_10000_NotFound-12              2812            424294 ns/op           81920 B/op          1 allocs/op
Benchmark_BinarySearch_10000_10000_NotFound-12              1308            832696 ns/op           81984 B/op          3 allocs/op
Benchmark_binarySearch_10000_10000_NotFound-12              3100            397286 ns/op           81920 B/op          1 allocs/op
Benchmark_MapSearch____10000_10000_NotFound-12              2629            455588 ns/op          404192 B/op         13 allocs/op
Benchmark_mapSearch____10000_10000_NotFound-12              8488            121866 ns/op           81920 B/op          1 allocs/op
Benchmark_LinearSearch_50000_10000_NotFound-12              2432            478623 ns/op           81920 B/op          1 allocs/op
Benchmark_BinarySearch_50000_10000_NotFound-12               397           2959751 ns/op           81984 B/op          3 allocs/op
Benchmark_binarySearch_50000_10000_NotFound-12              2595            456126 ns/op           81920 B/op          1 allocs/op
Benchmark_MapSearch____50000_10000_NotFound-12               517           2185981 ns/op         1495315 B/op        849 allocs/op
Benchmark_mapSearch____50000_10000_NotFound-12              9222            115517 ns/op           81920 B/op          1 allocs/op
Benchmark_LinearSearch_1000__3-12                       10121917               120 ns/op              32 B/op          1 allocs/op
Benchmark_BinarySearch_1000__3-12                          35614             34405 ns/op              96 B/op          3 allocs/op
Benchmark_binarySearch_1000__3-12                        9787075               122 ns/op              32 B/op          1 allocs/op
Benchmark_MapSearch____1000__3-12                          39430             30123 ns/op           41193 B/op          8 allocs/op
Benchmark_mapSearch____1000__3-12                       23496907              49.5 ns/op            32 B/op          1 allocs/op
Benchmark_LinearSearch_1000__100-12                       306747              3912 ns/op             896 B/op          1 allocs/op
Benchmark_BinarySearch_1000__100-12                        29623             39694 ns/op             960 B/op          3 allocs/op
Benchmark_binarySearch_1000__100-12                       319347              3718 ns/op             896 B/op          1 allocs/op
Benchmark_MapSearch____1000__100-12                        38542             31520 ns/op           42057 B/op          8 allocs/op
Benchmark_mapSearch____1000__100-12                      1000000              1057 ns/op             896 B/op          1 allocs/op
Benchmark_LinearSearch_1000__1000-12                       16897             72637 ns/op            8192 B/op          1 allocs/op
Benchmark_BinarySearch_1000__1000-12                       10000            101842 ns/op            8256 B/op          3 allocs/op
Benchmark_binarySearch_1000__1000-12                       18055             66846 ns/op            8192 B/op          1 allocs/op
Benchmark_MapSearch____1000__1000-12                       22741             53092 ns/op           49353 B/op          8 allocs/op
Benchmark_mapSearch____1000__1000-12                       70160             16851 ns/op            8192 B/op          1 allocs/op
Benchmark_LinearSearch_10000_3-12                        7210419               162 ns/op              32 B/op          1 allocs/op
Benchmark_BinarySearch_10000_3-12                           2544            443615 ns/op              96 B/op          3 allocs/op
Benchmark_binarySearch_10000_3-12                        7960460               153 ns/op              32 B/op          1 allocs/op
Benchmark_MapSearch____10000_3-12                           3214            320853 ns/op          322324 B/op         13 allocs/op
Benchmark_mapSearch____10000_3-12                       24674314              49.6 ns/op            32 B/op          1 allocs/op
Benchmark_LinearSearch_10000_100-12                       206180              5199 ns/op             896 B/op          1 allocs/op
Benchmark_BinarySearch_10000_100-12                         2281            447746 ns/op             960 B/op          3 allocs/op
Benchmark_binarySearch_10000_100-12                       237156              5058 ns/op             896 B/op          1 allocs/op
Benchmark_MapSearch____10000_100-12                         3590            323855 ns/op          323206 B/op         13 allocs/op
Benchmark_mapSearch____10000_100-12                      1000000              1034 ns/op             896 B/op          1 allocs/op
Benchmark_LinearSearch_10000_1000-12                       15594             79584 ns/op            8192 B/op          1 allocs/op
Benchmark_BinarySearch_10000_1000-12                        2222            538851 ns/op            8256 B/op          3 allocs/op
Benchmark_binarySearch_10000_1000-12                       16951             72578 ns/op            8192 B/op          1 allocs/op
Benchmark_MapSearch____10000_1000-12                        3278            339575 ns/op          330472 B/op         13 allocs/op
Benchmark_mapSearch____10000_1000-12                       88249             13089 ns/op            8192 B/op          1 allocs/op
Benchmark_LinearSearch_10000_10000-12                       1574            788938 ns/op           81920 B/op          1 allocs/op
Benchmark_BinarySearch_10000_10000-12                        933           1183497 ns/op           81984 B/op          3 allocs/op
Benchmark_binarySearch_10000_10000-12                       1651            735125 ns/op           81920 B/op          1 allocs/op
Benchmark_MapSearch____10000_10000-12                       1926            563017 ns/op          404171 B/op         13 allocs/op
Benchmark_mapSearch____10000_10000-12                       4801            248436 ns/op           81920 B/op          1 allocs/op
Benchmark_LinearSearch_50000_10000-12                       1359            820530 ns/op           81920 B/op          1 allocs/op
Benchmark_BinarySearch_50000_10000-12                        349           3352729 ns/op           81984 B/op          3 allocs/op
Benchmark_binarySearch_50000_10000-12                       1556            758114 ns/op           81920 B/op          1 allocs/op
Benchmark_MapSearch____50000_10000-12                        519           2250849 ns/op         1495393 B/op        850 allocs/op
Benchmark_mapSearch____50000_10000-12                       5649            208536 ns/op           81920 B/op          1 allocs/op
PASS
ok      github.com/fujiwaram/slice-vs-map       110.890s
```