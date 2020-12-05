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
Benchmark_LinearSearch_1000__3_NotFound-12               1200898              1041 ns/op              32 B/op          1 allocs/op
Benchmark_BinarySearch_1000__3_NotFound-12                 35037             34245 ns/op              96 B/op          3 allocs/op
Benchmark_binarySearch_1000__3_NotFound-12               9712423               115 ns/op              32 B/op          1 allocs/op
Benchmark_MapSearch____1000__3_NotFound-12                 38277             32859 ns/op           41193 B/op          8 allocs/op
Benchmark_mapSearch____1000__3_NotFound-12              17367688              61.9 ns/op              32 B/op          1 allocs/op
Benchmark_LinearSearch_1000__100_NotFound-12               41889             29805 ns/op             896 B/op          1 allocs/op
Benchmark_BinarySearch_1000__100_NotFound-12               30975             37237 ns/op             960 B/op          3 allocs/op
Benchmark_binarySearch_1000__100_NotFound-12              342697              3282 ns/op             896 B/op          1 allocs/op
Benchmark_MapSearch____1000__100_NotFound-12               36074             32351 ns/op           42057 B/op          8 allocs/op
Benchmark_mapSearch____1000__100_NotFound-12              783559              1288 ns/op             896 B/op          1 allocs/op
Benchmark_LinearSearch_1000__1000_NotFound-12               3541            287076 ns/op            8192 B/op          1 allocs/op
Benchmark_BinarySearch_1000__1000_NotFound-12              18427             67671 ns/op            8256 B/op          3 allocs/op
Benchmark_binarySearch_1000__1000_NotFound-12              37875             31188 ns/op            8192 B/op          1 allocs/op
Benchmark_MapSearch____1000__1000_NotFound-12              27123             45042 ns/op           49353 B/op          8 allocs/op
Benchmark_mapSearch____1000__1000_NotFound-12              83547             12659 ns/op            8192 B/op          1 allocs/op
Benchmark_LinearSearch_10000_3_NotFound-12                133090             10778 ns/op              32 B/op          1 allocs/op
Benchmark_BinarySearch_10000_3_NotFound-12                  1992            587961 ns/op              96 B/op          3 allocs/op
Benchmark_binarySearch_10000_3_NotFound-12               6005101               195 ns/op              32 B/op          1 allocs/op
Benchmark_MapSearch____10000_3_NotFound-12                  3416            370013 ns/op          322289 B/op         13 allocs/op
Benchmark_mapSearch____10000_3_NotFound-12              19178061              63.1 ns/op              32 B/op          1 allocs/op
Benchmark_LinearSearch_10000_100_NotFound-12                4141            272541 ns/op             896 B/op          1 allocs/op
Benchmark_BinarySearch_10000_100_NotFound-12                2673            452406 ns/op             960 B/op          3 allocs/op
Benchmark_binarySearch_10000_100_NotFound-12              284740              4563 ns/op             896 B/op          1 allocs/op
Benchmark_MapSearch____10000_100_NotFound-12                2964            346300 ns/op          323174 B/op         13 allocs/op
Benchmark_mapSearch____10000_100_NotFound-12              938948              1285 ns/op             896 B/op          1 allocs/op
Benchmark_LinearSearch_10000_1000_NotFound-12                435           2870078 ns/op            8192 B/op          1 allocs/op
Benchmark_BinarySearch_10000_1000_NotFound-12               2119            486083 ns/op            8256 B/op          3 allocs/op
Benchmark_binarySearch_10000_1000_NotFound-12              26682             40681 ns/op            8192 B/op          1 allocs/op
Benchmark_MapSearch____10000_1000_NotFound-12               3466            358536 ns/op          330494 B/op         13 allocs/op
Benchmark_mapSearch____10000_1000_NotFound-12              90537             12546 ns/op            8192 B/op          1 allocs/op
Benchmark_LinearSearch_10000_10000_NotFound-12                38          27388279 ns/op           81920 B/op          1 allocs/op
Benchmark_BinarySearch_10000_10000_NotFound-12              1333            844088 ns/op           81984 B/op          3 allocs/op
Benchmark_binarySearch_10000_10000_NotFound-12              2542            417791 ns/op           81920 B/op          1 allocs/op
Benchmark_MapSearch____10000_10000_NotFound-12              2362            537530 ns/op          404225 B/op         13 allocs/op
Benchmark_mapSearch____10000_10000_NotFound-12              9194            120524 ns/op           81920 B/op          1 allocs/op
Benchmark_LinearSearch_50000_10000_NotFound-12                 7         171894166 ns/op           81920 B/op          1 allocs/op
Benchmark_BinarySearch_50000_10000_NotFound-12               382           3537735 ns/op           81984 B/op          3 allocs/op
Benchmark_binarySearch_50000_10000_NotFound-12              2385            471677 ns/op           81920 B/op          1 allocs/op
Benchmark_MapSearch____50000_10000_NotFound-12               452           2471348 ns/op         1494987 B/op        847 allocs/op
Benchmark_mapSearch____50000_10000_NotFound-12              9882            178866 ns/op           81920 B/op          1 allocs/op
Benchmark_LinearSearch_1000__3-12                        1355110               887 ns/op              32 B/op          1 allocs/op
Benchmark_BinarySearch_1000__3-12                          33398             33985 ns/op              96 B/op          3 allocs/op
Benchmark_binarySearch_1000__3-12                        9556736               138 ns/op              32 B/op          1 allocs/op
Benchmark_MapSearch____1000__3-12                          33405             31856 ns/op           41193 B/op          8 allocs/op
Benchmark_mapSearch____1000__3-12                       21335660              50.7 ns/op              32 B/op          1 allocs/op
Benchmark_LinearSearch_1000__100-12                       501254              2388 ns/op             896 B/op          1 allocs/op
Benchmark_BinarySearch_1000__100-12                        30468             41578 ns/op             960 B/op          3 allocs/op
Benchmark_binarySearch_1000__100-12                       269874              3892 ns/op             896 B/op          1 allocs/op
Benchmark_MapSearch____1000__100-12                        35611             32581 ns/op           42057 B/op          8 allocs/op
Benchmark_mapSearch____1000__100-12                      1000000              1081 ns/op             896 B/op          1 allocs/op
Benchmark_LinearSearch_1000__1000-12                        8306            152265 ns/op            8192 B/op          1 allocs/op
Benchmark_BinarySearch_1000__1000-12                       10000            105576 ns/op            8256 B/op          3 allocs/op
Benchmark_binarySearch_1000__1000-12                       17820             68080 ns/op            8192 B/op          1 allocs/op
Benchmark_MapSearch____1000__1000-12                       21610             53298 ns/op           49353 B/op          8 allocs/op
Benchmark_mapSearch____1000__1000-12                       58798             18423 ns/op            8192 B/op          1 allocs/op
Benchmark_LinearSearch_10000_3-12                        1377502               874 ns/op              32 B/op          1 allocs/op
Benchmark_BinarySearch_10000_3-12                           2472            447825 ns/op              96 B/op          3 allocs/op
Benchmark_binarySearch_10000_3-12                        7272561               164 ns/op              32 B/op          1 allocs/op
Benchmark_MapSearch____10000_3-12                           3445            337668 ns/op          322346 B/op         13 allocs/op
Benchmark_mapSearch____10000_3-12                       23896603              51.3 ns/op              32 B/op          1 allocs/op
Benchmark_LinearSearch_10000_100-12                       504427              2480 ns/op             896 B/op          1 allocs/op
Benchmark_BinarySearch_10000_100-12                         2534            454719 ns/op             960 B/op          3 allocs/op
Benchmark_binarySearch_10000_100-12                       210144              5019 ns/op             896 B/op          1 allocs/op
Benchmark_MapSearch____10000_100-12                         3038            337526 ns/op          323174 B/op         13 allocs/op
Benchmark_mapSearch____10000_100-12                       971736              1129 ns/op             896 B/op          1 allocs/op
Benchmark_LinearSearch_10000_1000-12                        7874            143806 ns/op            8192 B/op          1 allocs/op
Benchmark_BinarySearch_10000_1000-12                        2115            530711 ns/op            8256 B/op          3 allocs/op
Benchmark_binarySearch_10000_1000-12                       16612             71856 ns/op            8192 B/op          1 allocs/op
Benchmark_MapSearch____10000_1000-12                        2916            381503 ns/op          330458 B/op         13 allocs/op
Benchmark_mapSearch____10000_1000-12                       81838             14044 ns/op            8192 B/op          1 allocs/op
Benchmark_LinearSearch_10000_10000-12                         86          13614430 ns/op           81920 B/op          1 allocs/op
Benchmark_BinarySearch_10000_10000-12                        945           1193840 ns/op           81984 B/op          3 allocs/op
Benchmark_binarySearch_10000_10000-12                       1503            790767 ns/op           81920 B/op          1 allocs/op
Benchmark_MapSearch____10000_10000-12                       1629            672223 ns/op          404231 B/op         13 allocs/op
Benchmark_mapSearch____10000_10000-12                       3331            309937 ns/op           81920 B/op          1 allocs/op
Benchmark_LinearSearch_50000_10000-12                         62          17120934 ns/op           81920 B/op          1 allocs/op
Benchmark_BinarySearch_50000_10000-12                        337           3661868 ns/op           81984 B/op          3 allocs/op
Benchmark_binarySearch_50000_10000-12                       1270           1031030 ns/op           81920 B/op          1 allocs/op
Benchmark_MapSearch____50000_10000-12                        374           2872006 ns/op         1495159 B/op        848 allocs/op
Benchmark_mapSearch____50000_10000-12                       4281            263885 ns/op           81920 B/op          1 allocs/op
PASS
ok      github.com/fujiwaram/slice-vs-map       122.127s
```