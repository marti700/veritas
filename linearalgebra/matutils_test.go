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
		return r.Get(0, 0) > 2
	}, 0)

	if !lintest.MatrixEq1(ans, result) {
		t.Error("The anwer should be ", ans, " but was ", result)
	}

	m1 := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	ans1 := linearalgebra.NewMatrix([][]float64{
		{2, 3},
		{5, 6},
		{8, 9},
	})

	result1 := linearalgebra.Filter(m1, func(r linearalgebra.Matrix) bool {
		return r.Get(0, 0) >= 2
	}, 1)

	if !lintest.MatrixEq1(ans1, result1) {
		t.Error("The anwer should be ", ans1, " but was ", result1)
	}
}

func TestFilter2(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{0, 8, 9},
	})

	ansT := linearalgebra.NewMatrix([][]float64{
		{4, 5, 6},
		{7, 8, 9},
	})

	ansF := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{0, 8, 9},
	})


	resultT, resultF := linearalgebra.Filter2(m, func(r linearalgebra.Matrix) bool {
		return r.Get(0, 0) > 2
	}, 0)

	if !lintest.MatrixEq1(ansT, resultT) {
		t.Error("The anwer should be ", ansT, " but was ", resultT)
	}

if !lintest.MatrixEq1(ansF, resultF) {
		t.Error("The anwer should be ", ansF, " but was ", resultF)
	}

	m1 := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3, 0},
		{4, 5, 6, 1},
		{7, 8, 9, 3},
	})

	ans1T := linearalgebra.NewMatrix([][]float64{
		{2, 3},
		{5, 6},
		{8, 9},
	})

ans1F := linearalgebra.NewMatrix([][]float64{
		{1, 0},
		{4, 1},
		{7, 3},
	})

	result1T, result1F := linearalgebra.Filter2(m1, func(r linearalgebra.Matrix) bool {
		return r.Get(0, 0) >= 2
	}, 1)

	if !lintest.MatrixEq1(ans1T, result1T) {
		t.Error("The anwer should be ", ans1T, " but was ", result1T)
	}

if !lintest.MatrixEq1(ans1F, result1F) {
		t.Error("The anwer should be ", ans1F, " but was ", result1F)
	}
}

func TestInsert(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{{1, 2}, {3, 4}})
	ans := linearalgebra.NewMatrix([][]float64{{8, 9}, {1, 2}, {3, 4}})
	result := m.InsertAt(linearalgebra.NewMatrix([][]float64{{8, 9}}), 0)

	if !lintest.MatrixEq1(result, ans) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}

	m1 := linearalgebra.NewMatrix([][]float64{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}})
	ans1 := linearalgebra.NewMatrix([][]float64{{0, 1, 1, 2}, {3, 2, 4, 5}, {6, 3, 7, 8}, {9, 4, 10, 11}})
	result1 := m1.InsertAt(linearalgebra.NewMatrix([][]float64{{1}, {2}, {3}, {4}}), 1)

	if !lintest.MatrixEq1(result1, ans1) {
		t.Error("answer should be: ", ans1, "but was: ", result1)
	}
}