// This files contains implementations of the usual operations that are performed on vectors
// in veritas vectors are matrices
// matrices with more than one column and one row are row vectors
// matrices with more that one row and one column are column vectors
package linearalgebra

import "errors"

// returns the number of elements of a vector
func Size(v Matrix) (int, error) {
	if v.Col > 1 && v.Row > 1 {
		return 0, errors.New("Matrix is not a vector")
	}
	return v.Col * v.Row, nil
}
