package tdigest

import "sort"

type float interface{ ~float32 | ~float64 }

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Centroid[T float] struct{ Mean, Weight T }

func (c *Centroid[T]) Add(sum, weight T) {
	sum += c.Mean * c.Weight
	c.Weight += weight
	c.Mean = sum / c.Weight
}

type TDigest[T float] struct {
	Count         int
	Sum, Min, Max T
	Centroids     []Centroid[T]
}

func (s TDigest[T]) IsZero() bool { return s.Count == 0 }

// Insert value without compression
func (s *TDigest[T]) Insert(v, w T) {
	if s.Count == 0 {
		s.Min, s.Max = v, v
	}

	s.Sum += v
	s.Count++
	s.Max = max(s.Max, v)
	s.Min = min(s.Min, v)

	pos := sort.Search(len(s.Centroids), func(i int) bool { return s.Centroids[i].Mean >= v })

	s.Centroids = append(s.Centroids, Centroid[T]{})
	if pos < len(s.Centroids)-1 {
		copy(s.Centroids[pos+1:], s.Centroids[pos:len(s.Centroids)-1])
	}

	s.Centroids[pos] = Centroid[T]{Mean: v, Weight: w}
}

// Merge without compression
func (s *TDigest[T]) Merge(digests ...TDigest[T]) {
	for _, d := range digests {
		if d.Count > 0 {
			if s.Count == 0 {
				s.Min, s.Max = d.Min, d.Max
			} else {
				s.Min = min(s.Min, d.Min)
				s.Max = max(s.Max, d.Max)
			}

			s.Count += d.Count
			s.Sum += d.Sum
			s.Centroids = append(s.Centroids, d.Centroids...)
		}
	}

	sort.Slice(s.Centroids, func(i, j int) bool { return s.Centroids[i].Mean < s.Centroids[j].Mean })
}

func (s *TDigest[T]) Compress(maxCentroids int) {
	if len(s.Centroids) == 0 || len(s.Centroids) <= maxCentroids {
		return
	}

	compressed := make([]Centroid[T], 0, maxCentroids)

	var kLimit T = 1
	qLimitTimesCount := kToQ(kLimit, T(maxCentroids)) * T(s.Count)

	cur := s.Centroids[0]
	weightSoFar := cur.Weight
	var sumsToMerge, weightsToMerge T

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
			qLimitTimesCount = kToQ(kLimit, T(maxCentroids)) * T(s.Count)
			kLimit++
			cur = next
		}
	}
	compressed = append(compressed, cur)

	sort.Slice(compressed, func(i, j int) bool { return compressed[i].Mean < compressed[j].Mean })

	s.Centroids = compressed
}

func (s TDigest[T]) Quantile(q T) T {
	if len(s.Centroids) == 0 {
		return 0
	}
	rank := q * T(s.Count)

	var pos int
	var t T
	if q > 0.5 {
		if q >= 1 {
			return s.Max
		}
		t = T(s.Count)
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

	var delta T
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

func (s *TDigest[T]) Mul(v T) {
	s.Sum *= v
	s.Min *= v
	s.Max *= v
	for i := range s.Centroids {
		s.Centroids[i].Mean *= v
	}
}

func kToQ[T float](k, d T) T {
	kDivD := k / d
	if kDivD >= 0.5 {
		base := 1 - kDivD
		return 1 - (2 * base * base)
	}
	return 2 * kDivD * kDivD
}

func clamp[T float | integer](v, lo, hi T) T {
	switch {
	case v > hi:
		return hi
	case v < lo:
		return lo
	default:
		return v
	}
}
