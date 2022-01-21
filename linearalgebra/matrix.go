package linearalgebra

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
	for i := 0; i < len(temp); i++ {
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

// retun a new matrix with all the entries of this matrix scaled by the given factor
func (m Matrix) ScaleBy(factor float64) Matrix {
	newMatrix := make([]float64, len(m.Data))
	for i := range m.Data {
		newMatrix[i] = factor * m.Data[i]
	}

	return Matrix{
		Row:  m.Row,
		Col:  m.Col,
		Data: newMatrix,
	}
}

// MATRIX ARITMETIC

// Returns a new matrix that represents the element wise multiplicaton of this matrix and another
// panic if the provided matrix dimensions are not of the same as this matrix
func (m Matrix) HadamardProduct(m1 Matrix) Matrix {
	if m.Col != m1.Col && m.Row != m1.Row {
		panic("error, matrices must be of the same dimension")
	}

	newMatrix := make([]float64, len(m.Data))
	for i := range m.Data {
		newMatrix[i] = m.Data[i] * m1.Data[i]
	}

	return Matrix{
		Row:  m.Row,
		Col:  m.Col,
		Data: newMatrix,
	}
}

// Returns a new matrix which represents the result of adding this matrix to another
// panic if the provided matrix dont have the same dimensions of this matrix
func (m Matrix) Sum(m1 Matrix) Matrix {
	if m.Col != m1.Col || m.Row != m1.Row {
		panic("can't add matrices with of diferen dimensions")
	}

	matrixSum := make([]float64, len(m.Data))
	for i := range m.Data {
		matrixSum[i] = m.Data[i] + m1.Data[i]
	}

	return Matrix{
		Row:  m.Row,
		Col:  m.Col,
		Data: matrixSum,
	}
}

// Returns a new matrix which represents the result of substracting this matrix from another
// panic if the provided matrix is not of the same dimensions of this matrix
func (m Matrix) Substract(m1 Matrix) Matrix {
	if m.Col != m1.Col || m.Row != m1.Row {
		panic("can't substract matrices with of diferen dimensions")
	}

	matrixSum := make([]float64, len(m.Data))
	for i := range m.Data {
		matrixSum[i] = m.Data[i] - m1.Data[i]
	}

	return Matrix{
		Row:  m.Row,
		Col:  m.Col,
		Data: matrixSum,
	}
}

// Returns a matrix that represents the result of multiplying this matrix to another
// panics if the provided matrix number of rows differs from this matrix number of columns
func (m Matrix) Mult(m1 Matrix) Matrix {

	if m.Col != m1.Row {
		panic("can't multiply matrices with diferent number of rows and columns")
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
	}
}

// MATRIX ARITMETIC END

// inserts a vector into this matrix at the provided index
// if vector is a column vector a new column will be added to the marix at the specified index
// if vector is a row vector a new row will be inserted into this matrix at the provided index
// panics if the provided matrix is not a vector
func (m Matrix) InsertAt(v Matrix, index int) Matrix {
	if !isVector(v) {
		panic("provided matrix should be a Vector")
	}

	if v.Row > 1 {
		return insertCol(m, v.Data, index)
	} else {
		return insertRow(m, v.Data, index)
	}
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

// finds the invere of a matrix using the Gauss-Jordan algorithm
// returns the inverse of this matrix
func (m Matrix) Inv() Matrix {
	// Generates an identity matrix
	iMat := GenIdenityMatrix(m.Col)

	//Build augmented matrix
	augmentedMatrix := m.InsertAt(iMat.GetCol(0), m.Col)
	for i := 1; i < iMat.Col; i++ {
		augmentedMatrix = augmentedMatrix.InsertAt(iMat.GetCol(i), augmentedMatrix.Col)
	}

	//to set the new values to the augmanted matrix
	apply := func(row int, m1, m2 Matrix) []float64 {
		gIndex := m1.Col * row
		for i := 0; i < len(m2.Data); i++ {
			m1.Data[gIndex] = m2.Data[i]
			gIndex++
		}
		return m1.Data
	}

	// to swap rows. This function search for entries that are not equal zero in the same
	// column of the curren pivot, when it find one, it swaps the rows and return a new augmented Matrix
	// which current pivot is a not zero entry
	swapRow := func(augM Matrix, row int) Matrix {
		col := row
		for k := row; k < augM.Row; k++ {
			if k != row && augM.Get(k, col) != 0 {
				r1 := augM.GetRow(row)           // get the first row
				r2 := augM.GetRow(k)             // get the first row with a non zero element
				augM.Data = apply(row, augM, r2) // apply first row to the augmented matrix
				augM.Data = apply(k, augM, r1)   // puts r1 where r2 were
				return augM
			}
		}
		panic("Not invertible matrix")
	}

	for i := 0; i < m.Col; i++ {
		if augmentedMatrix.Get(i, i) == 0 {
			augmentedMatrix = swapRow(augmentedMatrix, i)
		}
		pivotRow := augmentedMatrix.GetRow(i).ScaleBy(1 / augmentedMatrix.Get(i, i))
		// make the pivot (the element in the current row that is part of the main diagonal) 1 by multiplying the whole row by it's inverse
		augmentedMatrix.Data = apply(i, augmentedMatrix, pivotRow)
		for j := 0; j < m.Row; j++ {
			// if the element is not in the main diagonal
			if i != j {
				scaledPivot := pivotRow.ScaleBy(1 * augmentedMatrix.Get(j, i))
				n := augmentedMatrix.GetRow(j).Substract(scaledPivot)
				augmentedMatrix.Data = apply(j, augmentedMatrix, n)
			}
		}
	}
	return Slice(augmentedMatrix, m.Col, m.Col*2, "y")
}
