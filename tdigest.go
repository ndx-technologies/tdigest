package tdigest

import (
	"math"
	"sort"
)

type Centroid struct {
	Mean   float64
	Weight float64
}

/*
func (s Centroid) AppendBinary(b []byte) ([]byte, error) {
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(float32(s.Mean)))
	b = binary.LittleEndian.AppendUint32(b, math.Float32bits(float32(s.Weight)))
	return b, nil
}

func (s Centroid) MarshalBinary() ([]byte, error) { return s.AppendBinary(make([]byte, 0, 32)) }

func (s *Centroid) UnmarshalBinary(data []byte) error {
	if len(data) != 32 {
		return ErrInvalidDataLength{Length: len(data)}
	}
	for i := 0; i < 4; i++ {
		s.words[i] = binary.LittleEndian.Uint64(data[i*8 : (i+1)*8])
	}
	return nil
}
*/

func (c *Centroid) Add(sum, weight float64) {
	sum += c.Mean * c.Weight
	c.Weight += weight
	c.Mean = sum / c.Weight
}

type TDigest struct {
	Centroids []Centroid
	Sum       float64
	Count     int
	Max       float64
	Min       float64
	MaxSize   int
}

// Add value without compression
func (s *TDigest) Add(v, w float64) {
	if s.Count == 0 {
		s.Min = v
		s.Max = v
	}

	s.Sum += v
	s.Count++
	s.Max = math.Max(s.Max, v)
	s.Min = math.Min(s.Min, v)

	s.Centroids = append(s.Centroids, Centroid{Mean: v, Weight: w})
	sort.Slice(s.Centroids, func(i, j int) bool { return s.Centroids[i].Mean < s.Centroids[j].Mean })
}

// Merge without compression
func (s *TDigest) Merge(digests ...TDigest) {
	for _, d := range digests {
		if d.Count > 0 {
			s.Centroids = append(s.Centroids, d.Centroids...)
			s.Count += d.Count
			s.Sum += d.Sum
			s.Min = math.Min(s.Min, d.Min)
			s.Max = math.Max(s.Max, d.Max)
		}
	}

	sort.Slice(s.Centroids, func(i, j int) bool { return s.Centroids[i].Mean < s.Centroids[j].Mean })
}

func (s *TDigest) Compress() {
	compressed := make([]Centroid, 0, s.MaxSize)

	kLimit := 1.0
	qLimitTimesCount := kToQ(kLimit, float64(s.MaxSize)) * float64(s.Count)

	cur := s.Centroids[0]
	weightSoFar := cur.Weight
	sumsToMerge := 0.0
	weightsToMerge := 0.0

	for _, next := range s.Centroids[1:] {
		weightSoFar += next.Weight
		if weightSoFar <= qLimitTimesCount {
			sumsToMerge += next.Mean * next.Weight
			weightsToMerge += next.Weight
		} else {
			cur.Add(sumsToMerge, weightsToMerge)
			sumsToMerge = 0
			weightsToMerge = 0
			compressed = append(compressed, cur)
			qLimitTimesCount = kToQ(kLimit, float64(s.MaxSize)) * float64(s.Count)
			kLimit++
			cur = next
		}
	}
	compressed = append(compressed, cur)

	sort.Slice(compressed, func(i, j int) bool { return compressed[i].Mean < compressed[j].Mean })

	s.Centroids = compressed
}

func (s TDigest) Quantile(q float64) float64 {
	if len(s.Centroids) == 0 {
		return 0
	}
	rank := q * float64(s.Count)

	var pos int
	var t float64
	if q > 0.5 {
		if q >= 1 {
			return s.Max
		}
		t = float64(s.Count)
		for i := len(s.Centroids) - 1; i >= 0; i-- {
			t -= s.Centroids[i].Weight
			if rank >= t {
				pos = i
				break
			}
		}
	} else {
		if q <= 0 {
			return s.Min
		}
		pos = len(s.Centroids) - 1
		t = 0
		for i, c := range s.Centroids {
			if rank < t+c.Weight {
				pos = i
				break
			}
			t += c.Weight
		}
	}

	var delta float64
	min, max := s.Min, s.Max
	if len(s.Centroids) > 1 {
		if pos == 0 {
			delta = s.Centroids[pos+1].Mean - s.Centroids[pos].Mean
			max = s.Centroids[pos+1].Mean
		} else if pos == len(s.Centroids)-1 {
			delta = s.Centroids[pos].Mean - s.Centroids[pos-1].Mean
			min = s.Centroids[pos-1].Mean
		} else {
			delta = (s.Centroids[pos+1].Mean - s.Centroids[pos-1].Mean) / 2
			min = s.Centroids[pos-1].Mean
			max = s.Centroids[pos+1].Mean
		}
	}

	value := s.Centroids[pos].Mean + ((rank-t)/s.Centroids[pos].Weight-0.5)*delta
	return clamp(value, min, max)
}

type number interface{ float | integer }

type float interface{ ~float32 | ~float64 }

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func kToQ[T float](k, d T) T {
	kDivD := k / d
	if kDivD >= 0.5 {
		base := 1 - kDivD
		return 1 - 2*base*base
	}
	return 2 * kDivD * kDivD
}

func clamp[T number](v, lo, hi T) T {
	if v > hi {
		return hi
	} else if v < lo {
		return lo
	}
	return v
}
