// This file has descriptive statistics functions

package stats

// returns the average of a slice of float64 numbers
func Mean(data []float64) float64 {
	var sum float64
	for _, elm := range data {
		sum += elm
	}

	return sum / float64(len(data))
}

// returns the minimun value in this slice
func Min(s []float64) float64 {
	min := s[0]
	for _, val := range s {
		if min > val {
			min = val
		}
	}
	return min
}