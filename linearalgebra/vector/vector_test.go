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