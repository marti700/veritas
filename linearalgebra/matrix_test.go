package linearalgebra_test

import (
	"fmt"
	"testing"

	"github.com/marti700/veritas/linearalgebra"
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

	if !linearalgebra.MatrixEq1(ans, expectedAnswer) {
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

	ans := m1.Mult(m2)
	if !linearalgebra.MatrixEq1(ans, result) {
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

	ans := m1.Mult(m2)
	if !linearalgebra.MatrixEq1(ans, result) {
		fmt.Println(ans)
		t.Error("answer should be: ", result, "but was: ", ans)
	}

}

func TestInsertAt(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{{1, 2}, {3, 4}})
	ans := linearalgebra.NewMatrix([][]float64{{8, 9}, {1, 2}, {3, 4}})
	result := m.InsertAt(linearalgebra.NewMatrix([][]float64{{8, 9}}), 0)

	if !linearalgebra.MatrixEq1(result, ans) {
		fmt.Println(ans)
		t.Error("answer should be: ", ans, "but was: ", result)
	}

	m1 := linearalgebra.NewMatrix([][]float64{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}})
	ans1 := linearalgebra.NewMatrix([][]float64{{0, 1, 1, 2}, {3, 2, 4, 5}, {6, 3, 7, 8}, {9, 4, 10, 11}})
	result1 := m1.InsertAt(linearalgebra.NewMatrix([][]float64{{1}, {2}, {3}, {4}}), 1)

	if !linearalgebra.MatrixEq1(result1, ans1) {
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

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestGetCol(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	})

	ans := linearalgebra.NewMatrix([][]float64{{1}, {3}, {5}})
	result := m.GetCol(0)

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestNewColumnVector(t *testing.T) {
	ans := linearalgebra.NewMatrix([][]float64{
		{1},
		{2},
		{3},
	})

	result := linearalgebra.NewColumnVector([]float64{1, 2, 3})

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestNewRowVector(t *testing.T) {
	ans := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
	})

	result := linearalgebra.NewRowVector([]float64{1, 2, 3})

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestSum(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2},
		{3, -2},
	})

	m1 := linearalgebra.NewMatrix([][]float64{
		{4, 5},
		{-6, -4},
	})

	ans := linearalgebra.NewMatrix([][]float64{
		{5, 7},
		{-3, -6},
	})

	result := m.Sum(m1)

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestSubsctract(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2},
		{3, -2},
	})

	m1 := linearalgebra.NewMatrix([][]float64{
		{4, 5},
		{-6, -4},
	})

	ans := linearalgebra.NewMatrix([][]float64{
		{-3, -3},
		{9, 2},
	})

	result := m.Substract(m1)

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}
func TestScleBy(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2},
		{3, -2},
	})

	ans := linearalgebra.NewMatrix([][]float64{
		{2, 4},
		{6, -4},
	})

	result := m.ScaleBy(2.0)

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestHadamrdProduct(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{
		{1, 2, 3},
		{2, 2, 2},
	})

	m1 := linearalgebra.NewMatrix([][]float64{
		{2, 2, 2},
		{1, 2, 3},
	})

	ans := linearalgebra.NewMatrix([][]float64{
		{2, 4, 6},
		{2, 4, 6},
	})

	result := m.HadamardProduct(m1)

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestInv(t *testing.T) {
m := linearalgebra.NewMatrix([][]float64{
		{2, 0, 1},
		{3, 0, 2},
		{5, -6, 7},
	})

	result := m.Mult(m.Inv())

	//A*A^(-1) = I
	if !linearalgebra.IsAnIdentityMatrix(result) {
		t.Error("the result: ", result, " is not equl the the identity matrix")
	}
}

func TestMap(t *testing.T) {

	m1 := linearalgebra.NewMatrix([][]float64 {
		{1,2,3,4},
		{5,6,7,8},
	})

	ans := linearalgebra.NewMatrix([][]float64 {
		{1,4,9,16},
		{25,36,49,64},
	})

	squareF := func (x float64) float64 {
		return x*x
	}

	result := m1.Map(squareF)

	if !linearalgebra.MatrixEq1(ans, result) {
		t.Error("answer should be: ", ans, "but was: ", result)
	}
}

func TestIsEmpty(t *testing.T) {
	m := linearalgebra.NewMatrix([][]float64{{}})
	m1 := linearalgebra.NewMatrix([][]float64 {
		{1,2,3,4},
		{5,6,7,8},
	})

	if !linearalgebra.IsEmpty(m) {
		t.Error("m is empty an empty")
	}

	if linearalgebra.IsEmpty(m1) {
		t.Error("M1 is not an empty matrix")
	}

}