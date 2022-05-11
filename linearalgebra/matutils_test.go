package linearalgebra_test

import (
	"github.com/marti700/veritas/linearalgebra"
	"github.com/marti700/veritas/linearalgebra/lintest"
	"testing"
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

func TestSumMatOutil(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})

	ans := 10.0
	result := linearalgebra.ElementsSum(m)

	if ans != result {
		t.Error("The anwer should be ", ans, " but was ", result)
	}
}

func TestGenIdentityMatrix(t *testing.T) {
	ans := linearalgebra.NewMatrix([][]float64{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	})

	result := linearalgebra.GenIdenityMatrix(5)

	if !lintest.MatrixEq1(ans, result) {
		t.Error("The anwer should be ", ans, " but was ", result)
	}
}

func TestSlice(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// for the x axis
	ans := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	result := linearalgebra.Slice(m, 0, 2, "x")

	if !lintest.MatrixEq1(ans, result) {
		t.Error("The anwer should be ", ans, " but was ", result)
	}

	// for the y axix

	ans1 := linearalgebra.NewMatrix([][]float64{
		{1, 2},
		{4, 5},
		{7, 8},
	})
	result1 := linearalgebra.Slice(m, 0, 2, "y")

	if !lintest.MatrixEq1(ans1, result1) {
		t.Error("The anwer should be ", ans1, " but was ", result1)
	}

}

func TestFilter(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	ans := linearalgebra.NewMatrix([][]float64{
		{4, 5, 6},
		{7, 8, 9},
	})

	result := linearalgebra.Filter(m, func(r linearalgebra.Matrix) bool {
		return r.Get(0,0) > 2
	},0)

	if !lintest.MatrixEq1(ans, result) {
		t.Error("The anwer should be ", ans, " but was ", result)
	}
}
