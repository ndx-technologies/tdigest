<p align="center">
  ✨ Sponsored by <a href="https://apps.apple.com/app/id6738306589">NDX AI Shopping Assistant</a>
</p>

Go t-digest

- 100% test coverage
- fuzz tests
- 200 LOC

Benchmarks

```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/tdigest
cpu: Apple M3 Max
BenchmarkNoop______________________________________-16          1000000000             0 B/op          0 allocs/op
BenchmarkAddCompress/size=10-16                                  9717769               117.4 ns/op           295 B/op          3 allocs/op
BenchmarkAddCompress/size=100-16                                12581143               137.1 ns/op           611 B/op          0 allocs/op
BenchmarkAddCompress/size=1000-16                                5193721               354.0 ns/op          1472 B/op          0 allocs/op
BenchmarkAddCompress/size=10000-16                               1257470               961.4 ns/op           229 B/op          0 allocs/op
BenchmarkMergeCompress/size=10-16                                3237510               383.7 ns/op           666 B/op          6 allocs/op
BenchmarkMergeCompress/size=100-16                                286646              4325 ns/op            2800 B/op          6 allocs/op
BenchmarkMergeCompress/size=1000-16                                21424             55897 ns/op           29672 B/op          6 allocs/op
BenchmarkMergeCompress/size=10000-16                                1360            868039 ns/op          245758 B/op          6 allocs/op
BenchmarkQuantile/size=10/quantile=0.25-16                      301586848                3.979 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.50-16                      251046133                4.776 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.75-16                      301175084                3.981 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.99-16                      446203158                2.663 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.25-16                     71674243                16.36 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.50-16                     31303939                38.40 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.75-16                     100000000               10.81 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.99-16                     442189958                2.713 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.25-16                     3489037               337.9 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.50-16                     1786258               653.6 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.75-16                     8720684               137.2 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.99-16                    202513672                5.886 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.25-16                     385941              3221 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.50-16                     190620              6274 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.75-16                     948091              1284 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.99-16                   28220740                43.79 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/ndx-technologies/tdigest     30.658s
```

References

- https://github.com/tdunning/t-digest
- https://github.com/facebook/folly/blob/main/folly/stats/TDigest.cpp (no-unprocessed algorithm, inlined compression in creation, no add)
- https://github.com/MnO2/t-digest (same as folly)
- https://github.com/derrickburns/tdigest (processed/unprocessed algorithm, add via new centroid)
- https://github.com/influxdata/tdigest (last update 4y ago, open issues, uses errors in API, processed/unprocessed algorithm, add via new centroid)
- https://github.com/spenczar/tdigest (archived)
