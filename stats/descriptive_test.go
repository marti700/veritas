package stats

import "testing"

func TestAverage(t *testing.T) {
	v := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := 5.5

	result := Mean(v)

	if result != ans {
		t.Error("Result was: ", result, "But ", ans, "Was expected")
	}
}

func TestMin(t *testing.T) {
	v := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := 1.0

	result := Min(v)

	if result != ans {
		t.Error("Result was: ", result, "But ", ans, "Was expected")
	}
}
