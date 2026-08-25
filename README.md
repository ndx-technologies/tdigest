Go t-digest

- 100% test coverage
- fuzz tests
- 200 LOC

[![codecov](https://codecov.io/github/ndx-technologies/tdigest/graph/badge.svg?token=P7YNV35JYV)](https://codecov.io/github/ndx-technologies/tdigest)

Benchmarks

```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/tdigest
cpu: Apple M3 Max
BenchmarkNoop______________________________________-16          1000000000               0.0000001 ns/op               0 B/op          0 allocs/op
BenchmarkAddCompress/size=10-16                                 11996425                99.30 ns/op          296 B/op          4 allocs/op
BenchmarkAddCompress/size=100-16                                12134914               153.4 ns/op           617 B/op          0 allocs/op
BenchmarkAddCompress/size=1000-16                                5096000               760.6 ns/op          3066 B/op          0 allocs/op
BenchmarkAddCompress/size=10000-16                               1414242               848.5 ns/op           278 B/op          0 allocs/op
BenchmarkMergeCompress/size=10-16                                3262725               379.0 ns/op           672 B/op          7 allocs/op
BenchmarkMergeCompress/size=100-16                                283158              4446 ns/op            2800 B/op          6 allocs/op
BenchmarkMergeCompress/size=1000-16                                19807             64773 ns/op           24887 B/op          6 allocs/op
BenchmarkMergeCompress/size=10000-16                                1263           1068702 ns/op          245755 B/op          6 allocs/op
BenchmarkQuantile/size=10/quantile=0.25-16                      294905778                4.087 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.50-16                      238868716                5.038 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.75-16                      286459106                4.133 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.99-16                      440591511                2.718 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.25-16                     68564571                17.71 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.50-16                     28843784                41.85 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.75-16                     100000000               10.95 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.99-16                     442501465                2.733 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.25-16                     3440470               348.0 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.50-16                     1739107               689.8 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.75-16                     8104960               148.2 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.99-16                    198870909                6.010 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.25-16                     356914              3372 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.50-16                     179014              6838 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.75-16                     851617              1394 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.99-16                   27682272                44.10 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/ndx-technologies/tdigest     33.184s
```

References

- https://github.com/tdunning/t-digest
- https://github.com/facebook/folly/blob/main/folly/stats/TDigest.cpp (no-unprocessed algorithm, inlined compression in creation, no add)
- https://github.com/MnO2/t-digest (same as folly)
- https://github.com/derrickburns/tdigest (processed/unprocessed algorithm, add via new centroid)
- https://github.com/influxdata/tdigest (last update 4y ago, open issues, uses errors in API, processed/unprocessed algorithm, add via new centroid)
- https://github.com/spenczar/tdigest (archived)
