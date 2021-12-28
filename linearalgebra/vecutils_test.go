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
