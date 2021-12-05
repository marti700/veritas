package matrix_test

import (
	"fmt"
	"testing"
	"veritas/linearalgebra/matrix"
)

func TestTranspose(t *testing.T) {
	mat := matrix.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	ans := mat.T()

	expectedAnswer := matrix.NewMatrix([][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	})


	if !matrixEq1(ans, expectedAnswer) {
		t.Error("Expected result is: ", expectedAnswer, "but resultWas: ",   ans)
	}
}

func TestMult(t *testing.T) {
	m1 := matrix.NewMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	m2 := matrix.NewMatrix([][]float64{
		{10, 11},
		{20, 21},
		{30, 31},
	})

	result := matrix.NewMatrix([][]float64{
		{140, 146},
		{320, 335},
	})

	ans, _ := m1.Mult(m2)
	if !matrixEq1(ans, result) {
		fmt.Println(ans)
		t.Error("answer should be: ", result, "but was: ", ans)
	}

}

func TestMult1(t *testing.T) {
	m1 := matrix.NewMatrix([][]float64{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{9, 10, 11},
	})
	m2 := matrix.NewMatrix([][]float64{
		{0, 3, 6, 9},
		{1, 4, 7, 10},
		{2, 5, 8, 11},
	})

	result := matrix.NewMatrix([][]float64{
		{5,14,23,32},
		{14,50,86,122},
		{23,86,149,212},
		{32,122,212,302},

	})

	ans, _ := m1.Mult(m2)
	if !matrixEq1(ans, result) {
		fmt.Println(ans)
		t.Error("answer should be: ", result, "but was: ", ans)
	}

}

// Test utils

func matrixEq1(m1,m2 matrix.Matrix) bool {
	equalRows := m1.Row == m2.Row
	equalCols := m1.Col == m2.Col
	equalData := func() bool{
		if len(m1.Data) != len(m2.Data) {
			return false
		}

		for i,e := range m1.Data {
			if e != m2.Data[i] {
				return false
			}
		}
		return true
	}()

	return equalRows && equalCols && equalData
}
