package tdigest

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
)

func ExampleTDigest_Quantile() {
	d := TDigest{MaxSize: 100}

	for i := float32(1); i <= 100; i++ {
		d.Add(i, 1)
	}

	fmt.Println(d.Count, d.Sum, d.Min, d.Max)
	fmt.Println(d.Quantile(0.001), d.Quantile(0.01), d.Quantile(0.5), d.Quantile(0.99), d.Quantile(0.999))
	// Output: 100 5050 1 100
	// 1 1.5 50.5 99.5 100
}

func ExampleTDigest_Quantile_one() {
	d := TDigest{MaxSize: 1}
	d.Add(1, 1)

	fmt.Println(d.Count, d.Sum, d.Min, d.Max)
	fmt.Println(d.Quantile(0.001), d.Quantile(0.01), d.Quantile(0.5), d.Quantile(0.99), d.Quantile(0.999))
	// Output: 1 1 1 1
	// 1 1 1 1 1
}

func ExampleTDigest_Quantile_empty() {
	var d TDigest

	fmt.Println(d.Count, d.Sum, d.Min, d.Max)
	fmt.Println(d.Quantile(0.01), d.Quantile(0.5), d.Quantile(1))
	// Output: 0 0 0 0
	// 0 0 0
}

func ExampleTDigest_Merge() {
	d := TDigest{MaxSize: 100}
	for i := float32(1); i <= 100; i++ {
		d.Add(i, 1)
	}

	b := TDigest{MaxSize: 100}
	for i := float32(101); i <= 200; i++ {
		b.Add(i, 1)
	}

	d.Merge(b)
	d.Merge()

	fmt.Println(d.Count, d.Sum, d.Min, d.Max)
	fmt.Println(d.Quantile(0.001), d.Quantile(0.01), d.Quantile(0.5), d.Quantile(0.99), d.Quantile(0.999))
	// Output: 200 20100 1 200
	// 1 2.5 100.5 198.5 200
}

func ExampleTDigest_Merge_compress() {
	d := TDigest{MaxSize: 100}
	for i := float32(1); i <= 100; i++ {
		d.Add(i, 1)
	}

	b := TDigest{MaxSize: 100}
	for i := float32(101); i <= 200; i++ {
		b.Add(i, 1)
	}

	d.Merge(b)
	d.Merge()
	d.Compress()

	fmt.Println(d.Count, d.Sum, d.Min, d.Max)
	fmt.Println(d.Quantile(0.001), d.Quantile(0.01), d.Quantile(0.5), d.Quantile(0.99), d.Quantile(0.999))
	// Output: 200 20100 1 200
	// 1 2.5 100.25 198.5 200
}

func ExampleTDigest_Merge_large() {
	d := TDigest{MaxSize: 100}
	for i := float32(1); i <= 1000; i++ {
		d.Add(i, 1)
	}
	d.Compress()

	fmt.Println(d.Count, d.Sum, d.Min, d.Max)
	fmt.Println(d.Quantile(0.001), d.Quantile(0.01), d.Quantile(0.5), d.Quantile(0.99), d.Quantile(0.999))
	// Output: 1000 500500 1 1000
	// 1.5 10.5 500.25 990.25 999.5
}

func ExampleTDigest_Merge_negative() {
	d := TDigest{MaxSize: 100}
	for i := float32(1); i <= 100; i++ {
		d.Add(i, 1)
		d.Add(-i, 1)
	}
	d.Compress()

	fmt.Println(d.Count, d.Sum, d.Min, d.Max)
	fmt.Println(d.Quantile(0), d.Quantile(0.001), d.Quantile(0.01), d.Quantile(0.99), d.Quantile(0.999), d.Quantile(1))
	// Output: 200 0 -100 100
	// -100 -100 -98.5 98.5 100 100
}

func TestMergeLargeAsDigests(t *testing.T) {
	values := make([]float32, 1000)
	for i := range values {
		values[i] = float32(i + 1)
	}
	rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })

	digests := make([]TDigest, 0, 10)
	for i := range 10 {
		d := TDigest{MaxSize: 100}
		for j := i * 100; j < (i+1)*100; j++ {
			d.Add(values[j], 1)
		}
		digests = append(digests, d)
	}

	d := digests[0]
	d.Merge(digests[1:]...)
	d.Compress()

	if d.Count != 1000 {
		t.Error(d.Count)
	}
	if d.Sum != 500500 {
		t.Error(d.Sum)
	}
	if mean := d.Sum / float32(d.Count); mean != 500.5 {
		t.Error(mean)
	}
	if d.Min != 1 {
		t.Error(d.Min)
	}
	if d.Max != 1000 {
		t.Error(d.Max)
	}

	tests := map[float32]float32{
		0.001: 1.5,
		0.01:  10.5,
		0.5:   500.25,
		0.99:  990.25,
		0.999: 999.5,
	}
	for q, v := range tests {
		if g := d.Quantile(q); g != v {
			t.Error(g, q, v)
		}
	}
}

func TestNegativeValuesMergeDigests(t *testing.T) {
	d := TDigest{MaxSize: 100}
	for i := float32(1); i <= 100; i++ {
		d.Add(i, 1)
	}

	d2 := TDigest{MaxSize: 100}
	for i := float32(1); i <= 100; i++ {
		d2.Add(-i, 1)
	}

	d.Merge(d2)
	d.Compress()

	if d.Count != 200 {
		t.Error(d.Count)
	}
	if d.Sum != 0 {
		t.Error(d.Sum)
	}
	if mean := d.Sum / float32(d.Count); mean != 0 {
		t.Error(mean)
	}
	if d.Min != -100 {
		t.Error(d.Min)
	}
	if d.Max != 100 {
		t.Error(d.Max)
	}

	tests := map[float32]float32{
		0.0:   -100,
		0.001: -100,
		0.01:  -98.5,
		0.99:  98.5,
		0.999: 100,
		1.0:   100,
	}
	for q, v := range tests {
		if g := d.Quantile(q); g != v {
			t.Error(g, q, v)
		}
	}
}

func TestLargeOutlier(t *testing.T) {
	d := TDigest{MaxSize: 100}
	for i := range 20 {
		v := float32(i)
		if i == 0 {
			v = 100_000.0
		}
		d.Add(v, 1)
	}
	d.Compress()

	if q50, q90 := d.Quantile(0.5), d.Quantile(0.90); q50 >= q90 {
		t.Error(q50, q90)
	}
}

func TestFloatingPointSorted(t *testing.T) {
	const v = 1.4

	d1 := TDigest{MaxSize: 100}
	for range 100 {
		d1.Add(v, 1)
	}

	d2 := TDigest{MaxSize: 100}
	for range 100 {
		d1.Add(v, 1)
	}

	md1 := TDigest{MaxSize: 100}
	md1.Merge(d1)
	md1.Merge(d2)
	md1.Compress()

	d3 := TDigest{MaxSize: 100}
	for range 100 {
		d3.Add(v, 1)
	}

	md2 := TDigest{MaxSize: 100}
	md2.Merge(md1)
	md2.Merge(d3)
	md2.Compress()

	for i := 1; i < len(md2.Centroids); i++ {
		if md2.Centroids[i-1].Mean > md2.Centroids[i].Mean {
			t.Error(md2.Centroids[i-1].Mean, md2.Centroids[i].Mean)
		}
	}
}

const (
	kNumSamples    = 3000
	kNumRandomRuns = 10
	kSeed          = 0
)

func TestDistribution(t *testing.T) {
	dists := []struct {
		logarithmic bool
		modes       int
	}{
		{logarithmic: true, modes: 1},
		{logarithmic: true, modes: 3},
		{logarithmic: false, modes: 1},
		{logarithmic: false, modes: 10},
	}

	for _, dist := range dists {
		for _, q := range []float32{0.001, 0.01, 0.25, 0.5, 0.75, 0.99, 0.999} {
			for _, merge := range []bool{true, false} {
				t.Run("", func(t *testing.T) {
					var reasonableError float32
					switch q {
					case 0.001, 0.999:
						reasonableError = 0.001
					case 0.01, 0.99:
						reasonableError = 0.01
					default:
						reasonableError = 0.04
					}

					errors := make([]float32, kNumRandomRuns)

					for i := range errors {
						values := make([]float32, kNumSamples)
						if dist.logarithmic {
							for j := range values {
								mode := rand.Intn(dist.modes)
								values[j] = float32(math.Log(rand.Float64()*math.E+1) + 100*float64(mode))
							}
						} else {
							for j := range values {
								mode := rand.Intn(dist.modes)
								values[j] = float32(rand.NormFloat64()*25 + 100*float64(mode+1))
							}
						}

						var d TDigest

						if merge {
							digests := make([]TDigest, kNumSamples/1000)

							for j := range digests {
								for i := j * 1000; i < (j+1)*1000; i++ {
									digests[j].Add(values[i], 1)
								}
								digests[j].Compress()
							}

							d = digests[0]
							d.Merge(digests[1:]...)
						} else {
							d = TDigest{MaxSize: 100}

							for j := 0; j < kNumSamples/1000; j++ {
								for i := j * 1000; i < (j+1)*1000; i++ {
									d.Add(values[i], 1)
								}
							}

							d.Compress()
						}

						sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
						est := d.Quantile(q)
						idx := sort.Search(len(values), func(i int) bool { return values[i] >= est })
						actual := float32(idx) / float32(kNumSamples)
						errors[i] = actual - q
					}

					var mean, variance float32
					for _, e := range errors {
						mean += e
					}
					mean /= float32(kNumRandomRuns)

					for _, e := range errors {
						variance += (e - mean) * (e - mean)
					}
					stddev := math.Sqrt(float64(variance) / float64(kNumRandomRuns-1))

					if stddev > float64(reasonableError) {
						t.Error(stddev)
					}
				})
			}
		}
	}
}

func BenchmarkNoop______________________________________(b *testing.B) {}

func BenchmarkAddCompress(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			d := TDigest{MaxSize: size}
			v := rand.Float32() * 1000

			for b.Loop() {
				d.Add(v, 1.0)
				if len(d.Centroids) > d.MaxSize {
					d.Compress()
				}
			}
		})
	}
}

func BenchmarkMergeCompress(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			d := TDigest{MaxSize: size}
			other := TDigest{MaxSize: size}
			for range size {
				other.Add(rand.Float32()*1000, 1.0)
			}
			other.Compress()

			for b.Loop() {
				d.Merge(other)
				other.Compress()
			}
		})
	}
}

func BenchmarkQuantile(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000} {
		for _, q := range []float32{0.25, 0.5, 0.75, 0.99} {
			b.Run(fmt.Sprintf("size=%d/quantile=%.2f", size, q), func(b *testing.B) {
				d := TDigest{MaxSize: size}
				for range size {
					d.Add(rand.Float32()*1000, 1.0)
				}

				for b.Loop() {
					d.Quantile(q)
				}
			})
		}
	}
}

func FuzzCentroid_AppendBinary(f *testing.F) {
	f.Add([]byte{1, 2, 3}, float32(4), float32(5))

	f.Fuzz(func(t *testing.T, out []byte, mean, weight float32) {
		a := Centroid{Mean: mean, Weight: weight}

		outBefore := make([]byte, len(out))
		copy(outBefore, out)

		out, err := a.AppendBinary(out)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(outBefore, out[:len(outBefore)]) {
			t.Error(outBefore, out)
		}

		var b Centroid
		if err := (&b).UnmarshalBinary(out[len(outBefore):]); err != nil {
			t.Error(err)
		}

		if a != b {
			t.Error(a, b)
		}
	})
}

func FuzzCentroid_EncodeDecode(f *testing.F) {
	f.Add(float32(4), float32(5))

	f.Fuzz(func(t *testing.T, mean, weight float32) {
		a := Centroid{Mean: mean, Weight: weight}

		data, err := a.MarshalBinary()
		if err != nil {
			t.Error(err)
		}

		var b Centroid
		if err := (&b).UnmarshalBinary(data); err != nil {
			t.Error(err)
		}

		if a != b {
			t.Error(a, b)
		}
	})
}

func FuzzTDigest_AppendBinary(f *testing.F) {
	f.Add([]byte{1, 2, 3}, 10)

	f.Fuzz(func(t *testing.T, out []byte, n int) {
		if n < 0 {
			n = -n
		}
		if n > 100_000 {
			n = 100_000
		}

		d := TDigest{MaxSize: n}
		for range n {
			d.Add(rand.Float32()*1000, 1)
		}
		d.Compress()

		outBefore := make([]byte, len(out))
		copy(outBefore, out)

		out, err := d.AppendBinary(out)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(outBefore, out[:len(outBefore)]) {
			t.Error(outBefore, out)
		}

		var dout TDigest
		if err := (&dout).UnmarshalBinary(out[len(outBefore):]); err != nil {
			t.Error(err)
		}

		if !d.Equal(dout) {
			t.Error(d, dout)
		}
	})
}

func FuzzTDigest_EncodeDecode(f *testing.F) {
	f.Add(10)

	f.Fuzz(func(t *testing.T, n int) {
		if n < 0 {
			n = -n
		}
		if n > 100_000 {
			n = 100_000
		}

		a := TDigest{MaxSize: n}
		for range n {
			a.Add(rand.Float32()*1000, 1)
		}
		a.Compress()

		data, err := a.MarshalBinary()
		if err != nil {
			t.Error(err)
		}

		var b TDigest
		if err := (&b).UnmarshalBinary(data); err != nil {
			t.Error(err)
		}

		if !a.Equal(b) {
			t.Error(a, b)
		}
	})
}

func BenchmarkTDigest_AppendBinary(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10_000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			d := TDigest{MaxSize: size}
			for range size {
				d.Add(rand.Float32()*1000, 1)
			}
			d.Compress()

			out := make([]byte, 0, 24+8*size+100)

			for b.Loop() {
				d.AppendBinary(out)
			}
		})
	}
}

func BenchmarkTDigest_MarshalBinary(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10_000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			d := TDigest{MaxSize: size}
			for range size {
				d.Add(rand.Float32()*1000, 1)
			}
			d.Compress()

			for b.Loop() {
				d.MarshalBinary()
			}
		})
	}
}

func BenchmarkTDigest_UnmarshalBinary(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10_000} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			d := TDigest{MaxSize: size}
			for range size {
				d.Add(rand.Float32()*1000, 1)
			}
			d.Compress()

			var a TDigest
			data, _ := d.MarshalBinary()

			for b.Loop() {
				(&a).UnmarshalBinary(data)
			}
		})
	}
}
