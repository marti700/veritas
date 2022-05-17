// This files contains implementations of the usual operations that are performed on vectors
// in veritas vectors are matrices
// matrices with more than one column and one row are row vectors
// matrices with more that one row and one column are column vectors
package linearalgebra

import "errors"

// returns the number of elements of a vector
// an error is returned is the provides matrix is not a vector
func Size(v Matrix) (int, error) {
	if !IsVector(v) {
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

	if e1 != nil || e2 != nil {
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
func IsVector(v Matrix) bool {
	return IsColumnVector(v) || IsRowVector(v)
}

// returns true is the provided matrix is a row vector false otherwise
func IsRowVector(v Matrix) bool {
	return v.Row == 1
}

// returns true if the provided matrix is a column vector false otherwise
func IsColumnVector(v Matrix) bool {
	return v.Col == 1
}

// given a vector and a boolean function of the type float64 -> bool returns a new matrix with the elements for what the function returns true
//
// the axis parameter controls if the returned matrix will be a row or a column vector
// this function will panic if the axis parameter have any other value different from 0 or 1
func ElementWiseFilter(v Matrix, f func(n float64) bool, axis int) Matrix {
	if !IsVector(v) || (axis < 0 || axis > 1) {
		panic("Invalid parameters, make sure the provided matrix is a vector and that the axis parameter is either 0 or 1")
	}

	newVector := make([]float64, 0, len(v.Data))
	for _, e := range v.Data {
		if f(e) {
			newVector = append(newVector, e)
		}
	}

	if axis == 0 {
		return NewRowVector(newVector)
	}

	return NewColumnVector(newVector)
}

// given a vector and a boolean function of the type float64 -> bool returns two new matrix with the elements for what the function returns true
// the for which the function returns false
//
// the axis parameter controls if the returned matrix will be a row or a column vector
// this function will panic if the axis parameter have any other value different from 0 or 1
func ElementWiseFilter2(v Matrix, f func(n float64) bool, axis int) (Matrix, Matrix) {
	if !IsVector(v) || (axis < 0 || axis > 1) {
		panic("Invalid parameters, make sure the provided matrix is a vector and that the axis parameter is either 0 or 1")
	}

	newVectorT := make([]float64, 0, len(v.Data)) // holds true values
	newVectorF := make([]float64, 0, len(v.Data)) // holds false values
	for _, e := range v.Data {
		if f(e) {
			newVectorT = append(newVectorT, e)
		} else {
			newVectorF = append(newVectorF, e)
		}
	}

	if axis == 0 {
		return NewRowVector(newVectorT), NewRowVector(newVectorF)
	}

	return NewColumnVector(newVectorT), NewColumnVector(newVectorF)
}
