t-digest in Go

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
BenchmarkAddCompress/size=10-16                                 10087381               117.0 ns/op           296 B/op          3 allocs/op
BenchmarkAddCompress/size=100-16                                11190367               103.5 ns/op           429 B/op          0 allocs/op
BenchmarkAddCompress/size=1000-16                                7451500               391.8 ns/op          1483 B/op          0 allocs/op
BenchmarkAddCompress/size=10000-16                               1952947               693.2 ns/op           170 B/op          0 allocs/op
BenchmarkMergeCompress/size=10-16                                3163314               354.4 ns/op           525 B/op          6 allocs/op
BenchmarkMergeCompress/size=100-16                                281220              4415 ns/op            2800 B/op          6 allocs/op
BenchmarkMergeCompress/size=1000-16                                20407             58278 ns/op           29055 B/op          6 allocs/op
BenchmarkMergeCompress/size=10000-16                                1336            885533 ns/op          245756 B/op          6 allocs/op
BenchmarkQuantile/size=10/quantile=0.25-16                      300870988                4.006 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.50-16                      250298445                4.794 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.75-16                      300945816                3.996 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10/quantile=0.99-16                      399924069                3.000 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.25-16                     66874720                17.17 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.50-16                     30142264                39.79 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.75-16                     100000000               11.65 ns/op            0 B/op          0 allocs/op
BenchmarkQuantile/size=100/quantile=0.99-16                     387504800                3.090 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.25-16                     3711162               325.4 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.50-16                     1898356               639.2 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.75-16                     8525449               142.8 ns/op             0 B/op          0 allocs/op
BenchmarkQuantile/size=1000/quantile=0.99-16                    196872932                6.033 ns/op           0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.25-16                     385173              3136 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.50-16                     193035              6225 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.75-16                     921289              1263 ns/op               0 B/op          0 allocs/op
BenchmarkQuantile/size=10000/quantile=0.99-16                   28343498                43.12 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=10-16                        96190780                12.43 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=100-16                       13933243                86.83 ns/op            0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=1000-16                       1480885               809.2 ns/op             0 B/op          0 allocs/op
BenchmarkTDigest_AppendBinary/size=10000-16                       150110              8026 ns/op               0 B/op          0 allocs/op
BenchmarkTDigest_MarshalBinary/size=10-16                       42346652                27.51 ns/op          112 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=100-16                       8059311               149.5 ns/op           896 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=1000-16                       916420              1254 ns/op            8192 B/op          1 allocs/op
BenchmarkTDigest_MarshalBinary/size=10000-16                      103988             11453 ns/op           81920 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=10-16                     51248196                23.28 ns/op           80 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=100-16                     7656922               156.5 ns/op           896 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=1000-16                     864814              1425 ns/op            8192 B/op          1 allocs/op
BenchmarkTDigest_UnmarshalBinary/size=10000-16                     93374             12965 ns/op           81920 B/op          1 allocs/op
PASS
ok      github.com/ndx-technologies/tdigest     45.775s
```

References

- https://github.com/tdunning/t-digest
- https://github.com/facebook/folly/blob/main/folly/stats/TDigest.cpp (no-unprocessed algorithm, inlined compression in creation, no add)
- https://github.com/MnO2/t-digest (same as folly)
- https://github.com/derrickburns/tdigest (processed/unprocessed algorithm, add via new centroid)
- https://github.com/influxdata/tdigest (last update 4y ago, open issues, uses errors in API, processed/unprocessed algorithm, add via new centroid)
- https://github.com/spenczar/tdigest (archived)
