package linearalgebra

import (
	"errors"
)

// struct that represents a matrix
// The Data field store the contents of the matrix in row major order
// a matrix with one row and n columns is a called  row vector
// a matrix with one n rows and one column is called a column vector
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

// Builds a matrix with one column and N rows from a 1D slice
// N is the number of elements of the provided slice
func NewColumnVector(m []float64) Matrix {
	temp := make([][]float64, len(m))
	for i:= 0; i< len(temp); i++ {
		temp[i] = []float64{m[i]}
	}
	return NewMatrix(temp)
}

// Builds a matrix with N columns and 1 row from a 1D slice
// N is the number of elements of the provided slice
func NewRowVector(m []float64) Matrix {
	temp := make([][]float64, 1)
	temp[0] = m
	return NewMatrix(temp)
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
	for i := 0; i < m.Row; i++ {
		v1 := m.GetRow(i) // holds current row of this matrix as a vector
		for j := 0; j < m1.Col; j++ {
			v2 := m1.GetCol(j)                      // holds the current col of the m1 matrix (the one passed as argument) as a vector
			result[r_index], _ = DotProduct(v1, v2) //sets the dot product of the two vectors to the result slice
			r_index++
		}
	}

	return Matrix{
		Row:  m.Row,
		Col:  m1.Col,
		Data: result,
	}, nil
}

// inserts a vector into this matrix at the provided index
// if vector is a column vector a new column will be added to the marix at the specified index
// if vector is a row vector a new row will be inserted into this matrix at the provided index
func (m Matrix) InsertAt(v Matrix, index int) (Matrix, error) {
	if !isVector(v) {
		return Matrix{}, errors.New("column should be a Vector")
	}

	if v.Row > 1 {
		return insertCol(m, v.Data, index)
	} else {
		return insertRow(m, v.Data, index)
	}
}

// Inserts a row to a matrix at a given index, the index parameter is the index at which the new row will be added
// E.X
// if IsertCall is called in this matrix
// [1 2]
// [3 4]
// with row [8,9] and index 1
// This function with return the row major order quivalent of the Matrix
// [1,2]
// [8,9]
// [3,4]
func insertRow(m Matrix, row []float64, index int) (Matrix, error) {
	newMatrix := make([]float64, len(m.Data)+len(row))
	var newMatrixIndex int

	for i := 0; i < len(m.Data); i++ {
		if i == index {
			for j := 0; j < len(row); j++ {
				newMatrix[newMatrixIndex] = row[j]
				newMatrixIndex++
			}
		}
		newMatrix[newMatrixIndex] = m.Data[i]
		newMatrixIndex++
	}

	return Matrix{
		Row:  m.Row + 1,
		Col:  m.Col,
		Data: newMatrix,
	}, nil
}

// Inserts a column to a matrix at a given index, the index parameter is the index at which the new column will be added
// E.X
// if IsertCall is called in this matrix
// [1 2]
// [3 4]
// with colum [8,9] and index 1
// This function with return the row major order quivalent of the Matrix
// [1,8,2]
// [3,9,4]
func insertCol(m Matrix, column []float64, index int) (Matrix, error) {
	if len(column) < m.Row || len(column) > m.Row {
		return Matrix{}, errors.New("invalid column. Column lenght should be equal to matrix row length")
	}

	newMatrix := make([]float64, len(m.Data)+len(column))

	var mDataIndex int
	for i := 0; i < len(newMatrix); i++ {
		if i == index {
			newMatrix[i] = column[0]
			index = i + m.Col + 1 // calculates the next index on which the next element should be inserted
			column = column[1:]   // remove the first element of the column slice
		} else {
			newMatrix[i] = m.Data[mDataIndex] // copy data to new matrix slice
			mDataIndex++
		}
	}

	return Matrix{
		Row:  m.Row,
		Col:  m.Col + 1,
		Data: newMatrix,
	}, nil
}

// returns the the specified row of this matrix as a Vector
// this method assumes zero based index matrix, the first row index is 0
// the second row index is one, and so on...
func (m Matrix) GetRow(index int) Matrix {
	start := coordsToRowMajorIndex(index, 0, m.Col)
	end := start + m.Col
	temp := m.Data[start:end]
	// len 1 because a row vector will be returned
	rowMatrix := make([][]float64, 1)
	rowMatrix[0] = temp
	return NewMatrix(rowMatrix)
}

// returns the the specified row of this matrix as a Nx1 matrix
// this method assumes zero based index matrix, the first column index is 0
// the second column index is one, and so on...
func (m Matrix) GetCol(index int) Matrix {
	// the index of the first element of the column
	mIndex := coordsToRowMajorIndex(0, index, m.Col)
	jumps := m.Row
	i := 0
	col := make([][]float64, m.Row)
	for jumps > 0 {
		col[i] = []float64{m.Data[mIndex]}
		mIndex += m.Col
		jumps--
		i++
	}
	// end := start + m.Col
	// return vector.NewVector(m.Data[start:end])
	return NewMatrix(col)
}

// returns the value of the matrix at index [i,j] assumming that matrix.Data holds a valid row major order matrix
func (m Matrix) Get(i, j int) float64 {
	return m.Data[coordsToRowMajorIndex(i, j, m.Col)]
}
