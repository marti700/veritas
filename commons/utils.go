package commons

func Sum(arr []float64) float64 {
	var acc float64
	for _, e := range arr {
		acc += e
	}
	return acc
}