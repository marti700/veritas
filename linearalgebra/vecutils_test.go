package linearalgebra_test

import (
	"github.com/marti700/veritas/linearalgebra"
	"testing"
)

func TestSize(t *testing.T) {
	v1 := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3, 4, 5, 6},
	})
	v2 := linearalgebra.NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
	})

	ans1 := 6
	ans2 := 4

	result1, _ := linearalgebra.Size(v1)
	result2, _ := linearalgebra.Size(v2)

	if ans1 != result1 {
		t.Error("The answer should be ", ans1, "but was ", result1)
	}
	if ans2 != result2 {
		t.Error("The answer should be ", ans2, "but was ", result2)
	}
}

func TestDotProductRowCol(t *testing.T) {
	v1 := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3, 4, 5, 6},
	})
	v2 := linearalgebra.NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	})

	ans := 91.0
	result, _ := linearalgebra.DotProduct(v1,v2)

	if ans != result {
		t.Error("The answer should be ", ans, "but was ", result)
	}
}
func TestDotProductRowRow(t *testing.T) {
	v1 := linearalgebra.NewMatrix([][]float64{
		{2, 3, 4, 5, 5},
	})
	v2 := linearalgebra.NewMatrix([][]float64{
		{2, 3, 4, 5, 5},
	})

	ans1 := 79.0

	result1, _ := linearalgebra.DotProduct(v1,v2)

	if ans1 != result1 {
		t.Error("The answer should be ", ans1, "but was ", result1)
	}
}
func TestDotProductColCol(t *testing.T) {
	v1 := linearalgebra.NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	})

	v2 := linearalgebra.NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
		{5},
		{30},
	})

	ans1 := 235.0

	result1, _ := linearalgebra.DotProduct(v1,v2)

	if ans1 != result1 {
		t.Error("The answer should be ", ans1, "but was ", result1)
	}
}
