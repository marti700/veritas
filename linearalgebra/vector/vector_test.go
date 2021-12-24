package vector_test

import (
	"github.com/marti700/veritas/linearalgebra/vector"
	"testing"
)

func TestDotProduct(t *testing.T) {
	vec1 := vector.NewVector([]float64{1,2,3})
	vec2 := vector.NewVector([]float64{4,5,6})

	ans := vec1.DotProduct(vec2)

	if ans != 32 {
		t.Error("answer should be 32: ", ans)
	}
}

func TestOnes(t *testing.T) {
	ans := vector.NewVector([]float64{1,1,1,1,1,1,1,1,1,1})
	result := vector.Ones(10)

	if !eq(ans, result)  {
		t.Error("Answer should be ", ans, "but was ",result)
	}
}

// TEST UTILS

// Tests if two vectors are equal
func eq (v1,v2 vector.Vector) bool {
	if v1.Size != v2.Size {
		return false
	}
	for i := 0; i<v1.Size; i++ {
		if v1.Data[i] != v2.Data[i] {
			return false
		}
	}
	return true
}