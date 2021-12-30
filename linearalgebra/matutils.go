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
