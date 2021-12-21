package commons_test

import (
	"testing"
	"github.com/marti700/veritas/commons"
)

func TestSum(t *testing.T) {

	var numbers = []float64{1,2,3,4,5,6,7,8,9}

	ans := commons.Sum(numbers)

	if ans != 45 {
		t.Error("commons.Sum expected result is: ", 45)
	}
}
