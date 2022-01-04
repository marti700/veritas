package linearalgebra

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

