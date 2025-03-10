t-digest in Go

- 100% test coverage
- fuzz tests
- binary encoding
- 300 LOC

## Benchmarks

```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/tdigest
cpu: Apple M3 Max
BenchmarkAddCompress/size=10-16                                  6347971               185.4 ns/op           352 B/op          5 allocs/op
BenchmarkAddCompress/size=100-16                                 2437504               497.8 ns/op           413 B/op          2 allocs/op
BenchmarkAddCompress/size=1000-16                                 502335              2796 ns/op             822 B/op          2 allocs/op
BenchmarkAddCompress/size=10000-16                                116229             18801 ns/op             612 B/op          2 allocs/op
BenchmarkMergeCompress/size=10-16                                  10000            799252 ns/op             601 B/op          5 allocs/op
BenchmarkMergeCompress/size=100-16                                  5858          10873835 ns/op            4605 B/op          5 allocs/op
BenchmarkMergeCompress/size=1000-16                                  474          13394475 ns/op           46533 B/op          5 allocs/op
BenchmarkMergeCompress/size=10000-16                                 100          31891518 ns/op          451982 B/op          5 allocs/op
BenchmarkQuantile/size=10/quantile=0.25-16                      299203184                4.061 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.50-16                      241388850                4.877 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.75-16                      298027831                4.056 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.99-16                      385692108                3.152 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.25-16                     67147268                17.93 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.50-16                     28660680                42.35 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.75-16                     100000000               11.84 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.99-16                     371180035                3.236 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.25-16                     3404103               351.5 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.50-16                     1751793               684.8 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.75-16                     7998908               153.5 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.99-16                    193594092                6.231 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.25-16                     355957              3279 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.50-16                     172063              6680 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.75-16                     881113              1371 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.99-16                   26819563                45.67 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=10-16                        90984910                13.11 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=100-16                       14939246                79.54 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=1000-16                       1691401               709.0 ns/op             0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=10000-16                       171600              7026 ns/op               0 B/op          0 allocs/op
BenchmarkTDigest_MarshalBinary/size=10-16                       41681983                28.39 ns/op          112 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=100-16                       8543856               140.7 ns/op           768 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=1000-16                       927478              1277 ns/op            8192 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=10000-16                      112681             10537 ns/op           73728 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=10-16                     47608263                24.79 ns/op           80 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=100-16                     7531831               151.4 ns/op           768 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=1000-16                     881078              1418 ns/op            8192 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=10000-16                     99382             11969 ns/op           73728 B/op          1 allocs/op
PASS
ok      github.com/ndx-technologies/tdigest     128.191s
```

## References

- https://github.com/tdunning/t-digest
- https://github.com/facebook/folly/blob/main/folly/stats/TDigest.cpp (no-unprocessed algorithm, inlined compression in creation, no add)
- https://github.com/MnO2/t-digest (same as folly)
- https://github.com/derrickburns/tdigest (processed/unprocessed algorithm, add via new centroid)
- https://github.com/influxdata/tdigest (not maintained, last update 4y ago, open issues, uses errors in API, processed/unprocessed algorithm, add via new centroid)
- https://github.com/spenczar/tdigest (archived)
