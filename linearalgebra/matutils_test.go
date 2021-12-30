package linearalgebra_test

import (
	"testing"
	"github.com/marti700/veritas/linearalgebra"
	"github.com/marti700/veritas/linearalgebra/lintest"
)

func TestOnes(t *testing.T) {
	ans := linearalgebra.NewMatrix([][]float64{
		{1, 1},
		{1, 1},
	})

	ans1 := linearalgebra.NewMatrix([][]float64{
		{1, 1, 1},
	})

	ans2 := linearalgebra.NewMatrix([][]float64{
		{1},
		{1},
		{1},
		{1},
	})

	result1 := linearalgebra.Ones(2, 2)
	result2 := linearalgebra.Ones(1, 3)
	result3 := linearalgebra.Ones(4, 1)

	if !lintest.MatrixEq1(ans, result1) {
		t.Error("The anwer should be ", ans, " but was ", result1)
	}

	if !lintest.MatrixEq1(ans1, result2) {
		t.Error("The anwer should be ", ans1, " but was ", result2)
	}

	if !lintest.MatrixEq1(ans2, result3) {
		t.Error("The anwer should be ", ans2, " but was ", result3)
	}
}
