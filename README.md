Go t-digest

- 100% test coverage
- fuzz tests
- binary encoding
- 300 LOC

Benchmarks

```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/tdigest
cpu: Apple M3 Max
BenchmarkNoop______________________________________-16          1000000000               0.0000000 ns/op               0 B/op          0 allocs/op
BenchmarkAddCompress/size=10-16                                 10171742               123.4 ns/op           295 B/op          3 allocs/op
BenchmarkAddCompress/size=100-16                                12780916                96.39 ns/op          393 B/op          0 allocs/op
BenchmarkAddCompress/size=1000-16                                3908049               879.4 ns/op          3619 B/op          0 allocs/op
BenchmarkAddCompress/size=10000-16                               1830889               676.9 ns/op           143 B/op          0 allocs/op
BenchmarkMergeCompress/size=10-16                                3081964               388.5 ns/op           666 B/op          6 allocs/op
BenchmarkMergeCompress/size=100-16                                263866              4248 ns/op            2799 B/op          6 allocs/op
BenchmarkMergeCompress/size=1000-16                                20254             59185 ns/op           24940 B/op          6 allocs/op
BenchmarkMergeCompress/size=10000-16                                1270            885016 ns/op          245749 B/op          6 allocs/op
BenchmarkQuantile/size=10/quantile=0.25-16                      298937866                4.056 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.50-16                      247593187                4.884 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.75-16                      297937402                4.022 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.99-16                      436768309                2.839 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.25-16                     56765433                18.42 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.50-16                     28905532                41.72 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.75-16                     100000000               11.41 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.99-16                     448402147                2.673 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.25-16                     3713737               322.0 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.50-16                     1876771               646.4 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.75-16                     8074629               148.4 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.99-16                    201317398                5.989 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.25-16                     373893              3134 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.50-16                     193239              6222 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.75-16                     868594              1360 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.99-16                   27529591                43.52 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=10-16                        100000000               10.44 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=100-16                       17123602                70.02 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=1000-16                       1808103               658.9 ns/op             0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=10000-16                       182833              6541 ns/op               0 B/op          0 allocs/op
BenchmarkTDigest_MarshalBinary/size=10-16                       45071620                23.94 ns/op          112 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=100-16                       8930018               133.7 ns/op           896 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=1000-16                      1000000              1150 ns/op            8192 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=10000-16                      118453             10106 ns/op           81920 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=10-16                     49401945                24.60 ns/op           80 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=100-16                     7441465               168.4 ns/op           896 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=1000-16                     818502              1555 ns/op            8192 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=10000-16                     90664             13060 ns/op           81920 B/op          1 allocs/op
PASS
ok      github.com/ndx-technologies/tdigest     45.601s
```

References

- https://github.com/tdunning/t-digest
- https://github.com/facebook/folly/blob/main/folly/stats/TDigest.cpp (no-unprocessed algorithm, inlined compression in creation, no add)
- https://github.com/MnO2/t-digest (same as folly)
- https://github.com/derrickburns/tdigest (processed/unprocessed algorithm, add via new centroid)
- https://github.com/influxdata/tdigest (last update 4y ago, open issues, uses errors in API, processed/unprocessed algorithm, add via new centroid)
- https://github.com/spenczar/tdigest (archived)
