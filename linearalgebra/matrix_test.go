package linearalgebra_test

import (
	"fmt"
	"testing"

	"github.com/marti700/veritas/linearalgebra"
	"github.com/marti700/veritas/linearalgebra/lintest"
)

func TestTranspose(t *testing.T) {
	mat := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	ans := mat.T()

	expectedAnswer := linearalgebra.NewMatrix([][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	})

	if !lintest.MatrixEq1(ans, expectedAnswer) {
		t.Error("Expected result is: ", expectedAnswer, "but resultWas: ", ans)
	}
}

func TestMult(t *testing.T) {
	m1 := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	m2 := linearalgebra.NewMatrix([][]float64{
		{10, 11},
		{20, 21},
		{30, 31},
	})

	result := linearalgebra.NewMatrix([][]float64{
		{140, 146},
		{320, 335},
	})

	ans, _ := m1.Mult(m2)
	if !lintest.MatrixEq1(ans, result) {
		fmt.Println(ans)
		t.Error("answer should be: ", result, "but was: ", ans)
	}

}

func TestMult1(t *testing.T) {
	m1 := linearalgebra.NewMatrix([][]float64{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9, 10, 11},
	})
	m2 := linearalgebra.NewMatrix([][]float64{
		{0, 3, 6, 9},
		{1, 4, 7, 10},
		{2, 5, 8, 11},
	})

	result := linearalgebra.NewMatrix([][]float64{
		{5, 14, 23, 32},
		{14, 50, 86, 122},
		{23, 86, 149, 212},
		{32, 122, 212, 302},
	})

	ans, _ := m1.Mult(m2)
	if !lintest.MatrixEq1(ans, result) {
		fmt.Println(ans)
		t.Error("answer should be: ", result, "but was: ", ans)
	}

}

func TestInsertCol(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{{1, 2}, {3, 4}})
	ans := linearalgebra.NewMatrix([][]float64{{8,9},{1, 2}, {3, 4}})
	result, _ := m.InsertAt(linearalgebra.NewMatrix([][]float64{{8,9}}), 0)

	if !lintest.MatrixEq1(result, ans) {
		fmt.Println(ans)
		t.Error("answer should be: ", ans, "but was: ", result)
	}

	m1 := linearalgebra.NewMatrix([][]float64{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}})
	ans1 := linearalgebra.NewMatrix([][]float64{{0, 1, 1, 2}, {3, 2, 4, 5}, {6, 3, 7, 8}, {9, 4, 10, 11}})
	result1, _ := m1.InsertAt(linearalgebra.NewMatrix([][]float64{{1}, {2}, {3}, {4}}), 1)

	if !lintest.MatrixEq1(result1, ans1) {
		t.Error("answer should be: ", ans1, "but was: ", result1)
	}
}

func TestGetRow(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	ans := linearalgebra.NewMatrix([][]float64{{1, 2, 3}})
	result := m.GetRow(0)

	if !lintest.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestGetCol(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1,2},
		{3,4},
		{5,6},
	})

	ans := linearalgebra.NewMatrix([][]float64{{1}, {3}, {5}})
	result := m.GetCol(0)

	if !lintest.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}
