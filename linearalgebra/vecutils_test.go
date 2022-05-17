package linearalgebra

import (
	"testing"
)

func TestSize(t *testing.T) {
	v1 := NewMatrix([][]float64{
		{1, 2, 3, 4, 5, 6},
	})
	v2 := NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
	})

	ans1 := 6
	ans2 := 4

	result1, _ := Size(v1)
	result2, _ := Size(v2)

	if ans1 != result1 {
		t.Error("The answer should be ", ans1, "but was ", result1)
	}
	if ans2 != result2 {
		t.Error("The answer should be ", ans2, "but was ", result2)
	}
}

func TestDotProductRowCol(t *testing.T) {
	v1 := NewMatrix([][]float64{
		{1, 2, 3, 4, 5, 6},
	})
	v2 := NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	})

	ans := 91.0
	result, _ := DotProduct(v1, v2)

	if ans != result {
		t.Error("The answer should be ", ans, "but was ", result)
	}
}
func TestDotProductRowRow(t *testing.T) {
	v1 := NewMatrix([][]float64{
		{2, 3, 4, 5, 5},
	})
	v2 := NewMatrix([][]float64{
		{2, 3, 4, 5, 5},
	})

	ans1 := 79.0

	result1, _ := DotProduct(v1, v2)

	if ans1 != result1 {
		t.Error("The answer should be ", ans1, "but was ", result1)
	}
}
func TestDotProductColCol(t *testing.T) {
	v1 := NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	})

	v2 := NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
		{5},
		{30},
	})

	ans1 := 235.0

	result1, _ := DotProduct(v1, v2)

	if ans1 != result1 {
		t.Error("The answer should be ", ans1, "but was ", result1)
	}
}

func TestIsVector(t *testing.T) {
	v1 := NewColumnVector([]float64{1, 2, 3, 4, 5, 6, 7})
	v2 := NewRowVector([]float64{1, 2, 3, 4, 5, 67})

	if !(IsVector(v1) && IsVector(v2)) {
		t.Error("v1 and v2 are vectors")
	}

	m := NewMatrix([][]float64{
		{2, 0, 1},
		{3, 0, 2},
		{5, -6, 7},
	})

	if IsVector(m) {
		t.Error("M is a matrix, not a vector")
	}

}

func TestIsRowVector(t *testing.T) {
	v1 := NewColumnVector([]float64{1, 2, 3, 4, 5, 6, 7})
	v2 := NewRowVector([]float64{1, 2, 3, 4, 5, 67})

	if !(IsRowVector(v2)) {
		t.Error("v1 and v2 are vectors")
	}

	if IsRowVector(v1) {
		t.Error("v1 is a column vector")
	}

	m := NewMatrix([][]float64{
		{2, 0, 1},
		{3, 0, 2},
		{5, -6, 7},
	})

	if IsRowVector(m) {
		t.Error("M is a matrix, not a Row vector")
	}

}

func TestIsColumnVector(t *testing.T) {
	v1 := NewColumnVector([]float64{1, 2, 3, 4, 5, 6, 7})
	v2 := NewRowVector([]float64{1, 2, 3, 4, 5, 67})

	if !(IsColumnVector(v1)) {
		t.Error("v1 is a column vector")
	}

	if IsColumnVector(v2) {
		t.Error("v2 is a row vector")
	}

	m := NewMatrix([][]float64{
		{2, 0, 1},
		{3, 0, 2},
		{5, -6, 7},
	})

	if IsVector(m) {
		t.Error("M is a matrix, not a column vector")
	}

}

func TestElementWiseFilter(t *testing.T) {
	v1 := NewColumnVector([]float64{1, 2, 3, 4, 5, 6, 7})
	v2 := NewRowVector([]float64{1, 2, 3, 4, 5, 67})

	expected1 := NewColumnVector([]float64{2, 4, 6})
	expected2 := NewRowVector([]float64{2, 4})

	result1 := ElementWiseFilter(v1, func(n float64) bool {
		return (int(n) % 2) == 0
	}, 1)

	result2 := ElementWiseFilter(v2, func(n float64) bool {
		return (int(n) % 2) == 0
	}, 0)

	if !MatrixEq1(result1, expected1) {
		t.Error("answer should be: ", expected1, "but actual was: ", result1)
	}

	if !MatrixEq1(result2, expected2) {
		t.Error("answer should be: ", expected2, "but actual was: ", result2)
	}
}
