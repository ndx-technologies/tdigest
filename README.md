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
BenchmarkAddCompress/size=10-16                                  8362072               131.7 ns/op           295 B/op          3 allocs/op
BenchmarkAddCompress/size=100-16                                 3880452               315.9 ns/op          1328 B/op          1 allocs/op
BenchmarkAddCompress/size=1000-16                                5255559               256.2 ns/op           862 B/op          0 allocs/op
BenchmarkAddCompress/size=10000-16                               2157460               480.1 ns/op           224 B/op          0 allocs/op
BenchmarkMergeCompress/size=10-16                                  10000            807682 ns/op             465 B/op          2 allocs/op
BenchmarkMergeCompress/size=100-16                                  4974          11731933 ns/op            4294 B/op          2 allocs/op
BenchmarkMergeCompress/size=1000-16                                  408          13009113 ns/op           44409 B/op          2 allocs/op
BenchmarkMergeCompress/size=10000-16                                 100          37038880 ns/op          410557 B/op          2 allocs/op
BenchmarkQuantile/size=10/quantile=0.25-16                      296270143                4.056 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.50-16                      246909472                4.865 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.75-16                      296573799                4.048 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.99-16                      384321291                3.120 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.25-16                     65808562                18.09 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.50-16                     28376620                42.43 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.75-16                     100000000               11.93 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.99-16                     373714771                3.209 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.25-16                     3401463               353.2 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.50-16                     1747731               686.4 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.75-16                     7929656               151.6 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.99-16                    195276451                6.136 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.25-16                     354295              3360 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.50-16                     176695              6763 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.75-16                     868041              1376 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.99-16                   26568044                45.23 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=10-16                        94844790                12.72 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=100-16                       13354563                88.84 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=1000-16                       1464591               818.9 ns/op             0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=10000-16                       147856              8138 ns/op               0 B/op          0 allocs/op
BenchmarkTDigest_MarshalBinary/size=10-16                       41075317                28.73 ns/op          112 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=100-16                       7826073               155.6 ns/op           896 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=1000-16                       884018              1345 ns/op            8192 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=10000-16                       99562             11986 ns/op           81920 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=10-16                     48807934                24.27 ns/op           80 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=100-16                     7245828               166.3 ns/op           896 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=1000-16                     784621              1537 ns/op            8192 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=10000-16                     87409             13640 ns/op           81920 B/op          1 allocs/op
PASS
ok      github.com/ndx-technologies/tdigest     114.440s
```

## References

- https://github.com/tdunning/t-digest
- https://github.com/facebook/folly/blob/main/folly/stats/TDigest.cpp (no-unprocessed algorithm, inlined compression in creation, no add)
- https://github.com/MnO2/t-digest (same as folly)
- https://github.com/derrickburns/tdigest (processed/unprocessed algorithm, add via new centroid)
- https://github.com/influxdata/tdigest (not maintained, last update 4y ago, open issues, uses errors in API, processed/unprocessed algorithm, add via new centroid)
- https://github.com/spenczar/tdigest (archived)
