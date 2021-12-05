package matrix

import (
	"errors"
	"veritas/linearalgebra/vector"
)

// struct that represents a matrix
// The Data field store the contents of the matrix in row major order
type Matrix struct {
	Row  int
	Col  int
	Data []float64
}

// Construct a matrix from a slice
func NewMatrix(m [][]float64) Matrix {
	rows := len(m)
	cols := len(m[0])

	return Matrix{
		Row:  rows,
		Col:  cols,
		Data: toRowMajor(m),
	}
}

//Returns the transpose of this matrix
func (m Matrix) T() Matrix {

	var newMatrix = make([]float64, m.Row*m.Col)
	var k int
	for i := 0; i < m.Col; i++ {
		for j := 0; j < m.Row; j++ {
			newMatrix[k] = m.Get(j, i)
			k++
		}
	}

	return Matrix{
		Row:  m.Col,
		Col:  m.Row,
		Data: newMatrix,
	}
}

// Returns a matrix which represents the result of multiplying this matrix wiht another
func (m Matrix) Mult(m1 Matrix) (Matrix, error) {

	if m.Col != m1.Row {
		return Matrix{}, errors.New("can't multiply matrices with diferent number of rows and columns")
	}
	var result = make([]float64, m.Row*m1.Col)
	var r_index int

	// kind of a unconventional algorithm, basically what is does is take the rows of this matrix (one by one)
	// and the columns of the m1 (the matrix passed as argument) and by using the dot product in each interaction
	// produces the result matrix in row major order. Basically it takes to matrices in row major order multiplies them
	// and the result will be also in row major order
	for i := 0; i < len(m.Data); i = (i+m.Col) {
		v1 := vector.NewVector(m.Data[i:(i + m.Col)]) // holds current row of this matrix as a vector
		for j := 0; j < m1.Col; j++ {
			v2 := vector.NewVector(selectElements(m1, j, m1.Col)) // holds the current row of the m1 matrix (the one passed as argument) as a vector
			result[r_index] = v1.DotProduct(v2) //sets the dot product of the two vectors to the result slice
			r_index++
		}
	}

	return Matrix{
		Row:  m.Row,
		Col:  m1.Col,
		Data: result,
	}, nil
}

// MATRIX UTILS

// transforms a slice to row major order
func toRowMajor(m [][]float64) []float64 {
	var rmo []float64
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			rmo = append(rmo, m[i][j])
		}
	}
	return rmo
}

// returns the value of the matrix at index [i,j] assumming that matrix#Data is a valid 1D slice and is in row major order
func (m Matrix) Get(i, j int) float64 {
	return m.Data[(i*m.Col)+j]
}
// Select x distant elements from matrix expresed in row major order
// EJ:
//[1,2,3,4,5,6] which is let's soupouse a 2X3 matrix
//a call to selectElements with stardIndex = 2 and increment = 3 will result in
// [3,6]
func selectElements(m Matrix, startIndex, increment int) []float64 {
	var v = make([]float64, m.Row)
	var vIndex int
	for i := startIndex; vIndex < m.Row; {
		v[vIndex] = m.Data[i]
		i = i + increment
		vIndex++
	}
	return v
}
