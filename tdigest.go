package tdigest

import (
	"encoding/binary"
	"errors"
	"math"
	"slices"
	"sort"
)

var byteOrder = binary.LittleEndian

type Centroid struct {
	Mean   float32
	Weight float32
}

func (s Centroid) AppendBinary(b []byte) ([]byte, error) {
	b = byteOrder.AppendUint32(b, math.Float32bits(s.Mean))
	b = byteOrder.AppendUint32(b, math.Float32bits(s.Weight))
	return b, nil
}

func (s Centroid) MarshalBinary() ([]byte, error) { return s.AppendBinary(make([]byte, 0, 8)) }

func (s *Centroid) UnmarshalBinary(data []byte) error {
	if len(data) != 8 {
		return errors.New("invalid input")
	}
	s.Mean = math.Float32frombits(byteOrder.Uint32(data[:4]))
	s.Weight = math.Float32frombits(byteOrder.Uint32(data[4:8]))
	return nil
}

func (c *Centroid) Add(sum, weight float32) {
	sum += c.Mean * c.Weight
	c.Weight += weight
	c.Mean = sum / c.Weight
}

type TDigest struct {
	Centroids []Centroid
	Count     int
	Sum       float32
	Max       float32
	Min       float32
	MaxSize   int
}

func (s TDigest) Equal(other TDigest) bool {
	return s.Count == other.Count &&
		s.Sum == other.Sum &&
		s.Min == other.Min &&
		s.Max == other.Max &&
		s.MaxSize == other.MaxSize &&
		slices.Equal(s.Centroids, other.Centroids)
}

func (s TDigest) AppendBinary(b []byte) ([]byte, error) {
	if s.Count > math.MaxUint32 || s.MaxSize > math.MaxUint32 || len(s.Centroids) > math.MaxUint32 {
		return b, errors.New("too large for uint32 based encoding")
	}

	// metadata header first
	b = byteOrder.AppendUint32(b, uint32(s.Count))
	b = byteOrder.AppendUint32(b, math.Float32bits(s.Sum))
	b = byteOrder.AppendUint32(b, math.Float32bits(s.Min))
	b = byteOrder.AppendUint32(b, math.Float32bits(s.Max))
	b = byteOrder.AppendUint32(b, uint32(s.MaxSize))

	// centroids
	b = byteOrder.AppendUint32(b, uint32(len(s.Centroids)))

	var err error
	for _, c := range s.Centroids {
		b, err = c.AppendBinary(b)
		if err != nil {
			return b, err
		}
	}

	return b, nil
}

func (s TDigest) MarshalBinary() ([]byte, error) {
	return s.AppendBinary(make([]byte, 0, 8*len(s.Centroids)+4*6))
}

func (s *TDigest) UnmarshalBinary(data []byte) error {
	if len(data) < 24 {
		return errors.New("data too short")
	}

	s.Count = int(byteOrder.Uint32(data[:4]))
	s.Sum = math.Float32frombits(byteOrder.Uint32(data[4:8]))
	s.Min = math.Float32frombits(byteOrder.Uint32(data[8:12]))
	s.Max = math.Float32frombits(byteOrder.Uint32(data[12:16]))
	s.MaxSize = int(byteOrder.Uint32(data[16:20]))

	lenCentroids := int(byteOrder.Uint32(data[20:24]))

	s.Centroids = make([]Centroid, lenCentroids)

	data = data[24:]
	offset := 0

	for i := range s.Centroids {
		if offset+8 > len(data) {
			return errors.New("invalid centroid data")
		}

		if err := (&s.Centroids[i]).UnmarshalBinary(data[offset : offset+8]); err != nil {
			return err
		}

		offset += 8
	}

	return nil
}

// Add value without compression
func (s *TDigest) Add(v, w float32) {
	if s.Count == 0 {
		s.Min = v
		s.Max = v
	}

	s.Sum += v
	s.Count++
	if v > s.Max {
		s.Max = v
	}
	if v < s.Min {
		s.Min = v
	}

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
			if d.Min < s.Min {
				s.Min = d.Min
			}
			if d.Max > s.Max {
				s.Max = d.Max
			}
		}
	}

	sort.Slice(s.Centroids, func(i, j int) bool { return s.Centroids[i].Mean < s.Centroids[j].Mean })
}

func (s *TDigest) Compress() {
	if len(s.Centroids) == 0 {
		return
	}

	compressed := make([]Centroid, 0, s.MaxSize)

	var kLimit float32 = 1.0
	qLimitTimesCount := kToQ(kLimit, float32(s.MaxSize)) * float32(s.Count)

	cur := s.Centroids[0]
	weightSoFar := cur.Weight
	var sumsToMerge, weightsToMerge float32

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
			qLimitTimesCount = kToQ(kLimit, float32(s.MaxSize)) * float32(s.Count)
			kLimit++
			cur = next
		}
	}
	compressed = append(compressed, cur)

	sort.Slice(compressed, func(i, j int) bool { return compressed[i].Mean < compressed[j].Mean })

	s.Centroids = compressed
}

func (s TDigest) Quantile(q float32) float32 {
	if len(s.Centroids) == 0 {
		return 0
	}
	rank := q * float32(s.Count)

	var pos int
	var t float32
	if q > 0.5 {
		if q >= 1 {
			return s.Max
		}
		t = float32(s.Count)
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

	var delta float32
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
