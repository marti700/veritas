// This files contains implementations of the usual operations that are performed on vectors
// in veritas vectors are matrices
// matrices with more than one column and one row are row vectors
// matrices with more that one row and one column are column vectors
package linearalgebra

import "errors"

// returns the number of elements of a vector
// an error is returned is the provides matrix is not a vector
func Size(v Matrix) (int, error) {
	if !isVector(v) {
		return 0, errors.New("Matrix is not a vector")
	}
	return v.Col * v.Row, nil
}

// Calculates the dot product of two vectors
// A error is returned if th matrices are not vectors or if they have different sizes
func DotProduct(v1, v2 Matrix) (float64, error) {
	var sum float64

	v1Size, e1 := Size(v1)
	v2Size, e2 := Size(v2)

	if e1 !=nil || e2 !=nil {
		return 0, errors.New("one or both of the matrices are not a vectors")
	}

	if v1Size != v2Size {
		return 0, errors.New("vectors must be of the same size")
	}

	for i := range v1.Data {
		sum += v1.Data[i] * v2.Data[i]
	}

	return sum, nil
}

// returns true if a matrix is a vector false otherwise
func isVector(v Matrix) bool {
	if v.Row == 1 && v.Col > 1 {
		return true
	} else if v.Row > 1 && v.Col == 1 {
		return true
	} else {
		return false
	}
}
