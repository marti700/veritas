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
	ans := linearalgebra.NewMatrix([][]float64 {
		{1,0,0,0,0},
		{0,1,0,0,0},
		{0,0,1,0,0},
		{0,0,0,1,0},
		{0,0,0,0,1},
	})

	result := linearalgebra.GenIdenityMatrix(5)

if !lintest.MatrixEq1(ans,result)  {
		t.Error("The anwer should be ", ans, " but was ", result)
	}
}
