package stat

import "math"

type Stat struct {
	Mean      float64
	StdDev    float64
	Count     int
	Sum       float64
	SquareSum float64
}

func (s *Stat) InsertNewValue(value float64) {
	s.Count++
	s.Sum += value
	s.SquareSum += value * value
	s.Mean += s.Sum / float64(s.Count)
	if s.Count > 1 {
		variance := (s.SquareSum / float64(s.Count)) - (s.Mean * s.Mean)
		if variance < 0 {
			variance = 0
		}
		s.StdDev = math.Sqrt(variance)
	}
}

func (s *Stat) CheckAnomaly(value float64, k float64) bool {
	if value > s.Mean {
		if value < s.Mean-k*s.StdDev || value > s.Mean+k*s.StdDev {
			return true
		}
	}
	return false
}
