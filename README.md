# slice-vs-map
Performance investigation slice vs map for golang



# Usage

```
go test -bench . -benchmem
```

# Results

## go 1.14.7 - mac OS Mojave - CPU 2.2 GHz Intel Core i7 - Memory 16GB

```
Benchmark_LinearSearch_1000__3_NotFound-12        	 8324360	       127 ns/op	      32 B/op	       1 allocs/op
Benchmark_BinarySearch_1000__3_NotFound-12        	   38431	     30327 ns/op	      96 B/op	       3 allocs/op
Benchmark_binarySearch_1000__3_NotFound-12        	10638997	       124 ns/op	      32 B/op	       1 allocs/op
Benchmark_MapSearch____1000__3_NotFound-12        	   36813	     31075 ns/op	   41194 B/op	       8 allocs/op
Benchmark_mapSearch____1000__3_NotFound-12        	16177371	        70.3 ns/op	      32 B/op	       1 allocs/op
Benchmark_LinearSearch_1000__100_NotFound-12      	  338710	      4020 ns/op	     896 B/op	       1 allocs/op
Benchmark_BinarySearch_1000__100_NotFound-12      	   36348	     32902 ns/op	     960 B/op	       3 allocs/op
Benchmark_binarySearch_1000__100_NotFound-12      	  385741	      3246 ns/op	     896 B/op	       1 allocs/op
Benchmark_MapSearch____1000__100_NotFound-12      	   32944	     42803 ns/op	   42057 B/op	       8 allocs/op
Benchmark_mapSearch____1000__100_NotFound-12      	  934680	      1386 ns/op	     896 B/op	       1 allocs/op
Benchmark_LinearSearch_1000__1000_NotFound-12     	   34416	     41181 ns/op	    8192 B/op	       1 allocs/op
Benchmark_BinarySearch_1000__1000_NotFound-12     	   17850	     61995 ns/op	    8256 B/op	       3 allocs/op
Benchmark_binarySearch_1000__1000_NotFound-12     	   35134	     32229 ns/op	    8192 B/op	       1 allocs/op
Benchmark_MapSearch____1000__1000_NotFound-12     	   26571	     52370 ns/op	   49353 B/op	       8 allocs/op
Benchmark_mapSearch____1000__1000_NotFound-12     	   95305	     13841 ns/op	    8192 B/op	       1 allocs/op
Benchmark_LinearSearch_10000_3_NotFound-12        	 7914811	       154 ns/op	      32 B/op	       1 allocs/op
Benchmark_BinarySearch_10000_3_NotFound-12        	    2979	    503525 ns/op	      96 B/op	       3 allocs/op
Benchmark_binarySearch_10000_3_NotFound-12        	 5179592	       217 ns/op	      32 B/op	       1 allocs/op
Benchmark_MapSearch____10000_3_NotFound-12        	    2632	    463745 ns/op	  322328 B/op	      13 allocs/op
Benchmark_mapSearch____10000_3_NotFound-12        	15496969	        72.7 ns/op	      32 B/op	       1 allocs/op
Benchmark_LinearSearch_10000_100_NotFound-12      	  289971	      4465 ns/op	     896 B/op	       1 allocs/op
Benchmark_BinarySearch_10000_100_NotFound-12      	    2836	    474940 ns/op	     960 B/op	       3 allocs/op
Benchmark_binarySearch_10000_100_NotFound-12      	  214908	      5664 ns/op	     896 B/op	       1 allocs/op
Benchmark_MapSearch____10000_100_NotFound-12      	    2847	    478482 ns/op	  323193 B/op	      13 allocs/op
Benchmark_mapSearch____10000_100_NotFound-12      	  626500	      1966 ns/op	     896 B/op	       1 allocs/op
Benchmark_LinearSearch_10000_1000_NotFound-12     	   23624	     51129 ns/op	    8192 B/op	       1 allocs/op
Benchmark_BinarySearch_10000_1000_NotFound-12     	    2246	    586345 ns/op	    8256 B/op	       3 allocs/op
Benchmark_binarySearch_10000_1000_NotFound-12     	   21776	     54574 ns/op	    8192 B/op	       1 allocs/op
Benchmark_MapSearch____10000_1000_NotFound-12     	    2296	    440016 ns/op	  330494 B/op	      13 allocs/op
Benchmark_mapSearch____10000_1000_NotFound-12     	   71128	     15588 ns/op	    8192 B/op	       1 allocs/op
Benchmark_LinearSearch_10000_10000_NotFound-12    	    2389	    546483 ns/op	   81920 B/op	       1 allocs/op
Benchmark_BinarySearch_10000_10000_NotFound-12    	    1185	   1032770 ns/op	   81984 B/op	       3 allocs/op
Benchmark_binarySearch_10000_10000_NotFound-12    	    2372	    517075 ns/op	   81920 B/op	       1 allocs/op
Benchmark_MapSearch____10000_10000_NotFound-12    	    2000	    610571 ns/op	  404185 B/op	      13 allocs/op
Benchmark_mapSearch____10000_10000_NotFound-12    	    7140	    174533 ns/op	   81920 B/op	       1 allocs/op
Benchmark_LinearSearch_50000_10000_NotFound-12    	    1897	    581003 ns/op	   81920 B/op	       1 allocs/op
Benchmark_BinarySearch_50000_10000_NotFound-12    	     319	   3519102 ns/op	   81984 B/op	       3 allocs/op
Benchmark_binarySearch_50000_10000_NotFound-12    	    2544	    668809 ns/op	   81920 B/op	       1 allocs/op
Benchmark_MapSearch____50000_10000_NotFound-12    	     358	   3033506 ns/op	 1495107 B/op	     848 allocs/op
Benchmark_mapSearch____50000_10000_NotFound-12    	   10000	    118264 ns/op	   81920 B/op	       1 allocs/op
Benchmark_LinearSearch_1000__3-12                 	 8851676	       136 ns/op	      32 B/op	       1 allocs/op
Benchmark_BinarySearch_1000__3-12                 	   38305	     32138 ns/op	      96 B/op	       3 allocs/op
Benchmark_binarySearch_1000__3-12                 	 8444023	       145 ns/op	      32 B/op	       1 allocs/op
Benchmark_MapSearch____1000__3-12                 	   29637	     36981 ns/op	   41193 B/op	       8 allocs/op
Benchmark_mapSearch____1000__3-12                 	21617029	      60.0 ns/op	      32 B/op	       1 allocs/op
Benchmark_LinearSearch_1000__100-12               	  288295	      4852 ns/op	     896 B/op	       1 allocs/op
Benchmark_BinarySearch_1000__100-12               	   31184	     39326 ns/op	     960 B/op	       3 allocs/op
Benchmark_binarySearch_1000__100-12               	  307419	      4132 ns/op	     896 B/op	       1 allocs/op
Benchmark_MapSearch____1000__100-12               	   34598	     33346 ns/op	   42057 B/op	       8 allocs/op
Benchmark_mapSearch____1000__100-12               	 1000000	      1127 ns/op	     896 B/op	       1 allocs/op
Benchmark_LinearSearch_1000__1000-12              	   17692	     66693 ns/op	    8192 B/op	       1 allocs/op
Benchmark_BinarySearch_1000__1000-12              	   10000	    105409 ns/op	    8256 B/op	       3 allocs/op
Benchmark_binarySearch_1000__1000-12              	   17301	     69291 ns/op	    8192 B/op	       1 allocs/op
Benchmark_MapSearch____1000__1000-12              	   21204	     53141 ns/op	   49352 B/op	       8 allocs/op
Benchmark_mapSearch____1000__1000-12              	   69103	     18094 ns/op	    8192 B/op	       1 allocs/op
Benchmark_LinearSearch_10000_3-12                 	 7516464	       159 ns/op	      32 B/op	       1 allocs/op
Benchmark_BinarySearch_10000_3-12                 	    3108	    389350 ns/op	      96 B/op	       3 allocs/op
Benchmark_binarySearch_10000_3-12                 	 6973698	       156 ns/op	      32 B/op	       1 allocs/op
Benchmark_MapSearch____10000_3-12                 	    3772	    334742 ns/op	  322315 B/op	      13 allocs/op
Benchmark_mapSearch____10000_3-12                 	24772274	      49.7 ns/op	      32 B/op	       1 allocs/op
Benchmark_LinearSearch_10000_100-12               	  239727	      5009 ns/op	     896 B/op	       1 allocs/op
Benchmark_BinarySearch_10000_100-12               	    3048	    404893 ns/op	     960 B/op	       3 allocs/op
Benchmark_binarySearch_10000_100-12               	  226522	      5168 ns/op	     896 B/op	       1 allocs/op
Benchmark_MapSearch____10000_100-12               	    3184	    331067 ns/op	  323196 B/op	      13 allocs/op
Benchmark_mapSearch____10000_100-12               	 1000000	      1094 ns/op	     896 B/op	       1 allocs/op
Benchmark_LinearSearch_10000_1000-12              	   16344	     75934 ns/op	    8192 B/op	       1 allocs/op
Benchmark_BinarySearch_10000_1000-12              	    2149	    513521 ns/op	    8256 B/op	       3 allocs/op
Benchmark_binarySearch_10000_1000-12              	   16557	     71083 ns/op	    8192 B/op	       1 allocs/op
Benchmark_MapSearch____10000_1000-12              	    3584	    336511 ns/op	  330467 B/op	      13 allocs/op
Benchmark_mapSearch____10000_1000-12              	   74372	     15622 ns/op	    8192 B/op	       1 allocs/op
Benchmark_LinearSearch_10000_10000-12             	    1546	    759260 ns/op	   81920 B/op	       1 allocs/op
Benchmark_BinarySearch_10000_10000-12             	    1052	   1142117 ns/op	   81984 B/op	       3 allocs/op
Benchmark_binarySearch_10000_10000-12             	    1608	    747372 ns/op	   81920 B/op	       1 allocs/op
Benchmark_MapSearch____10000_10000-12             	    2046	    623346 ns/op	  404204 B/op	      13 allocs/op
Benchmark_mapSearch____10000_10000-12             	    4190	    258955 ns/op	   81920 B/op	       1 allocs/op
Benchmark_LinearSearch_50000_10000-12             	    1544	    779130 ns/op	   81920 B/op	       1 allocs/op
Benchmark_BinarySearch_50000_10000-12             	     396	   3037436 ns/op	   81984 B/op	       3 allocs/op
Benchmark_binarySearch_50000_10000-12             	    1424	    826831 ns/op	   81920 B/op	       1 allocs/op
Benchmark_MapSearch____50000_10000-12             	     552	   2160789 ns/op	 1494904 B/op	     846 allocs/op
Benchmark_mapSearch____50000_10000-12             	    5691	    210062 ns/op	   81920 B/op	       1 allocs/op
```