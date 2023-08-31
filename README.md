# arena-test
Simple program with golang arenas (experimental in go 1.20)

```
$ export GOEXPERIMENT=arenas
```

```
$ go test -bench=. -benchtime=10x -benchmem arena_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkNoGC-8   	      10	       146.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGC-8     	      10	     16735 ns/op	    2035 B/op	       2 allocs/op
PASS
ok  	command-line-arguments	0.005s
```

```
$ go test -bench=. -benchtime=100x -benchmem arena_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkNoGC-8   	     100	        97.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGC-8     	     100	      1295 ns/op	    1840 B/op	       2 allocs/op
PASS
ok  	command-line-arguments	0.004s
```

```
$ go test -bench=. -benchtime=1000x -benchmem arena_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkNoGC-8   	    1000	        44.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGC-8     	    1000	       415.4 ns/op	    1840 B/op	       2 allocs/op
PASS
ok  	command-line-arguments	0.005s
```

```
$ go test -bench=. -benchtime=10000x -benchmem arena_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkNoGC-8   	   10000	       186.6 ns/op	     838 B/op	       0 allocs/op
BenchmarkGC-8     	   10000	       316.0 ns/op	    1840 B/op	       2 allocs/op
PASS
ok  	command-line-arguments	0.010s
```

```
$ go test -bench=. -benchtime=100000x -benchmem arena_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkNoGC-8   	  100000	       284.8 ns/op	    1593 B/op	       0 allocs/op
BenchmarkGC-8     	  100000	       285.6 ns/op	    1840 B/op	       2 allocs/op
PASS
ok  	command-line-arguments	0.062s
```

```
$ go test -bench=. -benchtime=1000000x -benchmem arena_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkNoGC-8   	 1000000	       266.7 ns/op	    1644 B/op	       0 allocs/op
BenchmarkGC-8     	 1000000	       227.1 ns/op	    1840 B/op	       2 allocs/op
PASS
ok  	command-line-arguments	0.501s
```

```
$ go test -bench=. -benchtime=10000000x -benchmem arena_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkNoGC-8   	10000000	       649.1 ns/op	    1647 B/op	       0 allocs/op
BenchmarkGC-8     	10000000	       210.8 ns/op	    1840 B/op	       2 allocs/op
PASS
ok  	command-line-arguments	8.668s
```

```
$ go test -bench=. -benchtime=100000000x -benchmem arena_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkNoGC-8   	signal: killed
FAIL	command-line-arguments	46.783s
FAIL
```
