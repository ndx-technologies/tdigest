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
BenchmarkAddCompress/size=10-16                                  5350582               256.6 ns/op           656 B/op          7 allocs/op
BenchmarkAddCompress/size=100-16                                 1266465               940.6 ns/op          1342 B/op          4 allocs/op
BenchmarkAddCompress/size=1000-16                                 335818              5868 ns/op             621 B/op          3 allocs/op
BenchmarkAddCompress/size=10000-16                                123516             19323 ns/op             671 B/op          3 allocs/op
BenchmarkMergeCompress/size=10-16                                  10000           2000885 ns/op            1228 B/op          7 allocs/op
BenchmarkMergeCompress/size=100-16                                  2288          10416689 ns/op            9813 B/op          7 allocs/op
BenchmarkMergeCompress/size=1000-16                                  188          12622178 ns/op           86224 B/op          7 allocs/op
BenchmarkMergeCompress/size=10000-16                                 100          78736989 ns/op          878768 B/op          7 allocs/op
BenchmarkQuantile/size=10/quantile=0.25-16                      299838242                4.027 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.50-16                      245923518                4.887 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.75-16                      297084456                4.034 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.99-16                      389778069                3.097 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.25-16                     65193768                18.03 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.50-16                     29017426                41.98 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.75-16                     100000000               11.87 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.99-16                     380789940                3.143 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.25-16                     3690636               323.9 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.50-16                     1891216               634.9 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.75-16                     8029159               148.9 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.99-16                    198243957                6.051 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.25-16                     373183              3124 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.50-16                     191590              6237 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.75-16                     883096              1360 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.99-16                   26880088                44.58 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/ndx-technologies/tdigest     92.221s
```

## References

- https://github.com/tdunning/t-digest
- https://github.com/facebook/folly/blob/main/folly/stats/TDigest.cpp (no-unprocessed algorithm, inlined compression in creation, no add)
- https://github.com/MnO2/t-digest (same as folly)
- https://github.com/derrickburns/tdigest (processed/unprocessed algorithm, add via new centroid)
- https://github.com/influxdata/tdigest (not maintained, last update 4y ago, open issues, uses errors in API, processed/unprocessed algorithm, add via new centroid)
- https://github.com/spenczar/tdigest (archived)
