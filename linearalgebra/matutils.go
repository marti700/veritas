package linearalgebra

import (
	"errors"

	"github.com/marti700/veritas/commons"
)

// returns a matrix with the specified number of rows and columns
// which contents are ones
func Ones(r, c int) Matrix {
	temp := make([][]float64, r)
	for i := 0; i < r; i++ {
		temp[i] = make([]float64, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			temp[i][j] = 1
		}
	}
	return NewMatrix(temp)
}

// transforms a 2d slice to a row major order matrix
func toRowMajor(m [][]float64) []float64 {
	var rmo []float64
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			rmo = append(rmo, m[i][j])
		}
	}
	return rmo
}

// given the coordinates of a matrix (row,column) and the matrix total number of columns
// returns the equivalent matrix index in row major order
// i: the rows of the matrix
// j: the columns of the matrix
// m: the total number of columns of the matrix
func coordsToRowMajorIndex(i, j, m int) int {
	return (i * m) + j
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
	if len(row) < m.Col || len(row) > m.Col {
		return Matrix{}, errors.New("invalid row. row lenght should be equal to matrix column length")
	}
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

// Returns the result of adding all elements of provided matrix
func ElementsSum(m Matrix) float64 {
	return commons.Sum(m.Data)
}

// Returns an NXN identity matrix
// the argument d represents the dimensions of the matrx
func GenIdenityMatrix(d int) Matrix {
	newMatrix := make([]float64, d*d)
	for i := 0; i <= d*d; i+=d+1 {
		newMatrix[i] = 1
	}

	return Matrix{
		Row: d,
		Col: d,
		Data: newMatrix,
	}
}
